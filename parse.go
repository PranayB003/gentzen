package main

import (
  "strings"
  "log"
)

func Tokenise(exp string) []string {
  return strings.Split(exp, " ")
}

func Parse(tokens []string) Expression {
 op, pos := 0, 0

 for i:= 0; i < len(tokens); i++ {
   if tokens[i] == _dimp {
     op, pos = 5, i
     break
   } else if tokens[i] == _imp && op < 4 {
     op, pos = 4, i
   } else if tokens[i] == _or  && op < 3 {
     op, pos = 3, i
   } else if tokens[i] == _and  && op < 2 {
     op, pos = 2, i
   } else if tokens[i] == _not  && op < 1 {
     op, pos = 1, i
   }
 } 

 var exp Expression
 if op == 0 {
   if len(tokens) > 1 {
     log.Fatal("No operator between operands")
   }
   exp = Expression{nil, nil, tokens[0], 1}
 } else if op == 1 {
   if tokens[0] != _not {
     log.Fatal("Invalid usage of NOT operator")
   }
   rightExp := Parse(tokens[pos+1:])
   exp = Expression{nil, &rightExp, _not, 2}
 } else if op == 2 {
   leftExp, rightExp := Parse(tokens[:pos]), Parse(tokens[pos+1:])
   exp = Expression{&leftExp, &rightExp, _and, 3}
 } else if op == 3 {
   leftExp, rightExp := Parse(tokens[:pos]), Parse(tokens[pos+1:])
   exp = Expression{&leftExp, &rightExp, _or, 3}
 } else if op == 4 {
   leftExp, rightExp := Parse(tokens[:pos]), Parse(tokens[pos+1:])
   exp = Expression{&leftExp, &rightExp, _imp, 3}
 } else if op == 5 {
   leftExp, rightExp := Parse(tokens[:pos]), Parse(tokens[pos+1:])
   exp = Expression{&leftExp, &rightExp, _dimp, 3}
 }

 return exp
}
