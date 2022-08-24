package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
	Реализовать утилиту аналог консольной команды cut (man cut).
	Утилита должна принимать строки через STDIN, разбивать по разделителю (TAB) на колонки и выводить запрошенные.
	Реализовать поддержку утилитой следующих ключей:
	-f - "fields" - выбрать поля (колонки)
	-d - "delimiter" - использовать другой разделитель
	-s - "separated" - только строки с разделителем


	Программа единожды принимает аргументы командной строки, а затем подтягивает строки из stdin
	и обрабатывает их до тех пор, пока не встретит пустую строку.
	Аргументы -f подаются в виде чисел, раздленных только запятыми.
	Например, "... -f 2,3,5 ..."
*/

func main() {

	myCutter := Cutter{ParseFlags()}
	var scanner = bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		data := scanner.Text()
		err := scanner.Err()
		if err != nil {
			fmt.Println("Err!")
			panic(err)
		}

		if data == "" {
			fmt.Println("Closing...")
			break
		}

		fmt.Println(myCutter.Process(data))
	}
}
