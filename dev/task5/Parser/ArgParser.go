package Parser

import (
	"flag"
	"fmt"
	"io"
	"os"
)

type Flags struct {
	A         int  // "after" печатать +N строк после совпадения
	B         int  // "before" печатать +N строк до совпадения
	C         int  // "context" (A+B) печатать ±N строк вокруг совпадения
	OnlyCount bool // "count" (только количество строк)
	Ignore    bool // "ignore-case" (игнорировать регистр)
	Invert    bool // "invert" (вместо совпадения, исключать)
	Fixed     bool // "fixed", точное совпадение со строкой, не паттерн
	StrNum    bool // "line num", напечатать номер строки
	Input     io.Reader
	SearchStr string
}

func ParseFlags() Flags {
	flags := Flags{}
	flag.IntVar(&flags.A, "A", 0, "after - печатать +N строк после совпадения")
	flag.IntVar(&flags.B, "B", 0, "before - печатать +N строк до совпадения")
	flag.IntVar(&flags.C, "C", 0, "context - (A+B) печатать ±N строк вокруг совпадения")
	flag.BoolVar(&flags.OnlyCount, "c", false, "count - (количество строк)")
	flag.BoolVar(&flags.Ignore, "i", false, "ignore-case - (игнорировать регистр)")
	flag.BoolVar(&flags.Invert, "v", false, "invert - (вместо совпадения, исключать)")
	flag.BoolVar(&flags.Fixed, "F", false, "fixed - точное совпадение со строкой, не паттерн")
	flag.BoolVar(&flags.StrNum, "n", false, "line num - напечатать номер строки")
	flag.Parse()

	if _, err := os.Stat(os.Args[len(os.Args)-1]); os.IsNotExist(err) {
		flags.SearchStr = os.Args[len(os.Args)-1]
		flags.Input = os.Stdin
	} else {
		flags.Input, err = os.Open(os.Args[len(os.Args)-1])
		if err != nil {
			fmt.Println("File \"", os.Args[len(os.Args)-1], "\" not found! Panic!\n At the disco")
			panic(err)
		}
		flags.SearchStr = os.Args[len(os.Args)-2]
	}
	flags.correctnessCheck()
	return flags
}

func (f *Flags) correctnessCheck() {
	if f.C > 0 {
		f.A, f.B = f.C, f.C
	}
}
