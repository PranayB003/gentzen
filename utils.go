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

func (exps Expressions) AppendUnique(newElems ...Expression) Expressions {
  res := make(Expressions, len(exps))
  copy(res, exps)
  for _, newEl := range newElems {
    unique := true
    for _, exp := range exps {
      if exp.Equals(newEl) {
        unique = false
      }
    }
    if unique {
      res = append(res, newEl)
    }
  }
  return res
}

func (a Expression) Equals(b Expression) bool {
  if a.etype != b.etype {
    return false
  }
  if (a.left == nil && b.left != nil) || (a.left != nil && b.left == nil) ||
     (a.right == nil && b.right != nil) || (a.right != nil && b.right == nil) {
    return false
  }
  if (a.left != nil && b.left != nil && !a.left.Equals(*b.left)) ||
     (a.right != nil && b.right != nil && !a.right.Equals(*b.right)) ||
     (a.mid != b.mid) {
    return false
  }
  return true
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
  for i, val := range ant {
    if i == 0 {
      res += val.Printexp()
    } else {
      res = res + ", " + val.Printexp()
    }
  }
  res = res + " => "
  for i, val := range con {
    if i == 0 {
      res += val.Printexp()
    } else {
      res = res + ", " + val.Printexp()
    }
  }
  return
}

func (seqs Sequents) Printseqs() (res string) {
  for _, val := range seqs {
    res = res + val.Printseq() + "     "
  }
  return
}
