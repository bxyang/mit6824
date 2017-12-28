


package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {

	var a []string = strings.Fields(s)
	var m map[string]int = make(map[string]int)
	for _, v := range a {
		r, ok := m[v]
		if ! ok {
			r = 0
		}
		m[v] = r + 1
	}
	
	return m
}

func main() {
	wc.Test(WordCount)
}

