package main

import (
  "bufio"
  "os"
  "fmt"
  "log"
)

func main() {
  scanner := bufio.NewScanner(os.Stdin)

  options := os.Args[1:]
  optMap := make(map[string]bool)
  for _, opt := range options {
    optMap[opt] = true
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

    if optMap["--debug-parse"] {
      fmt.Println("tokens: ", tokens)
      fmt.Println("expression: ", exp.Printexp(), "\n")
    }

    valid := exp.Proove()
    if !valid {
      fmt.Println("\nVALID")
    } else {
      fmt.Println("\nINVALID")
    }
  }
}


