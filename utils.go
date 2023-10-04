package main

// returns a unique number for each operator,
// a lower number signifies higher precedence
func OpStrToNum(operator string) (res int) {
  switch operator {
  case _not:
    res = 1
  case _and:
    res = 2
  case _or:
    res = 3
  case _imp:
    res = 4
  case _dimp:
    res = 5
  default:
    res = 0
  }
  return
}

func (exps Expressions) Append(newElems ...Expression) Expressions {
  res := make(Expressions, len(exps))
  copy(res, exps)
  res = append(res, newElems...)
  return res
} 

func (exp Expression) Printexp() string {
  var res string
  if exp.etype == _term {
    res = exp.mid
  } else if exp.etype == _unary {
    res = exp.mid + " " + exp.right.Printexp()
  } else if exp.etype == _binary {
    res = exp.left.Printexp() + " " + exp.mid + " " + exp.right.Printexp()
  }
  res = "(" + res + ")"
  return res
}

func (seq Sequent) Printseq() (res string) {
  ant, con := seq.ant, seq.con
  for i := 0; i < len(ant); i++ {
    if i == 0 {
      res += ant[i].Printexp()
    } else {
      res = res + ", " + ant[i].Printexp()
    }
  }
  res = res + " => "
  for i := 0; i < len(con); i++ {
    if i == 0 {
      res += con[i].Printexp()
    } else {
      res = res + ", " + con[i].Printexp()
    }
  }
  return
}

func (seqs Sequents) Printseqs() (res string) {
  for i := 0; i < len(seqs); i++ {
    res = res + seqs[i].Printseq() + "     "
  }
  return
}
