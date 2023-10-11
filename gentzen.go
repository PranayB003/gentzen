package main

import (
  "fmt"
)

// checks for validity of an expression`
func (exp Expression) Proove() bool {
  var seqs, queue = make(Sequents, 0), make(Sequents, 0)
  queue = append(queue, Sequent{Expressions{}, Expressions{exp}})

  if !_optMap["--validity-only"] {
    fmt.Println("Proof Tree:")
    fmt.Println(queue.Printseqs())
  }

  for curLvl, nxtLvl := 1, 0; len(queue) > 0; {
    top := queue[0]
    if len(queue) == 1 {
      queue = []Sequent{}
    } else {
      queue = queue[1:]
    }
    curLvl-- 

    if top.IsLowestForm() == false {
      decSeqs := top.Decompose() 
      queue = append(queue, decSeqs...)
      nxtLvl += len(decSeqs)
    } else {
      seqs = append(seqs, top)
    }

    if !_optMap["--validity-only"] {
      if curLvl == 0 {
        fmt.Println(queue.Printseqs())
        curLvl, nxtLvl = nxtLvl, 0
      }
    }
  }

  if !_optMap["--validity-only"] {
    fmt.Println("Final Sequents:")
    fmt.Println(seqs.Printseqs(), "\n")
  }

  res := true
  for i := 0; i < len(seqs); i++ {
    res = res && seqs[i].HasContradiction()
  }

  return res
}

// decompose a sequent by using one of the 8 rules
func (seq Sequent) Decompose() Sequents {
  ant, con := seq.ant, seq.con

  // right rules without fanout (right-NOT, right-OR, right-IMPLICATION)
  for i, exp := range con {
    newCon := con[:i].Append(con[i+1:]...)
    if exp.etype == 2 && exp.mid == _not {
      return Sequents{Sequent{ant.AppendUnique(*(exp.right)), newCon}}
    } else if exp.etype == 3 && exp.mid == _or {
      return Sequents{Sequent{ant, newCon.AppendUnique(*(exp.left), *(exp.right))}}
    } else if exp.etype == 3 && exp.mid == _imp {
      newCon1 := newCon.AppendUnique(*(exp.right))
      newAnt := ant.AppendUnique(*(exp.left))
      return Sequents{Sequent{newAnt, newCon1}}
    }
  }

  // left rules without fanout (left-NOT, left-AND)
  for i, exp := range ant {
    newAnt := ant[:i].Append(ant[i+1:]...)
    if exp.etype == 2 && exp.mid == _not {
      return Sequents{Sequent{newAnt, con.AppendUnique(*(exp.right))}}
    } else if exp.etype == 3 && exp.mid == _and {
      return Sequents{Sequent{newAnt.AppendUnique(*(exp.left), *(exp.right)), con}}
    }
  }

  // right rules with fanout (right-AND)
  for i, exp := range con {
    newCon := con[:i].Append(con[i+1:]...)
    if exp.etype == 3 && exp.mid == _and {
      newCon1 := newCon.AppendUnique(*(exp.left))
      newCon2 := newCon.AppendUnique(*(exp.right))
      return Sequents{Sequent{ant, newCon1}, Sequent{ant, newCon2}}
    }  
  }

  // left rules with fanout (left-OR, left-IMPLICATION)
  for i, exp := range ant {
    newAnt := ant[:i].Append(ant[i+1:]...)
    if exp.etype == 3 && exp.mid == _or {
      newAnt1 := newAnt.AppendUnique(*(exp.left))
      newAnt2 := newAnt.AppendUnique(*(exp.right))
      return Sequents{Sequent{newAnt1, con}, Sequent{newAnt2, con}}
    } else if exp.etype == 3 && exp.mid == _imp {
      newAnt1 := newAnt.AppendUnique(*(exp.right))
      newCon2 := con.AppendUnique(*(exp.left))
      return Sequents{Sequent{newAnt, newCon2}, Sequent{newAnt1, con}}
    }
  }
  
  return Sequents{}
}

// returns true if the sequent has a contradiction
// the sequent must be in the lowest form (not decomposable further)
func (seq Sequent) HasContradiction() (res bool) {
  res = false
  ant, con := seq.ant, seq.con
  antTerms := make(map[string]bool)

  for i := 0; i < len(ant); i++ {
    if ant[i].mid == _false {
      res = true
      return
    }
    antTerms[ant[i].mid] = true
  }

  for i := 0; i < len(con); i++ {
    if antTerms[con[i].mid] == true || con[i].mid == _true {
      res = true
      return
    }
  }

  return 
}

// returns true if a sequent cannot be decomposed further
func (seq Sequent) IsLowestForm() (res bool) {
  res = true
  ant, con := seq.ant, seq.con

  for i := 0; i < len(ant); i++ {
    if ant[i].etype != 1 {
      res = false
      return
    }
  }

  for i := 0; i < len(con); i++ {
    if con[i].etype != 1 {
      res = false
      return
    }
  }

  return 
}
