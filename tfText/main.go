package main

import (
	"strings"
	"fmt"

)
func cleanUnaccessSymbol(text string, dataUnaccessSymbol []string) string {
	var claenText string = text
	for _, symbol := range dataUnaccessSymbol {
		claenText = strings.Replace(claenText, symbol, "", -1)
	}

	return claenText
}
func tf(text string) string { // map[string]int
	var lowerText = strings.ToLower(text)
	var cleanText string = cleanUnaccessSymbol(lowerText, strings.Split("!@#$%^&*()))))_+-=±§><|[]?/,.;:~`", ""))
	var dataWordText = strings.Split(cleanText, " ")

	fmt.Println(dataWordText)
	return ""
}

func main() {
	var inputText string = "As far as eye could reach he saw nothing but the stems of the great plants about him receding in the violet shade, and far overhead the multiple transparency of huge leaves filtering the sunshine to the solemn splendour of twilight in which he walked. Whenever he felt able he ran again; the ground continued soft and springy, covered with the same resilient weed which was the first thing his hands had touched in Malacandra. Once or twice a small red creature scuttled across his path, but otherwise there seemed to be no life stirring in the wood; nothing to fear—except the fact of wandering unprovisioned and alone in a forest of unknown vegetation thousands or millions of miles beyond the reach or knowledge of man."
	inputText = "Привет мир, я надеюсь не умереть"
	tf(inputText)
}