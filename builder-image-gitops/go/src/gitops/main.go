package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/src-d/go-git.v4"

	// "gopkg.in/src-d/go-git.v4/storage/memory"
	// "gopkg.in/src-d/go-git.v4/plumbing/object"
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

func main() {
	fmt.Println("Hi!")
	Info("git clone https://github.com/liatrio/lead-environments.git")

	r, err := git.PlainClone("/repo", false, &git.CloneOptions{
		Auth: &http.BasicAuth{
			Username: "ChrisSchreiber",
			Password: "xxxx",
		},
		URL:      "https://github.com/liatrio/lead-environments.git",
		Progress: os.Stdout,
	})

	CheckIfError(err)

	// // Gets the HEAD history from HEAD, just like this command:
	// Info("git log")

	// ... retrieves the branch pointed by HEAD
	ref, err := r.Head()
	_ = ref
	CheckIfError(err)

	// // ... retrieves the commit history
	// cIter, err := r.Log(&git.LogOptions{From: ref.Hash()})
	// CheckIfError(err)

	// // ... just iterates over the commits, printing it
	// err = cIter.ForEach(func(c *object.Commit) error {
	// 	fmt.Println(c)
	// 	return nil
	// })
	// CheckIfError(err)

	data, err := ioutil.ReadFile("/repo/lead-environments/aws/liatrio-sandbox/terragrunt.hcl")

	hclParseTree, err := hcl.Parse(string(data))
	CheckIfError(err)

	fmt.Println(string(data))

	fmt.Println(spew.Sdump(hclParseTree))
}
