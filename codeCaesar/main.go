package main

import (
	"fmt"
	"unicode/utf8"
)

func decodeingLineInCodeCeasar(line string) string{
	var new_line = make([]rune, 0, utf8.RuneCountInString(line))
	for _, word := range line {
		if (word >= 'a' && word <= 'z'){
			word -= 3 
			if word < 'a'{
				word = word + 26
			}
		}
		if (word >= 'A' && word <= 'Z'){
			word -= 3 
			if word < 'A'{
				word = word + 26
			}
		}
		new_line = append(new_line, word)
	}
	return string(new_line)
}

func codeingLineInCodeCeasar(line string) string{
	var new_line = make([]rune, 0, utf8.RuneCountInString(line))
	for _, word := range line {
		if (word >= 'a' && word <= 'z'){
			word += 3 
			if word > 'z'{
				word = word + 26
			}
		}
		if (word >= 'A' && word <= 'Z'){
			word += 3 
			if word > 'Z'{
				word = word + 26
			}
		}
		new_line = append(new_line, word)
	}
	return string(new_line)
}

func main(){
	var line string = "L fdph, L vdz, L frqtxhuhg."
	fmt.Print(decodeingLineInCodeCeasar(line))
}