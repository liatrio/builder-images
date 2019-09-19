package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"path/filepath"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
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
	fmt.Printf("Cloning git repo %s\n", url)
	return git.PlainClone(path, false, &git.CloneOptions{
		Auth:     auth,
		URL:      url,
	})
}

func gitCommit() {

}

func gitPush() {

}

func parseFile(basePath string, repoFile string) (map[string]interface{}, error) {
	path := basePath + repoFile
	ext := filepath.Ext(repoFile)

	fmt.Printf("Parsing file %s\n", path)

	contents, err := ioutil.ReadFile(path)
	if (err != nil) {
		return nil, err
	}
	switch ext {
	case ".hcl":
		return decodeHcl(string(contents))
	default:
		return nil, fmt.Errorf("Unhandled file type '%s' for file '%s'", ext, repoFile)
	}
}

func decodeHcl(hclString string) (map[string]interface{}, error) {
	hclObj := make(map[string]interface{})

	fmt.Println("Decoding HCL data")

	err := hcl.Decode(&hclObj, hclString)
	if (err != nil) {
		return nil, err
	}

	return hclObj, nil
}

func encodeHcl() {

}

func setValueInMap(target interface{}, path []string, value interface{}) error {
	if reflect.TypeOf(target).Kind() == reflect.Slice {
		for _, refSlice := range target.([]map[string]interface{}) {
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
		fmt.Printf("Update value %s: %s -> %s\n",path[0], target.(map[string]interface{})[path[0]], value)
		target.(map[string]interface{})[path[0]] = value
	} else {
		setValueInMap(target.(map[string]interface{})[path[0]], path[1:], value)
	}

	return nil
}

type valuePath struct {
	path []string
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
	if (message != "") {
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
	gitAuthPassword := flag.String(
		"gitPassword", 
		os.Getenv("GITOPS_GIT_PASSWORD"), 
		"Password or token to authenticate with git. Can also use GITOPS_GIT_PASSWORD environment variable")
	repoPath := "/home/gitops/repo/"
	repoFile := flag.String(
		"repoFile", 
		os.Getenv("GITOPS_REPO_FILE"), 
		"File in git repo to apply changes to. Can also use GITOPS_REPO_FILE environment variable")
	values := flag.String(
		"values", 
		os.Getenv("GITOPS_VALUES"), 
		"List of variables and coresponding values to update. Variables paths are a list of keys separated with periods. Each variable is separated with a colon. Example '-values=input.one=foo:input.two=bar'. Can also use GITOPS_VALUES environment variable")

	flag.Parse()

	fmt.Println("Start GitOps")

	if (*gitURL == "") {
		usage("ERROR: Git URL is required!")
	}

	if (*gitUsername == "") {
		usage("ERROR: Git username is required!")
	}

	if (*gitAuthPassword == "") {
		usage("ERROR: Git password is required!")
	}

	if *repoFile == "" {
		usage("ERROR: File is required!")
	}

	if *values == "" {
		usage("ERROR: Values are required!")
	}
	valuePaths, err := parseValues(*values)
	if (err != nil) {
		usage("ERROR: Could not parse values")
	}

	gitAuth := &http.BasicAuth{
		Username: *gitUsername,
		Password: *gitAuthPassword,
	}

	_, err = gitClone(*gitURL, gitAuth, repoPath)
	CheckIfError(err)

	content, err := parseFile(repoPath, *repoFile)
	CheckIfError(err)

	for _, value := range valuePaths {
		err = setValueInMap(content, value.path, value.value)
		CheckIfError(err)	
	}

	fmt.Println("SPEW\n", spew.Sdump(content))
}
