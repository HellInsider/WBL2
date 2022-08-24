package main

import (
	"fmt"
	"math/rand"
)

/*
	Команда - паттерн оборачивающий запросы в объекты,
	тем самым позволяя передавать их как аргументы при вызове методов,ставить запросы в очередь,
	логировать их, а также поддерживать отмену операций.

Плюсы:
1. Позволяет иметь эффективное представление запросов к некоторой системе,
	не обладая при этом знаниями ни об их природе ни о способах их обработки.
Минусы:
1. Усложняет код из-за необходимости реализации дополнительных классов
*/

func main() {
	a := AAACommand{}
	b := BBBCommand{}
	h := History{}

	for i := 0; i < 10; i++ {
		var cmd ICommand
		if rand.Int()%2 == 0 {
			cmd = &a
		} else {
			cmd = &b
		}
		h.history = append(h.history, cmd)
	}

	h.Execute()

}

type ICommand interface {
	getName() string
	Execute()
}

type AAACommand struct {
}

func (c *AAACommand) getName() string {
	return "AAACommand"
}

func (c *AAACommand) Execute() {
	fmt.Println("AAA")
}

type BBBCommand struct {
}

func (c *BBBCommand) getName() string {
	return "BBBCommand"
}

func (c *BBBCommand) Execute() {
	fmt.Println("BBB")
}

type History struct {
	history []ICommand
}

func (c *History) Execute() {
	fmt.Println("Command history is: ")
	for _, cmd := range c.history {
		fmt.Println(cmd.getName())
	}
}
