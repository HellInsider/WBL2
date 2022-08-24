package main

import "fmt"

/*
Строитель -  порождающий шаблон проектирования предоставляет способ создания составного объекта.
Содан чтобы отделять конструирование сложного объекта от его представления. В результате чего,
после действий одного и того же процесса конструирования могут получаться разные представления.

Плюсы:
1. Позволяет изменять внутреннее представление продукта;
2. Изолирует код, реализующий конструирование и представление;
3. Дает более тонкий контроль и вариативность над процессом конструирования.
Минусы:
1. Алгоритм создания сложного объекта не должен зависеть от того, из каких частей состоит объект и как они стыкуются между собой;
2. Процесс конструирования должен обеспечивать различные представления конструируемого объекта.
*/

func main() {
	var b1 DormBuilder
	b1.BuildDoor("plastic")
	b1.BuildDoor("wood")
	b1.BuildRoof("plastic bag")
	b1.BuildWall("wood")
	b1.BuildWall("wood")
	b1.BuildWall("wood")
	b1.BuildWindow("glass")

	dorm := b1.Build()
	dorm.Info()

	var b2 PenthouseBuilder
	b2.BuildDoor("steel")
	b2.BuildDoor("gold")
	b2.BuildRoof("glass")
	b2.BuildWall("hedge")
	b2.BuildWall("hedge")
	b2.BuildWall("hedge")
	b2.BuildWindow("glass and gold")

	pent := b2.Build()
	pent.Info()
}

type Penthouse struct {
	walls   []Wall
	windows []Window
	doors   []Door
	roof    Roof
}

func (h *Penthouse) Info() {
	fmt.Printf("Penthouse info:")
	for _, i := range h.walls {
		fmt.Println("wall of ", i.details)
	}

	for _, i := range h.windows {
		fmt.Println("window of ", i.details)
	}

	for _, i := range h.doors {
		fmt.Println("door of ", i.details)
	}

	fmt.Println("roof of ", h.roof.details)
}

type Dorm struct {
	walls   []Wall
	windows []Window
	doors   []Door
	roof    Roof
}

func (h *Dorm) Info() {
	fmt.Printf("Dorm info:")
	for _, i := range h.walls {
		fmt.Println("wall of ", i.details)
	}

	for _, i := range h.windows {
		fmt.Println("window of ", i.details)
	}

	for _, i := range h.doors {
		fmt.Println("door of ", i.details)
	}

	fmt.Println("roof of ", h.roof.details)
}

type Wall struct {
	details string
}
type Window struct {
	details string
}

type Door struct {
	details string
}

type Roof struct {
	details string
}
