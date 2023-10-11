package main

import (
  "bufio"
  "os"
  "fmt"
  "log"
)

var _optMap map[string]bool

func main() {
  options := os.Args[1:]
  _optMap = make(map[string]bool)
  for _, opt := range options {
    _optMap[opt] = true
  }

  if _optMap["--no-interactive"] {
    var expPtr *Expression
    var exp Expression
    // define a custom expression to be evaluated
    a := Term("a")
    b := Term("b")
    c := Term("c")

    expPtr = Implication( Implication(a, c), Implication( Implication(b, c), Implication(Disjunction(a, b), c)))
    exp = *expPtr

    if _optMap["--debug-parse"] {
      fmt.Println("expression: ", exp.Printexp(), "\n")
    }

    valid := exp.Proove()
    if valid {
      fmt.Println("VALID")
    } else {
      fmt.Println("INVALID")
    }

    return
  }

  scanner := bufio.NewScanner(os.Stdin)

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


