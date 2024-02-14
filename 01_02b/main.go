package main

import (
	"log"
	"strings"
	"time"
)

const delay = 700 * time.Millisecond

// print outputs a message and then sleeps for a pre-determined amount
func print(msg string) {
	log.Println(msg)
	time.Sleep(delay)
}

// slowDown takes the given string and repeats its characters
// according to their index in the string.
func slowDown(msg string) {
	for _, word := range strings.Split(msg, " ") {
		print(expand(word))
	}
}

func expand(word string) string {
	var result string
	for i, c := range word {
		result += strings.Repeat(string(c),i+1)
	}
	return result
}

func main() {
	msg := "Time to learn about Go strings!"
	slowDown(msg)
}
