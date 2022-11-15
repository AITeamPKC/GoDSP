package main

import (
	"fmt"
	"strings"
)

type ExpressionVisitor interface {
	VisitDoubleExpression(de *DoubleExpression)
	VisitAdditionExpression(ae *AdditionExpression)
}

type Expression interface {
	Accept(ev ExpressionVisitor)
}
type DoubleExpression struct {
	value float64
}

type AdditionExpression struct {
	left, right Expression
}

func (d *DoubleExpression) Accept(ev ExpressionVisitor) {
	ev.VisitDoubleExpression(d)
}

type AdditonExpression struct {
	left, right Expression
}

func (a *AdditionExpression) Accept(ev ExpressionVisitor) {
	ev.VisitAdditionExpression(a)
}

type ExpressPrinter struct {
	sb strings.Builder
}

func (e *ExpressPrinter) VisitDoubleExpression(de *DoubleExpression) {
	e.sb.WriteString(fmt.Sprintf("%g", de.value))
}

func (e *ExpressPrinter) VisitAdditionExpression(ae *AdditionExpression) {
	e.sb.WriteString("(")
	ae.left.Accept(e)
	e.sb.WriteString("+")
	ae.right.Accept(e)
	e.sb.WriteString(")")
}

func NewExpressionPrinter() *ExpressPrinter {
	return &ExpressPrinter{strings.Builder{}}
}

func (e *ExpressPrinter) String() string {
	return e.sb.String()
}

func main() {
	e := &AdditionExpression{
		&DoubleExpression{1},
		&AdditionExpression{
			left:  &DoubleExpression{2},
			right: &DoubleExpression{3},
		},
	}

	ep := NewExpressionPrinter()
	ep.VisitAdditionExpression(e)
	fmt.Println(ep.String())
}
