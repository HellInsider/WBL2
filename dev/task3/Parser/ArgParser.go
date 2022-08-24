package Parser

import (
	"flag"
)

func ParseFlags() (*map[string]bool, int) {
	var n, r, u bool
	var column int
	flag.IntVar(&column, "k", 0, "указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)")
	flag.BoolVar(&n, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&r, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&u, "u", false, "не выводить повторяющиеся строки")
	flag.Parse()

	if column < 1 {
		column = 1
	}

	return &map[string]bool{"n": n, "r": r, "u": u}, column
}
