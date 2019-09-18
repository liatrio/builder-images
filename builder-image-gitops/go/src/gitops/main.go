package main

import (
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"flag"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/hcl"
)

// CheckIfError should be used to naively panics if an error is not nil.
func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

// Info should be used to describe the example commands that are about to run.
func Info(format string, args ...interface{}) {
	fmt.Printf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

func gitClone(url string, auth transport.AuthMethod, path string) (*git.Repository, error) {
	return git.PlainClone(path, false, &git.CloneOptions{
		Auth: 		auth,
		URL:      url,
		Progress: os.Stdout,
	})
}

func gitCommit() {

}

func gitPush() {

}

func decodeHcl() {

}

func encodeHcl() {

}

func setValueInMap(target interface{}, path []string, value interface{}) error {
	if reflect.TypeOf(target).Kind() == reflect.Slice {
		for _, refSlice := range target.([]map[string]interface{}) {
			fmt.Printf("T: %T\n", refSlice)
			setValueInMap(refSlice, path, value)
		}
		return nil
	}

	if reflect.TypeOf(target).Kind() != reflect.Map {
		return fmt.Errorf(fmt.Sprintf("Cannot map path %s type '%T'", path, target))
	}

	if _, exists := target.(map[string]interface{})[path[0]]; !exists {
		return fmt.Errorf(fmt.Sprintf("Cannot map path %s. '%s' not found.", path, path[0]))
	}

	if len(path) == 1 {
		fmt.Printf("Replace %s with %s\n", target.(map[string]interface{})[path[0]], value)
		target.(map[string]interface{})[path[0]] = value
	} else {
		setValueInMap(target.(map[string]interface{})[path[0]], path[1:], value)
	}

	return nil
}

func main() {
	gitURL := flag.String("gitUrl", os.Getenv("gitUrl"), "URL of git repository")
	gitAuthUsername := flag.String("gitAuthUsername", os.Getenv("gitAuthUsername"), "Username to authenticate with git")
	gitAuthPassword := flag.String("gitAuthPassword", os.Getenv("gitAuthPassword"), "Password or token to authenticate with git")
	repoPath := "/home/gitops/repo"

	fmt.Printf("Start GitOps (git URL: %s, git auth user: %s, git auth password %s)", *gitURL, *gitAuthU)

	gitAuth := &http.BasicAuth{
		Username: *gitAuthUsername,
		Password: *gitAuthPassword,
	}

	_, err := gitClone(*gitURL, gitAuth, repoPath)
	CheckIfError(err)

	data, err := ioutil.ReadFile("/home/gitops/repo/aws/liatrio-sandbox/terragrunt.hcl")

	hclout := make(map[string]interface{})

	err = hcl.Decode(&hclout, string(data))
	CheckIfError(err)

	err = setValueInMap(hclout, []string{"inputs", "sdm_version"}, "v100000.0.0")
	CheckIfError(err)

	fmt.Println("SPEW\n", spew.Sdump(hclout))
}
