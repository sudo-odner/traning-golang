package main

import (
	"fmt"
	"strings"
)

type Universe [][]bool

const (
	width = 80
	height = 15
	live = "#"
	dead = " "
)

func (univers Universe) createSize(width int, height int) Universe{
	new_univers := make([][]bool, height)
	for i := 0; i < height; i++ {
		new_univers[i] = make([]bool, width)
	}
	return Universe(new_univers)
}

func (univers Universe) printUniverse() {
	fmt.Printf("%v\n", strings.Repeat("=", len(univers[0]) + 2))
	for _, line := range univers {
		for index, info := range line {
			if index == 0 {
				fmt.Printf("|")
			}
			if info{
				fmt.Printf("%v", live)
			} else {
				fmt.Printf("%v", dead)
			}
			if index == len(univers[0]) - 1 {
				fmt.Printf("|")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("%v\n", strings.Repeat("=", len(univers[0]) + 2))
}

func main() {
	var univers Universe
	univers = univers.createSize(width, height)
	univers.printUniverse()

	// fmt.Println(univers)
}