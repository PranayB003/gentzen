package main

import (
  "bufio"
  "os"
  "fmt"
  "log"
)

var _optMap map[string]bool

func main() {
  scanner := bufio.NewScanner(os.Stdin)

  options := os.Args[1:]
  _optMap = make(map[string]bool)
  for _, opt := range options {
    _optMap[opt] = true
  }

  for {
    fmt.Print(">> ")
    
    scanner.Scan()
    err := scanner.Err()
    if err != nil {
      log.Fatal(err)
    }

    if scanner.Text() == "" {
      continue
    }
    
    tokens := Tokenise(scanner.Text())
    
    exp := Parse(tokens)

    if _optMap["--debug-parse"] {
      fmt.Println("tokens: ", tokens)
      fmt.Println("expression: ", exp.Printexp(), "\n")
    }

    valid := exp.Proove()
    if valid {
      fmt.Println("VALID")
    } else {
      fmt.Println("INVALID")
    }
  }
}


