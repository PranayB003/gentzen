package main

func Term(a string) *Expression {
  return &Expression{nil, nil, a, 1}
}

func Negation(a *Expression) *Expression {
  return &Expression{nil, a, _not, 2}
}

func Conjunction(a, b *Expression) *Expression {
  return &Expression{a, b, _and, 3}
}

func Disjunction(a, b *Expression) *Expression {
  return &Expression{a, b, _or, 3}
}

func Implication(a, b *Expression) *Expression {
  return &Expression{a, b, _imp, 3}
}

func Dimplication(a, b *Expression) *Expression {
  return &Expression{a, b, _dimp, 3}
}
