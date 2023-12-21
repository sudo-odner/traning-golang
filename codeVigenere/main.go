package main

import (
	"fmt"
	"unicode/utf8"
)

func codeInCodeVigenere(message string, key string) string{
	var sizeMessage = utf8.RuneCountInString(message)
	var newMessage = make([]rune, 0, sizeMessage)
	var sliceKey = []rune(key)

	for i := 0; i < sizeMessage; i++ {
		newMessage = append(newMessage, sliceKey[(i % len(sliceKey))])
	}
	fmt.Println(string(newMessage))
	return ""
}

func main() {
	var message string = "CSOITEUIWUIZNSROCNKFD"
	var key string = "GOLANG"

	codeInCodeVigenere(message, key)
}