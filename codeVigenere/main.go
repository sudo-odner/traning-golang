package main

import (
	"fmt"
	"unicode"
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
	if 'A' <= symbolFirst && symbolFirst <= 'Z' {
		var differenceSymbol rune = (symbolFirst - unicode.ToUpper(symboleSecond))
		if differenceSymbol < 0 {
			differenceSymbol *= -1
		}
		var newSymbol = symbolFirst + differenceSymbol
		if newSymbol > 'Z' {
			newSymbol -= 26
		}
		return newSymbol
	}

	var differenceSymbol rune = symbolFirst - symboleSecond
	if differenceSymbol < 0 {
		differenceSymbol *= -1
	}
	var newSymbol = symbolFirst + differenceSymbol
	if newSymbol > 'z' {
		newSymbol -= 26
	}
	return newSymbol
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

func codeInCodeVigenere(message string, key string) string{
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
	return ""
}

func main() {
	var message string = "CSOITEUIWUIZNSROCNKFD"
	var key string = "GOLANG"

	codeInCodeVigenere(message, key)
}