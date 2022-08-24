package SortsSrc

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type IComparator interface {
	Compare(a, b []rune, column int, separator rune) bool
}

type NumericComparator struct{} //compares 2 strings in numeric format {(>) => true}

func (c *NumericComparator) Compare(a, b []rune, column int, separator rune) bool {

	strA, strB := strings.Split(string(a), string(separator)), strings.Split(string(b), string(separator))
	intA, err := strconv.ParseInt(strA[column-1], 10, 64)
	if err != nil {
		fmt.Println("Err! Not found a number")
		return false
	}

	intB, err := strconv.ParseInt(strB[column-1], 10, 64)
	if err != nil {
		fmt.Println("Err! Not found a number")
		return true
	}

	if intA < intB {
		return false
	}

	return true
}

type StringComparator struct{} //compares 2 strings in string format {(>) => true}

func (c *StringComparator) Compare(a, b []rune, column int, separator rune) bool {

	strA, strB := strings.Split(string(a), string(separator)), strings.Split(string(b), string(separator))
	if strings.Compare(strA[column-1], strB[column-1]) < 0 {
		return false
	}

	return true
}

type Sorter struct {
	Bytes      []byte
	Column     int
	Args       map[string]bool
	Comparator IComparator
}

func (s *Sorter) sort() []string {

	tmp := strings.Split(string(s.Bytes), "\n")

	runes := make([][]rune, len(tmp))
	for i, v := range tmp {
		if v[len(v)-1] == '\r' {
			v = v[:len(v)-1]
		}

		runes[i] = []rune(v)
	}

	swapper := reflect.Swapper(runes)
	for i := 0; i < len(runes); i++ {
		for j := i + 1; j < len(runes); j++ {
			if s.Comparator.Compare(runes[i], runes[j], s.Column, ' ') {
				swapper(i, j)

			}
		}
	}

	result := make([]string, len(runes))
	for i, str := range runes {
		result[i] = string(str)
	}
	return result
}

func deleteRepeated(data []string) []string {
	tmp := make(map[string]uint32, len(data))
	for _, v := range data {
		tmp[v] = tmp[v] + 1
	}
	newData := make([]string, 0, len(tmp))
	for _, v := range data {
		tmp[v]--
		if tmp[v] == 0 {
			newData = append(newData, v)
		}
	}
	return newData
}

func reverse(data []string) []string {
	swapper := reflect.Swapper(data)
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		swapper(i, j)
	}
	return data
}

func (s *Sorter) prepare() []func([]string) []string {
	var execList []func([]string) []string
	if ok, _ := s.Args["n"]; ok {
		s.Comparator = &NumericComparator{}
	} else {
		s.Comparator = &StringComparator{}
	}

	if ok, _ := s.Args["u"]; ok {
		execList = append(execList, deleteRepeated)
	}
	if ok, _ := s.Args["r"]; ok {
		execList = append(execList, reverse)
	}
	return execList
}

func (s *Sorter) Run() []string {
	execList := s.prepare()
	result := s.sort()
	for _, exec := range execList {
		result = exec(result)
	}
	return result
}
