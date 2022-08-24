package main

import "fmt"

/*
Посетитель — поведенческий шаблон проектирования, описывающий операцию, которая выполняется над объектами других классов.
При изменении visitor нет необходимости изменять обслуживаемые классы.

Плюсы:
1. Упрощается добавление новых операций.
2. Класс Visitor может запоминать в себе какое-то состояние по мере обхода контейнера.
Минусы:
1. Затрудняет добавление новых классов, поскольку нужно обновлять иерархию посетителя и его сыновей.
*/

func main() {

	sc := SoundChecker{}
	ms := Musicians{}

	tr := Trumpet{38}
	dr := Drums{1}

	tr.accept(&sc)
	dr.accept(&sc)
	tr.accept(&ms)
	dr.accept(&ms)

}

type IVisitor interface {
	VisitDrums(d *Drums)
	VisitTrumpet(d *Trumpet)
}

type SoundChecker struct {
}

func (v *SoundChecker) VisitDrums(d *Drums) {
	fmt.Println("SoundChecker checking ", d.getName(), " Condition: ", d.condition)
	d.condition = 100
}

func (v *SoundChecker) VisitTrumpet(d *Trumpet) {
	fmt.Println("SoundChecker checking ", d.getName(), ". Condition: ", d.condition)
	d.condition = 100
}

type Musicians struct {
}

func (v *Musicians) VisitDrums(d *Drums) {
	d.condition = 37
	fmt.Println("Musicians use ", d.getName(), ". Condition: ", d.condition)

}

func (v *Musicians) VisitTrumpet(d *Trumpet) {
	d.condition = 50
	fmt.Println("Musicians use ", d.getName(), ". Condition: ", d.condition)

}

type Trumpet struct {
	condition int
}

func (d *Trumpet) getName() string {
	return "Trumpet"
}

func (d *Trumpet) accept(v IVisitor) {
	v.VisitTrumpet(d)
}

type Drums struct {
	condition int
}

func (d *Drums) getName() string {
	return "Drums"
}

func (d *Drums) accept(v IVisitor) {
	v.VisitDrums(d)
}
