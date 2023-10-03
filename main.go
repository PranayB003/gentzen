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
    fmt.Println(tokens)
  }
}


