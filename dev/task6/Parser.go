package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

/*
	Реализовать утилиту аналог консольной команды cut (man cut).
	Утилита должна принимать строки через STDIN, разбивать по разделителю (TAB) на колонки и выводить запрошенные.
	Реализовать поддержку утилитой следующих ключей:
	-f - "fields" - выбрать поля (колонки)
	-d - "delimiter" - использовать другой разделитель
	-s - "separated" - только строки с разделителем
*/

type Flags struct {
	f []int
	d string
	s bool
}

func ParseFlags() Flags {
	flags := Flags{}
	var tmpF string = ""

	flag.StringVar(&tmpF, "f", "", "fields - выбрать поля (колонки)")
	flag.StringVar(&flags.d, "d", "	", "delimiter - использовать другой разделитель")
	flag.BoolVar(&flags.s, "s", false, "separated - только строки с разделителем")
	flag.Parse()

	if tmpF != "" {
		flags.f = parseFieldsStr(tmpF)
	}

	return flags
}

func parseFieldsStr(str string) []int {
	var res []int = nil
	strs := strings.Split(str, ",")
	for _, s := range strs {
		if num, err := strconv.Atoi(s); err != nil {
			fmt.Println("Can't parse \"", s, "\" as num. Panic!")
			panic(err)
		} else {
			res = append(res, num-1)
		}

	}
	return res
}
