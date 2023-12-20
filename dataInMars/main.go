package main

import (
	"fmt"
	"math/rand"
)

var era = "AD"
 
func main() {
	for i := 0; i < 10; i++ {
		year := rand.Intn(10000)
		month := rand.Intn(12) + 1
		daysInMonth := 31
	
		switch month {
		case 2:
			switch year % 400 == 0 {
				case true:
					daysInMonth = 29
				case false:
					daysInMonth = 28
			}
		case 4, 6, 9, 11:
			daysInMonth = 30
		}
	
		day := rand.Intn(daysInMonth) + 1
		fmt.Println(era, year, month, day)
	}
}