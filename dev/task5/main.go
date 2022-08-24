package main

import (
	"fmt"
	"strings"
	"task5/GrepSrc"
	"task5/Parser"
)

/*
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", напечатать номер строки
*/
func main() {
	myGrep := GrepSrc.NewGrep(Parser.ParseFlags())
	fmt.Printf(strings.Join(myGrep.Run(), "\n"))
}
