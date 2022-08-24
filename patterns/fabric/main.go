package main

import (
	"fmt"
	"math/rand"
)

/*

 Фабрика - порождающий шаблон проектирования, делегирующий создание объектов наследникам родительского класса.
В момент создания наследники могут определить, какой класс создавать.

Плюсы:
1. Позволяет использовать в коде программы не конкретные классы, а манипулировать абстрактными объектами на более высоком уровне.
2. Упрощает добавление новых продуктов в программу.
*/

func main() {
	fabric := Fabric{}
	myCreature := fabric.Create()
	myCreature.Info()
	myCreature = fabric.Create()
	myCreature.Info()

}

type Fabric struct {
}

func (c *Fabric) Create() Creature {
	creature := Creature{}

	if rand.Int()%2 == 1 {
		creature.features = append(creature.features, &Swimming{})
	}
	if rand.Int()%2 == 1 {
		creature.features = append(creature.features, &WithLegs{})
	}

	if rand.Int()%2 == 1 {
		creature.features = append(creature.features, &Predator{})
	} else {
		creature.features = append(creature.features, &Herbivorous{})
	}
	return creature
}

type Creature struct {
	features []IFeature
}

func (c *Creature) Info() {
	fmt.Println("My creature is...")
	for _, f := range c.features {
		f.MakeAction()
	}
}

type IFeature interface {
	MakeAction()
}

type Swimming struct {
}

func (c *Swimming) MakeAction() {
	fmt.Println("Swimming")
}

type WithLegs struct {
}

func (c *WithLegs) MakeAction() {
	fmt.Println("Walking")
}

type Predator struct {
}

func (c *Predator) MakeAction() {
	fmt.Println("Predator")
}

type Herbivorous struct {
}

func (c *Herbivorous) MakeAction() {
	fmt.Println("Herbivorous")
}
