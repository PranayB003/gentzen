package main

import (
  "strings"
)

func Tokenise(exp string) []string {
  return strings.Split(exp, " ")
}
