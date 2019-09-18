package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"

	"gopkg.in/src-d/go-git.v4"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/hcl"

	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
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

func setValueInMap(values interface{}, path []string, value interface{}) (error) {
	if (reflect.TypeOf(values).Kind() == reflect.Slice) {
		for _, refSlice := range values.([]map[string]interface {}) {
			fmt.Printf("T: %T\n", refSlice)
			setValueInMap(refSlice, path, value)
		}
		return nil
	}

	if reflect.TypeOf(values).Kind() != reflect.Map {
		return fmt.Errorf(fmt.Sprintf("Cannot map path %s type '%T'", path, values))
	}

	if _, exists := values.(map[string] interface{})[path[0]]; !exists {
		return fmt.Errorf(fmt.Sprintf("Cannot map path %s. '%s' not found.", path, path[0]))
	}

	if len(path) == 1 {
		fmt.Printf("Replace %s with %s\n", values.(map[string] interface{})[path[0]], value);
		values.(map[string] interface{})[path[0]] = value
	} else {
		setValueInMap(values.(map[string]interface{})[path[0]], path[1:], value)
	}

	return nil
}

func main() {
	fmt.Println("Hi!")

	Info("git clone https://github.com/liatrio/lead-environments.git")

	_, err := git.PlainClone("/home/gitops/repo", false, &git.CloneOptions{
		Auth: &http.BasicAuth{
			Username: "ChrisSchreiber",
			Password: os.Getenv("GITHUB_TOKEN"),
		},
		URL:      "https://github.com/liatrio/lead-environments.git",
		Progress: os.Stdout,
	})
	CheckIfError(err)

	data, err := ioutil.ReadFile("/home/gitops/repo/aws/liatrio-sandbox/terragrunt.hcl")

  hclout := make(map[string]interface{})

	err = hcl.Decode(&hclout, string(data))
	CheckIfError(err)

	err = setValueInMap(hclout, []string{"inputs", "sdm_version"}, "v100000.0.0")
	CheckIfError(err)
	
	fmt.Println("SPEW\n", spew.Sdump(hclout))
}
