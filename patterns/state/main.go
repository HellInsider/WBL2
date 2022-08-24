package main

import "fmt"

/*
Состояние — это поведенческий паттерн проектирования, меняющий поведение объекта в зависимости от привязки к состоянию.

Плюсы:
1. Избавляет от множества больших условных операторов машины состояний.
2. Поведение может быть целиком описано одним классом.

Минусы:
1. Может создать проблемы, если состояния должны обмениваться данными, или одно состояние настраивает свойства другого.
*/

func main() {
	gate := Gate{&sLower{}}

	gate.Raise()
	gate.Raise()
	gate.Lower()
	gate.Lower()
	gate.LowerButALittle()
	gate.LowerButALittle()
	gate.Raise()
}

type State interface {
	Raise()
	Lower()
	LowerButALittle()
}

type Gate struct {
	state State
}

func (s *Gate) Raise() {
	s.state.Raise()
	s.state = &sUp{}

}

func (s *Gate) Lower() {
	s.state.Lower()
	s.state = &sLower{}
}

func (s *Gate) LowerButALittle() {
	s.state.LowerButALittle()
	s.state = &sLowerButALittle{}
}

type sUp struct {
}

func (s *sUp) Raise() {
	fmt.Println("The gate is already raised")
}

func (s *sUp) Lower() {
	fmt.Println("The gate lowered")
}

func (s *sUp) LowerButALittle() {
	fmt.Println("The gate a little lowered")
}

type sLower struct {
}

func (s *sLower) Raise() {
	fmt.Println("The gate raised")
}

func (s *sLower) Lower() {
	fmt.Println("The gate is already lowered")
}

func (s *sLower) LowerButALittle() {
	fmt.Println("The gate a little rised")
}

type sLowerButALittle struct {
}

func (s *sLowerButALittle) Raise() {
	fmt.Println("The gate raised")
}

func (s *sLowerButALittle) Lower() {
	fmt.Println("The gate lowered")
}

func (s *sLowerButALittle) LowerButALittle() {
	fmt.Println("The gate is already a bit lowered")
}
