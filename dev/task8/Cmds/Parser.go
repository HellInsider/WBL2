package Cmds

import (
	"fmt"
	"strings"
)

func ParseLine(str string) {

	mapCmds := map[string]ICommand{
		"-kill": NewKILL(),
		"-echo": NewECHO(),
		"-cd":   NewCD(),
		"-ps":   NewPS(),
		"-pwd":  NewPWD(),
		"/quit": NewQUIT(),
	}

	commands := strings.Split(str, "|")
	var err error

	for _, com := range commands {
		err = func(str string) error {
			str = strings.TrimSpace(str)
			words := strings.Split(str, " ")
			args := words[1:]
			executor := mapCmds[words[0]]
			if executor == nil {
				fmt.Println("Unknown command \"", words[0], "\".")
				return nil
			} else {
				return executor.Exec(args)
			}
		}(com)

		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
