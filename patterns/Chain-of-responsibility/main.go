package main

import "fmt"

/*
Цепочка вызовов - используется, когда в разрабатываемой системе имеется группа объектов, которые могут обрабатывать
сообщения определенного типа, и/или все сообщения должны быть обработаны хотя бы одним объектом системы.

Плюсы:
1. Разнесение клиента и обработчиков, уменьшение их зависимости.
2. Реализация принципа единственной ответственности.
Минусы:
1. В каких-то реализациях запрос может быть не обработан.
*/

func main() {
	fmt.Println("Тимлид демонстрирует команде новый амбициозный проект." +
		"- Вот, товарищи бойцы, это новый секретный проект. Стажер, сделай его.")
	pipeline := Trainee{1, &Junior{2, &Middle{4, &Senior{10, nil}}}}
	pipeline.SolveTask(11)
	//pipeline.SolveTask(5)
	//pipeline.SolveTask(9)
}

type IProgrammist interface {
	SolveTask(requiredKnowledge int)
}

type Trainee struct {
	knowledge int
	next      IProgrammist
}

func (p *Trainee) SolveTask(requiredKnowledge int) {
	if requiredKnowledge <= p.knowledge {
		fmt.Println("Стажер осилил!")
	} else {
		fmt.Println("Стажер тужится, пыжится - не может. \n -Джун, помоги ему!")
		p.next.SolveTask(requiredKnowledge)
	}
}

type Junior struct {
	knowledge int
	next      IProgrammist
}

func (p *Junior) SolveTask(requiredKnowledge int) {
	if requiredKnowledge <= p.knowledge {
		fmt.Println("Джун осилил!")
	} else {
		fmt.Println("Джун тужится, пыжится - не может. \n -Мидл, помоги ему!")
		p.next.SolveTask(requiredKnowledge)
	}
}

type Middle struct {
	knowledge int
	next      IProgrammist
}

func (p *Middle) SolveTask(requiredKnowledge int) {
	if requiredKnowledge <= p.knowledge {
		fmt.Println("Мидл осилил!")
	} else {
		fmt.Println("Мидл тужится, пыжится - не может. \n -Сеньор, помоги ему!")
		p.next.SolveTask(requiredKnowledge)
	}
}

type Senior struct {
	knowledge int
	next      IProgrammist
}

func (p *Senior) SolveTask(requiredKnowledge int) {
	if requiredKnowledge <= p.knowledge {
		fmt.Println("Сеньор осилил!")
	} else {
		fmt.Println("Сеньор тужится, пыжится - не может. \nСлишком амбициозный проект(")
	}
}
