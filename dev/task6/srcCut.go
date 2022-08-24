package main

import "strings"

type Cutter struct {
	flags Flags
}

func (c *Cutter) Process(str string) []string {
	var cuttedStr = strings.Split(str, c.flags.d)
	var res []string
	if c.flags.s && len(cuttedStr) == 1 {
		return nil
	}

	if c.flags.f != nil {
		for _, ind := range c.flags.f {
			if tmp := c.getFromColumn(ind, cuttedStr); tmp != "" {
				res = append(res, tmp)
			}
		}
	} else {
		return cuttedStr
	}

	return res
}

func (c *Cutter) getFromColumn(ind int, strs []string) string {
	if ind > len(strs)-1 {
		return ""
	}

	return strs[ind]
}
