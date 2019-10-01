package main

import (
  "testing"
  "github.com/hashicorp/hcl"
)

func TestDecodeHcl(t *testing.T) {

  hclString := "inputs = {\n  one = \"bar\"\n}"
  ast, err := hcl.Parse(hclString)

  if err != nil {
    t.Error(err)
    t.Log(ast)
  }
}

func TestSetValueInAst(t *testing.T) {
  hclString := "inputs = {\n  one = \"bar\"\n}"

  ast, _ := hcl.Parse(hclString)
  values := "inputs.one=foo"
  valuePath, _ := parseValues(values)

  err := setValueInAst(*ast, valuePath[0].path, valuePath[0].value)
  if err != nil {
    t.Error(err)
  }
}

func TestParseValues(t *testing.T) {
  values := "input.one=foo"
  valuePath, err := parseValues(values)
  if err != nil {
    t.Error(err)
    t.Log(valuePath)
  }
}

