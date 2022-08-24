package main

import (
	"bufio"
	"fmt"
	"os"
	"task8/Cmds"
)

/*
	- cd <args> - смена директории (в качестве аргумента могут
	быть то-то и то)
	- pwd - показать путь до текущего каталога
	- echo <args> - вывод аргумента в STDOUT
	- kill <args> - "убить" процесс, переданный в качесте
	аргумента (пример: такой-то пример)
	- ps - выводит общую информацию по запущенным процессам в
	формате *такой-то формат*


	Пример ввода: -echo 123 321 44 | -pwd | -cd ../task9

*/

func main() {
	var scanner = bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		data := scanner.Text()
		err := scanner.Err()
		if err != nil {
			fmt.Println("Err!")
			panic(err)
		}
		Cmds.ParseLine(data)
	}
}
