package main

import (
  "fmt"
)

// checks for validity of an expression`
func (exp Expression) Proove() bool {
  var seqs, queue = make(Sequents, 0), make(Sequents, 0)
  queue = append(queue, Sequent{Expressions{}, Expressions{exp}})

  fmt.Println("\nProof Tree:")
  fmt.Println(queue.Printseqs())

  for curLvl, nxtLvl := 1, 0; len(queue) > 0; {
    top := queue[0]
    if len(queue) == 1 {
      queue = []Sequent{}
    } else {
      queue = queue[1:]
    }
    curLvl-- 

    if top.isLowestForm() == false {
      decSeqs := top.decompose() 
      queue = append(queue, decSeqs...)
      nxtLvl += len(decSeqs)
    } else {
      seqs = append(seqs, top)
    }

    if curLvl == 0 {
      fmt.Println(queue.Printseqs())
      curLvl, nxtLvl = nxtLvl, 0
    }
  }

  fmt.Println("Final Sequents:")
  fmt.Println(seqs.Printseqs())

  res := true
  for i := 0; i < len(seqs); i++ {
    res = seqs[i].validate()
  }

  return res
}

// decompose a sequent by using one of the 8 rules
func (seq Sequent) decompose() Sequents {
  res := make(Sequents, 0)
  ant, con := seq.ant, seq.con

  // left rules
  for i := 0; i < len(ant); i++ {
    exp := ant[i]
    newAnt := ant[:i].Append(ant[i+1:]...)
    if exp.etype == 2 && exp.mid == _not {
      res = append(res, Sequent{newAnt, con.Append(*(exp.right))})
    } else if exp.etype == 3 && exp.mid == _and {
      res = append(res, Sequent{newAnt.Append(*(exp.left), *(exp.right)), con})
    } else if exp.etype == 3 && exp.mid == _or {
      newAnt1 := newAnt.Append(*(exp.left))
      newAnt2 := newAnt.Append(*(exp.right))
      res = append(res, Sequent{newAnt1, con}, Sequent{newAnt2, con})
    } else if exp.etype == 3 && exp.mid == _imp {
      newAnt1 := newAnt.Append(*(exp.right))
      newCon2 := con.Append(*(exp.left))
      res = append(res, Sequent{newAnt, newCon2}, Sequent{newAnt1, con})
    }
  }

  // right rules
  for i := 0; i < len(con); i++ {
    exp := con[i]
    newCon := con[:i].Append(con[i+1:]...)
    if exp.etype == 2 && exp.mid == _not {
      res = append(res, Sequent{ant.Append(*(exp.right)), newCon})
    } else if exp.etype == 3 && exp.mid == _or {
      res = append(res, Sequent{ant, newCon.Append(*(exp.left), *(exp.right))})
    } else if exp.etype == 3 && exp.mid == _and {
      newCon1 := newCon.Append(*(exp.left))
      newCon2 := newCon.Append(*(exp.right))
      res = append(res, Sequent{ant, newCon1}, Sequent{ant, newCon2})
    } else if exp.etype == 3 && exp.mid == _imp {
      newCon1 := newCon.Append(*(exp.right))
      newAnt := ant.Append(*(exp.left))
      res = append(res, Sequent{newAnt, newCon1})
    }
  }

  return res
}

// returns false if the sequent has a contradiction
// the sequent must be in the lowest form (not decomposable further)
func (seq Sequent) validate() (res bool) {
  res = true
  ant, con := seq.ant, seq.con
  antTerms := make(map[string]bool)

  for i := 0; i < len(ant); i++ {
    if ant[i].mid == _false {
      res = false
      return
    }
    antTerms[ant[i].mid] = true
  }

  for i := 0; i < len(con); i++ {
    if antTerms[con[i].mid] == true || con[i].mid == _true {
      res = false
      return
    }
  }

  return 
}

// returns true if a sequent cannot be decomposed further
func (seq Sequent) isLowestForm() (res bool) {
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
