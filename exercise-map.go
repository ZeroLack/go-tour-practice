package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	ret := make(map[string]int)
	arr := strings.Fields(s)
	for _, word := range arr {
		ret[word]++
	}
	return ret
}

func main() {
	wc.Test(WordCount)
}
