package main

import (
	"fmt"
	"strings"
)

func alphaberCount(s string) map[string]int {
	res := make(map[string]int)

	for _, charachter := range s {
		ch := string(charachter)
		if res[ch] == 0 {
			res[ch] = 0
		}
		res[ch] += 1
	}
	return res
}

func wordCount(s string) map[string]int {
	words := make(map[string]int)
	splited := strings.Split(s, " ")

	for _, word := range splited {
		if words[word] == 0 {
			words[word] = 0
		}
		words[word] += 1

	}
	return words
}
func main() {
	tes := "Hi World Hi"
	fmt.Println(alphaberCount("Hi Hi World"))
	fmt.Println(wordCount(tes))

}
