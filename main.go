package main

import (
  "bufio"
  "os"
  "fmt"
  "log"
)

func main() {
  scanner := bufio.NewScanner(os.Stdin)

  for {
    fmt.Print(">> ")
    
    scanner.Scan()
    err := scanner.Err()
    if err != nil {
      log.Fatal(err)
    }
    
    tokens := Tokenise(scanner.Text())
    fmt.Println("tokens: ", tokens)
    
    exp := Parse(tokens)
    fmt.Println("expression: ", exp.Printexp())

    valid := exp.Proove()
    if !valid {
      fmt.Println("VALID")
    } else {
      fmt.Println("INVALID")
    }
  }
}


