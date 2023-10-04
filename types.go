package main

// operator types
var _not  string = "!"
var _and  string = "&"
var _or   string = "|"
var _imp  string = "->"
var _dimp string = "<->"

type Expression struct {
  left  *Expression
  right *Expression
  mid    string
  etype  byte
}

// etypes for Expressions
var _term   byte = 1
var _unary  byte = 2
var _binary byte = 3
