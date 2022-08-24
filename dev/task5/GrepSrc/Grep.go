package GrepSrc

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"task5/Parser"
)

type Grep struct {
	data           []string
	searchStr      string
	resultStrIndxs []int
	flags          Parser.Flags
}

func (g *Grep) compareString(s string) bool {
	if g.flags.Fixed {
		if strings.Compare(g.searchStr, s) == 0 {
			return true
		}
	} else {
		if strings.Contains(s, g.searchStr) {
			return true
		}
	}
	return false
}

func (g *Grep) contains(ind int) bool {
	for _, v := range g.resultStrIndxs {
		if v == ind {
			return true
		}
	}
	return false
}

func (g *Grep) addStrings(ind int) {
	var min, max int
	if g.flags.C != 0 {
		min, max = ind-g.flags.C, ind+g.flags.C
	} else {
		min, max = ind-g.flags.B, ind+g.flags.A
	}

	if min < 0 {
		min = 0
	}

	if max > len(g.data)-1 {
		max = len(g.data) - 1
	}

	for i := min; i <= max; i++ {
		if !g.contains(i) {
			g.resultStrIndxs = append(g.resultStrIndxs, i)
		}
	}
}

func (g *Grep) prepare() {
	if raw, ok := ioutil.ReadAll(g.flags.Input); ok != nil { //data collection
		log.Fatal(ok)
	} else {
		g.data = strings.Split(string(raw), "\n")
	}
	g.searchStr = g.flags.SearchStr

	if g.flags.Ignore {
		g.searchStr = strings.ToLower(g.searchStr) // Ignore flag preprocessing
		for i, str := range g.data {
			g.data[i] = strings.ToLower(str)
		}
	}
}

func (g *Grep) Run() []string {
	g.prepare()
	for i, str := range g.data {
		if g.compareString(str) {
			g.addStrings(i)
		}
	}
	return g.postPrecessing()
}

func (g *Grep) postPrecessing() []string {
	var resStrings []string

	if g.flags.Invert {
		if g.flags.OnlyCount {
			resStrings = append(resStrings, strconv.Itoa(len(g.data)-len(g.resultStrIndxs))) //If only strs count needed
		} else {
			for i, str := range g.data {
				if !g.contains(i) {
					if g.flags.StrNum {
						resStrings = append(resStrings, strconv.Itoa(i)+": "+str) //Writes string number in the beginning
					} else {
						resStrings = append(resStrings, str)
					}
				}
			}
		}
	} else {
		if g.flags.OnlyCount {
			resStrings = append(resStrings, strconv.Itoa(len(g.resultStrIndxs))) //If only strs count needed
		} else {
			for i, str := range g.data {
				if g.contains(i) {
					if g.flags.StrNum {
						resStrings = append(resStrings, strconv.Itoa(i)+": "+str) //Writes string number in the beginning
					} else {
						resStrings = append(resStrings, str)
					}
				}
			}
		}
	}

	return resStrings
}

func NewGrep(flags Parser.Flags) *Grep {
	return &Grep{nil, "", nil, flags}
}
