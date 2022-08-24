package main

import (
	"fmt"
	"testing"
	"unicode"
)

func decode(data string) string {
	str := []rune(data)
	var resStr []rune
	var savedSymb rune = 0
	for _, symb := range str {

		if unicode.IsDigit(symb) {

			if savedSymb == 0 {
				fmt.Println("Here is error!")
				return ""
			}

			for i := 0; i < int(symb-'0'); i++ {
				resStr = append(resStr, savedSymb)
			}
			savedSymb = 0

		} else {
			if savedSymb != 0 {
				resStr = append(resStr, savedSymb)
			}
			savedSymb = symb
		}
	}

	if savedSymb != 0 {
		resStr = append(resStr, savedSymb)
	}

	fmt.Println(string(resStr))
	return string(resStr)
}

type decodingTestStruct struct {
	inputStr    string
	expectedStr string
}

func testDecode(t *testing.T) {
	testingStrings := []decodingTestStruct{
		{
			"a4bc2d5e",
			"aaaabccddddde",
		},
		{
			"abcd",
			"abcd",
		},
		{
			"45",
			"",
		},
		{
			"",
			"",
		},
		{
			"abc3d",
			"abcccd",
		},
	}
	for _, testItem := range testingStrings {
		s := decode(testItem.inputStr)
		if s != testItem.expectedStr {
			t.Errorf("Error occured on string : %s", testItem.inputStr)
		}
	}
}
