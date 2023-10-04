package main

import (
  "fmt"
)

// checks for validity of an expression`
func (exp Expression) Proove() bool {
  var seqs, queue = make([]Sequent, 0), make([]Sequent, 0)
  queue = append(queue, Sequent{[]Expression{}, []Expression{exp}})

  fmt.Println("Sequents:")
  fmt.Println(Printseqs(seqs))

  for len(queue) > 0 {
    top := queue[0]
    if len(queue) == 1 {
      queue = []Sequent{}
    } else {
      queue = queue[1:]
    }

    if top.isLowestForm() == false {
      decSeqs := top.decompose() 
      queue = append(queue, decSeqs...)
    } else {
      seqs = append(seqs, top)
    }

    fmt.Println("Sequents:")
    fmt.Println(Printseqs(seqs))
  }

  res := true
  for i := 0; i < len(seqs); i++ {
    res = seqs[i].validate()
  }

  return res
}

// decompose a sequent by using one of the 8 rules
func (seq Sequent) decompose() []Sequent {
  res := make([]Sequent, 0)
  ant, con := seq.ant, seq.con

  // left rules
  for i := 0; i < len(ant); i++ {
    exp := ant[i]
    if exp.etype == 2 && exp.mid == _not {
      newAnt := append(ant[:i], ant[i+1:]...)
      newCon := append(con, *(exp.right))
      res = append(res, Sequent{newAnt, newCon})
    } else if exp.etype == 3 && exp.mid == _and {
      newAnt := append(ant[:i], ant[i+1:]...)
      newAnt = append(newAnt, *(exp.left), *(exp.right))
      res = append(res, Sequent{newAnt, con})
    } else if exp.etype == 3 && exp.mid == _or {
      newAnt := append(ant[:i], ant[i+1:]...)
      newAnt1 := append(newAnt, *(exp.left))
      newAnt2 := append(newAnt, *(exp.right))
      res = append(res, Sequent{newAnt1, con}, Sequent{newAnt2, con})
    } else if exp.etype == 3 && exp.mid == _imp {
      newAnt := append(ant[:i], ant[i+1:]...)
      newAnt1 := append(newAnt, *(exp.right))
      newCon2 := append(con, *(exp.left))
      res = append(res, Sequent{newAnt, newCon2}, Sequent{newAnt1, con})
    }
  }

  // right rules
  for i := 0; i < len(con); i++ {
    exp := con[i]
    if exp.etype == 2 && exp.mid == _not {
      newCon := append(con[:i], con[i+1:]...)
      newAnt := append(ant, *(exp.right))
      res = append(res, Sequent{newAnt, newCon})
    } else if exp.etype == 3 && exp.mid == _or {
      newCon := append(con[:i], con[i+1:]...)
      newCon = append(newCon, *(exp.left), *(exp.right))
      res = append(res, Sequent{ant, newCon})
    } else if exp.etype == 3 && exp.mid == _and {
      newCon := append(con[:i], con[i+1:]...)
      newCon1 := append(newCon, *(exp.left))
      newCon2 := append(newCon, *(exp.right))
      res = append(res, Sequent{ant, newCon1}, Sequent{ant, newCon2})
    } else if exp.etype == 3 && exp.mid == _imp {
      newCon := append(con[:i], con[i+1:]...)
      newCon = append(newCon, *(exp.right))
      newAnt := append(ant, *(exp.left))
      res = append(res, Sequent{newAnt, newCon})
    }
  }

  return res
}

// returns false if the sequent has a contradiction
// the sequent must be in the lowest form (not decomposable further)
func (seq Sequent) validate() bool {
  res := true
  ant, con := seq.ant, seq.con
  antTerms := make(map[string]bool)

  for i := 0; i < len(ant); i++ {
    exp := ant[i]
    antTerms[exp.mid] = true
  }

  for i := 0; i < len(con); i++ {
    exp := con[i]
    if antTerms[exp.mid] == true {
      res = false
      break
    }
  }

  return res
}

// returns true if a sequent cannot be decomposed further
func (seq Sequent) isLowestForm() bool {
  res := true
  ant, con := seq.ant, seq.con

  for i := 0; i < len(ant); i++ {
    if ant[i].etype != 1 {
      res = false
      break
    }
  }

  for i := 0; i < len(con); i++ {
    if con[i].etype != 1 {
      res = false
      break
    }
  }

  return res
}

func (seq Sequent) Printseq() (res string) {
  ant, con := seq.ant, seq.con
  for i := 0; i < len(ant); i++ {
    res = res + ", " + ant[i].Printexp()
  }
  res = res + " => "
  for i := 0; i < len(con); i++ {
    res = res + ", " + con[i].Printexp()
  }
  return
}

func Printseqs(seqs []Sequent) (res string) {
  for i := 0; i < len(seqs); i++ {
    res = res + seqs[i].Printseq() + "\n"
  }
  return
}
