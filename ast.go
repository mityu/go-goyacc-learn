package main

import (
	"fmt"
	"strconv"
)

type Expr interface {
	Eval() (Expr, error)
}

type ExprNum struct {
	literal string
}

func (e ExprNum) Eval() (Expr, error) {
	return e, nil
}

func (e ExprNum) Value() (int, error) {
	return strconv.Atoi(e.literal)
}

type ExprBinary struct {
	lhs, rhs Expr
	op       string
}

func (e ExprBinary) Eval() (Expr, error) {
	lhs, err := e.lhs.Eval()
	if err != nil {
		return nil, err
	}

	rhs, err := e.rhs.Eval()
	if err != nil {
		return nil, err
	}

	switch e.op {
	case "+":
		lhsv, err := lhs.(ExprNum).Value()
		if err != nil {
			return nil, err
		}
		rhsv, err := rhs.(ExprNum).Value()
		if err != nil {
			return nil, err
		}
		v := lhsv + rhsv
		return ExprNum{literal: fmt.Sprintf("%d", v)}, nil
	}
	panic("unreachable")
}
