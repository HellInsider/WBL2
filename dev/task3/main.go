package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"task3/Parser"
	"task3/SortsSrc"
)

/*
	My solution can be described in this axioms:
	1)	Simple 1 column sorting - is a subset of a table sorting. So I need make table sorting as base.
	2)	Some flags can be processed after sorting (f. e. -r and -u). I can execute them in the end.
	3)	Strings sorting and numbers sorting needs different comparison methods. I need different funcs for it.

	-k _ : указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)
	-n : сортировать по числовому значению
	-r : сортировать в обратном порядке
	-u : не выводить повторяющиеся строки
*/

func main() {
	flags, k := Parser.ParseFlags()
	bytes, ok := os.ReadFile(os.Args[len(os.Args)-1])
	if ok != nil {
		log.Fatal(ok)
	}
	mySort := SortsSrc.Sorter{Bytes: bytes, Column: k, Args: *flags}
	result := mySort.Run()
	fmt.Println(strings.Join(result, "\n"))
	Parser.WriteFile(result, "sortedData.txt")
}
