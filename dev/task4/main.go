package main

import (
	"fmt"
	"strings"
)

func main() {
	dictionary := []string{
		"пятак",
		"тяпка",
		"пятка",
		"листок",
		"слиток",
		"столик",
		"ЧерТог",
		"Горечь",
		"КЛОУН",
		"клоун",
		"УКЛОН",
		"КОЛУН",
		"КУЛОН",
	}

	res := findAnagrams(dictionary)
	fmt.Println(res)
}

type anagram struct {
	srcWord string
	wordSet []string
	wordMap map[rune]int
}

func findAnagrams(data []string) map[string][]string {

	anagrams := []anagram{}

	for _, word := range data {
		isAppended := false
		word = strings.ToLower(word)
		for i, anag := range anagrams {
			if isEqual(castToMap(word), anag.wordMap) { //if their letter set equals
				isAppended = true
				if strings.Compare(word, anag.srcWord) != 0 { //if source words are different
					var isUnique = true
					for _, sWord := range anag.wordSet {
						if strings.Compare(word, sWord) == 0 { //if new word is not unique in current set
							isUnique = false
							break
						}
					}
					if isUnique {
						anagrams[i].wordSet = append(anag.wordSet, word)
						//fmt.Println("Append to ", anag.srcWord, " word ", word)
						break
					}
				}
			}
		}
		if !isAppended {
			anagrams = append(anagrams, anagram{word, []string{}, castToMap(word)})
		}
	}

	var resMap = make(map[string][]string)

	for _, anag := range anagrams { //conver type []anagrams to map[string]string[]
		resMap[anag.srcWord] = anag.wordSet
	}

	return resMap
}

func castToMap(word string) map[rune]int {
	wordMap := make(map[rune]int)
	for _, symb := range []rune(word) {
		wordMap[symb] += 1
	}
	return wordMap
}

func isEqual(a, b map[rune]int) bool {

	if len(a) != len(b) {
		return false
	}

	for k, v := range a {
		if v != b[k] {
			return false
		}
	}
	return true
}
