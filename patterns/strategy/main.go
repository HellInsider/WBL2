package main

import "fmt"

/*
Стратегия  — поведенческий шаблон проектирования, предназначенный для определения семейства алгоритмов, инкапсуляции каждого
из них и обеспечения их взаимозаменяемости. Это позволяет выбирать алгоритм путём определения соответствующего класса.
Шаблон Strategy позволяет менять выбранный алгоритм независимо от объектов-клиентов, которые его используют.

Плюсы:
1. Возможность замены алгоритмов в рантайме
2. Отделение алгоритмов от остальной логики, сокрытие самих алгоритмов
Минусы:
1. Усложнение кода, засчет введения дополнительных объектов
2. Единый интерфейс
*/

func main() {
	machine := ResearchMachine{&sLake{"2 fish"}}
	machine.Research()

	machine.SetStrategy(&sDesert{"123 scorpions"})
	machine.Research()

	machine.SetStrategy(&sForest{"1 fox"})
	machine.Research()

}

type ResearchMachine struct {
	strategy Srategy
}

func (m *ResearchMachine) SetStrategy(srategy Srategy) {
	m.strategy = srategy
}

func (m *ResearchMachine) Research() {
	fmt.Println(m.strategy.SearchLive())
}

type Srategy interface {
	SearchLive() string
}

type sLake struct {
	animals string
}

func (s *sLake) SearchLive() string {
	return "After searching in lake you found " + s.animals
}

type sDesert struct {
	animals string
}

func (s *sDesert) SearchLive() string {
	return "After searching in desert you found " + s.animals
}

type sForest struct {
	animals string
}

func (s *sForest) SearchLive() string {
	return "After searching in forest you found " + s.animals
}
