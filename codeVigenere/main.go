package main

import (
	"fmt"
	// "unicode"
	"strings"
	"unicode/utf8"
)

func genereteKeyByMessage(message string, key string) []rune {
	var sliceMassage = []rune(message)
	var sliceKey = []rune(key)

	var newKey = make([]rune, 0, len(sliceMassage))

	for i := 0; i < len(sliceMassage); i++ {
		newKey = append(newKey, sliceKey[(i % len(sliceKey))])
	}
	return newKey
}

func codingSymbolByCodeVigenere(symbolFirst rune, symboleSecond rune) rune{
	var differenceSymbol rune = 'a' - symboleSecond
	if differenceSymbol < 0 {
		differenceSymbol *= -1
	}
	var newSymbol rune = symbolFirst + differenceSymbol

	if 'A' <= symbolFirst && symbolFirst <= 'Z'{
		return 'A' + ((newSymbol - 'A') % 26)
	}
	return 'a' + ((newSymbol - 'a') % 26)
}

func runeInSlice(desired rune, slice []rune) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i] == desired {
			return false
		}
	}
	return true
}

func checkAnotherSymol(message string) bool {
	var flag int = utf8.RuneCountInString(message)
	for _, symbol := range message {
		var unaccessSymbol = []rune("!@#$%^&*()[]{}|<>.,/±§œ∑´®†¥¨ˆøπ“‘åß∂ƒ©˙∆˚¬…æ«≈ç√∫˜µ≤≥÷")
		if runeInSlice(symbol, unaccessSymbol){
			flag -= 1
		}
	}
	if flag == 0{
		return true
	}
	return false
}

func codeByCodeVigenere(message string, key string) string{
	if !checkAnotherSymol(message) {
		return "Err in message, has unaccess symbol"
	}
	var sliceMassage = []rune(message)
	var sliceKey = genereteKeyByMessage(message, strings.ToLower(key))

	var newMessage = make([]rune, 0, len(sliceMassage))

	for i := 0; i < len(sliceMassage); i++ {
		newMessage = append(newMessage, codingSymbolByCodeVigenere(sliceMassage[i], sliceKey[i]))
	}
	fmt.Println(string(newMessage))
	return string(newMessage)
}

func uncodingSymbolByCodeVigenere(symbolFirst rune, symboleSecond rune) rune{
	var differenceSymbol rune = 'a' - symboleSecond
	if differenceSymbol < 0 {
		differenceSymbol *= -1
	}
	var newSymbol rune = symbolFirst - differenceSymbol

	if 'A' <= symbolFirst && symbolFirst <= 'Z' && newSymbol <= 'A' {
		return newSymbol + 26
	}
	if 'a' <= symbolFirst && symbolFirst <= 'z' && newSymbol <= 'a' {
		return newSymbol + 26
	}
	return newSymbol
}

func uncodeByCodeVigenere(message string, key string) string{
	if !checkAnotherSymol(message) {
		return "Err in message, has unaccess symbol"
	}
	var sliceMassage = []rune(message)
	var sliceKey = genereteKeyByMessage(message, strings.ToLower(key))

	var newMessage = make([]rune, 0, len(sliceMassage))

	for i := 0; i < len(sliceMassage); i++ {
		newMessage = append(newMessage, uncodingSymbolByCodeVigenere(sliceMassage[i], sliceKey[i]))
	}
	fmt.Println(string(newMessage))
	return string(newMessage)
}

func main() {
	var message string = "zsdt"
	var key string = "GOLANG"

	uncodeByCodeVigenere(message, key)
}