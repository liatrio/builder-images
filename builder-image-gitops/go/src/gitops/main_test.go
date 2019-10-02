package main

import (
  "os"
  "testing"
  "reflect"
  "github.com/hashicorp/hcl"
  "github.com/hashicorp/hcl/hcl/printer"
)

func TestDecodeHcl(t *testing.T) {
  hclString := "inputs = {\n  one = \"bar\"\n}"
  f, err := decodeHcl(hclString)

  if reflect.TypeOf(*f).Name() != "File" {
    t.Error("Incorrect type returned from decodeHcl")
  }

  if err != nil {
    t.Error(err)
  }
}

func TestSetValueInAst(t *testing.T) {
  hclString := "inputs = {\n  one = \"bar\"\n}"

  ast, _ := hcl.Parse(hclString)

  //First test runs correct input into setValueInAst to change 'bar' to 'foo'
  err := setValueInAst(*ast, []string{"inputs", "one"}, "foo")

  file, err := os.Create("./testsetvalue.hcl")
  err = printer.Fprint(file, ast.Node)

  fileInfo, err := file.Stat()

  data := make([]byte, fileInfo.Size())
  file.Seek(0,0)
  _, err = file.Read(data)
  hclStringResult := "inputs = {\n  one = \"foo\"\n}"

  if string(data) != hclStringResult {
    t.Error("Data does not match")
  }

  if err != nil {
    t.Error(err)
  }

  os.Remove("./testsetvalue.hcl")

  //Second test tries to access a value that doesn't exist 'two'
  err = setValueInAst(*ast, []string{"inputs", "missing", "two"}, "foo")

  if err.Error() != "Did not match value ([missing two] -> foo)" {
   t.Error(err)
  }
}

func TestParseValues(t *testing.T) {
  values := "input.one=foo"
  valuePath, err := parseValues(values)
  if err != nil {
    t.Error(err)
  }

  if valuePath[0].path[0] != "input" {
    t.Error("Values parsed incorrectly, expected 'input' got " + valuePath[0].path[0])
  }

  if valuePath[0].path[1] != "one" {
    t.Error("Values parsed incorrectly, expected 'one' got " + valuePath[0].path[1])
  }

  if valuePath[0].value != "foo" {
    str := valuePath[0].value.(string)
    t.Error("Values parsed incorrectly, expected 'foo' got " + str)
  }
}

