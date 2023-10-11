package main

import (
  "strings"
  "slices"
)

func Tokenise(exp string) []string {
  return strings.Split(exp, " ")
}

func Parse(tokens []string) Expression {
  level := make([]int, len(tokens))
  for lvl, i := 0, 0; i < len(tokens); i++ {
    switch tokens[i] {
    case "(": {
      lvl++
      level[i] = lvl
    }
    case ")": {
      level[i] = lvl
      lvl--
    }
    default: 
      level[i] = lvl
    }
  }

  // remove enclosing parenthesis if they exist
  remParenth := tokens[0] == "(" && tokens[len(tokens)-1] == ")"
  if remParenth {
    for i, val := range level {
      if 0 < i && i < len(level)-1 && val < level[0] {
        remParenth = false
      }
    }
  }
  if remParenth {
    tokens = tokens[1:len(tokens)-1]
    level = level[1:len(level)-1]
  }

  op, pos := 0, 0

  minLvl := slices.Min(level)
  for i, val := range tokens {
    opNum := OpStrToNum(val)
    if opNum == 0 || level[i] > minLvl {
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
