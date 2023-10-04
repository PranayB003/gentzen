package main

import (
  "strings"
)

func Tokenise(exp string) []string {
  return strings.Split(exp, " ")
}

func Parse(tokens []string) Expression {
  enclPrths := true
  for i, val := range tokens {
    if (0 < i && i < len(tokens)-1) && (val == "(" || val == ")") {
      enclPrths = false
      break
    }
  }
  enclPrths = enclPrths && tokens[0] == "(" && tokens[len(tokens)-1] == ")"
  if enclPrths {
    // the expression is encloses by parenthesis, which can be omitted
    tokens = tokens[1:len(tokens)-1]  
  }

  isOuterTerm := make([]bool, len(tokens))
  for lvl, i := 0, 0; i < len(tokens); i++ {
    switch tokens[i] {
    case "(": 
      lvl++
    case ")":
      lvl--
    default: {
      if lvl == 0 {
        isOuterTerm[i] = true
      }
    }
    }
  }

  op, pos := 0, 0

  for i, val := range tokens {
    opNum := OpStrToNum(val)
    if opNum == 0 || isOuterTerm[i] == false {
      continue
    }

    if val == _dimp {
      op, pos = OpStrToNum(_dimp), i
      break
    } else if val == _imp && op < opNum {
      op, pos = OpStrToNum(_imp), i
    } else if val == _or  && op < opNum {
      op, pos = OpStrToNum(_or), i
    } else if val == _and  && op < opNum {
      op, pos = OpStrToNum(_and), i
    } else if val == _not  && op < opNum {
      op, pos = OpStrToNum(_not), i
    }
  } 

  var exp Expression
  if op == 0 {
    exp = Expression{nil, nil, tokens[0], _term}
  } else if op == 1 {
    rightExp := Parse(tokens[pos+1:])
    exp = Expression{nil, &rightExp, _not, _unary}
  } else if op == 2 {
    leftExp, rightExp := Parse(tokens[:pos]), Parse(tokens[pos+1:])
    exp = Expression{&leftExp, &rightExp, _and, _binary}
  } else if op == 3 {
    leftExp, rightExp := Parse(tokens[:pos]), Parse(tokens[pos+1:])
    exp = Expression{&leftExp, &rightExp, _or, _binary}
  } else if op == 4 {
    leftExp, rightExp := Parse(tokens[:pos]), Parse(tokens[pos+1:])
    exp = Expression{&leftExp, &rightExp, _imp, _binary}
  } else if op == 5 {
    leftExp, rightExp := Parse(tokens[:pos]), Parse(tokens[pos+1:])
    exp = Expression{&Expression{&leftExp, &rightExp, _imp, _binary},
                     &Expression{&rightExp, &leftExp, _imp, _binary},
                     _and, _binary}
  }

  return exp
}
