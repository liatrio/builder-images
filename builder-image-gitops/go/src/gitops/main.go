package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"strings"
	"time"
	"net/http"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	githttp "gopkg.in/src-d/go-git.v4/plumbing/transport/http"

	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/ast"
	"github.com/hashicorp/hcl/hcl/printer"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// CheckIfError should be used to naively panics if an error is not nil.
func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

func gitClone(url string, auth transport.AuthMethod, repoPath string) (*git.Repository, error) {
	fmt.Printf("Cloning git repo %s with %s\n", url, auth.String())
	return git.PlainClone(repoPath, false, &git.CloneOptions{
		Auth: auth,
		URL:  url,
	})
}

func gitBranch(repo *git.Repository) (branch *plumbing.Reference, worktree *git.Worktree, err error) {
	name := fmt.Sprintf("gitops-%x", time.Now().Unix())
	fmt.Printf("Create branch %s\n", name)

	head, err := repo.Head()
	if err != nil {
		return
	}

	branch = plumbing.NewHashReference(plumbing.NewBranchReferenceName(name), head.Hash())

	err = repo.Storer.SetReference(branch)

	worktree, err = repo.Worktree()

	worktree.Checkout(&git.CheckoutOptions{
		Branch: branch.Name(),
	})

	return
}

func gitCommit(worktree *git.Worktree, file string) (err error) {
	fmt.Printf("Add changes to repo %s\n", file)
	_, err = worktree.Add(file)
	if err != nil {
		return
	}

	fmt.Println("Commiting changes")
	worktree.Commit(fmt.Sprintf("GitOps: Update %s", file), &git.CommitOptions{
		Author: &object.Signature{
			Name:  "GitOps Automation",
			Email: "gitops@liatr.io",
			When:  time.Now(),
		},
	})

	return
}

func gitPush(repo *git.Repository, auth transport.AuthMethod) (err error) {
	fmt.Println("Pushing changes")
	err = repo.Push(&git.PushOptions{
		Auth: auth,
	})
	return
}

func githubPullRequest(httpClient *http.Client, org string, repo string, branch *plumbing.Reference) (pullRequest *github.PullRequest, err error) {
	fmt.Printf("Create pull request for branch %s\n", branch.Name().Short())
	client := github.NewClient(httpClient)

	newPR := &github.NewPullRequest{
		Title:               github.String("GitOps"),
		Head:                github.String(branch.Name().Short()),
		Base:                github.String("master"),
		Body:                github.String("This pull request was automatically generated https://github.com/liatrio/builder-images/tree/master/builder-image-gitops"),
		MaintainerCanModify: github.Bool(true),
	}

	pullRequest, _, err = client.PullRequests.Create(context.Background(), org, repo, newPR)

	return
}

func parseFile(filePath string) (*ast.File, error) {
	ext := path.Ext(filePath)

	fmt.Printf("Parsing file %s\n", filePath)

	contents, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	switch ext {
	case ".hcl":
		return decodeHcl(string(contents))
	default:
		return nil, fmt.Errorf("Unhandled file type '%s' for file '%s'", ext, filePath)
	}
}

func decodeHcl(hclString string) (*ast.File, error) {
	fmt.Println("Decoding HCL data")

	ast, err := hcl.Parse(hclString)
	CheckIfError(err)

	return ast, nil
}

func updateFile(filePath string, data *ast.File) (err error) {
	ext := path.Ext(filePath)

	fmt.Printf("Updating file %s\n", filePath)

	switch ext {
	case ".hcl":
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer file.Close()
		err = printer.Fprint(file, data.Node)
	default:
		err = fmt.Errorf("Unhandeld file extension for %s", filePath)
	}
	return
}

func setValueInAst(target ast.File, path []string, value interface{}) error {
	fmt.Printf("Changing value of %v\n", path)
	matched := false
	ast.Walk(target.Node, func(n ast.Node) (ast.Node, bool) {
		if item, ok := n.(*ast.ObjectItem); ok {
			for _, key := range item.Keys {
				if key.Token.Type.IsIdentifier() && key.Token.Text == path[0] {
					if len(path) == 1 {
						if val, ok := item.Val.(*ast.LiteralType); ok {
							fmt.Printf("Changed value %s -> \"%s\"\n", item.Val.(*ast.LiteralType).Token.Text, value)
							val.Token.Text = fmt.Sprintf("\"%s\"", value)
							matched = true
						} else {
							fmt.Printf("Warning: Cannot change value of type %T\n", item.Val)
						}
						return n, false // we matched the end of the path
					}
					path = path[1:]
					return n, true // we matched part of the path
				}
				return n, false // this branch does not match our path
			}
		}
		return n, true // traverse all non item nodes
	})
	if matched == false {
		return fmt.Errorf("Did not match value (%v -> %s)", path, value)
	}
	return nil
}

type valuePath struct {
	path  []string
	value interface{}
}

func parseValues(values string) ([]valuePath, error) {
	items := strings.Split(values, ":")
	valueMapList := make([]valuePath, len(items))
	for index, item := range items {
		keyvalue := strings.Split(item, "=")
		valueMapList[index] = valuePath{strings.Split(keyvalue[0], "."), keyvalue[1]}
	}
	return valueMapList, nil
}

func usage(message string) {
	if message != "" {
		fmt.Println(message)
	}
	flag.Usage()
	os.Exit(1)
}

func main() {
	gitURL := flag.String(
		"gitUrl",
		os.Getenv("GITOPS_GIT_URL"),
		"URL of git repository. Can also use GITOPS_GIT_URL environment variable")
	gitUsername := flag.String(
		"gitUsername",
		os.Getenv("GITOPS_GIT_USERNAME"),
		"Username to authenticate with git. Can also useGITOPS_GIT_USERNAME environment variable ")
	gitPassword := flag.String(
		"gitPassword",
		os.Getenv("GITOPS_GIT_PASSWORD"),
		"Password or token to authenticate with git. Can also use GITOPS_GIT_PASSWORD environment variable")
	repoPath := "/home/jenkins/repo/"
	repoFile := flag.String(
		"repoFile",
		os.Getenv("GITOPS_REPO_FILE"),
		"File in git repo to apply changes to. Can also use GITOPS_REPO_FILE environment variable")
	values := flag.String(
		"values",
		os.Getenv("GITOPS_VALUES"),
		"List of variables and coresponding values to update. Variables paths are a list of keys separated with periods. Each variable is separated with a colon. Example '-values=input.one=foo:input.two=bar'. Can also use GITOPS_VALUES environment variable")

	flag.Parse()

	filePath := repoPath + *repoFile

	fmt.Println("Start GitOps")

	if *gitURL == "" {
		usage("ERROR: Git URL is required!")
	}

	if *gitUsername == "" || gitUsername == nil {
		usage("ERROR: Git username is required!")
	}

	if *gitPassword == "" || gitPassword == nil {
		usage("ERROR: Git password is required!")
	}

	if *repoFile == "" {
		usage("ERROR: File is required!")
	}

	if *values == "" {
		usage("ERROR: Values are required!")
	}
	valuePaths, err := parseValues(*values)
	if err != nil {
		usage("ERROR: Could not parse values")
	}

	gitAuth := &githttp.BasicAuth{
		Username: *gitUsername,
		Password: *gitPassword,
	}

	repo, err := gitClone(*gitURL, gitAuth, repoPath)
	CheckIfError(err)

	branch, worktree, err := gitBranch(repo)
	CheckIfError(err)

	ast, err := parseFile(filePath)
	CheckIfError(err)

	for _, value := range valuePaths {
		err = setValueInAst(*ast, value.path, value.value)
		CheckIfError(err)
	}

	err = updateFile(filePath, ast)
	CheckIfError(err)

	err = gitCommit(worktree, *repoFile)
	CheckIfError(err)

	err = gitPush(repo, gitAuth)
	CheckIfError(err)

	gitURLParts, err := url.Parse(*gitURL)
	CheckIfError(err)

	if gitURLParts.Host == "github.com" {
		tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: *gitPassword})
		tokenClient := oauth2.NewClient(context.Background(), tokenSource)

		pathParts := strings.Split(strings.TrimSuffix(gitURLParts.Path, ".git"), "/")

		pullRequest, err := githubPullRequest(tokenClient, pathParts[1], pathParts[2], branch)
		CheckIfError(err)

		fmt.Printf("Pull Request created: %s\n", pullRequest.GetHTMLURL())
	}
}
