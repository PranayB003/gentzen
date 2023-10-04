package main

// connectives
var _true  string = "TRUE"
var _false string = "FALSE"
var _not   string = "!"
var _and   string = "&"
var _or    string = "|"
var _imp   string = "->"
var _dimp  string = "<->"

// etypes for Expressions
var _term   byte = 1
var _unary  byte = 2
var _binary byte = 3

type Expression struct {
  left  *Expression
  right *Expression
  mid    string
  etype  byte
}

type Expressions []Expression

type Sequent struct {
  ant Expressions // antecedents
  con Expressions // consequents
}

type Sequents []Sequent
