package main

import (
	"fmt"
	"strings"
	"math/rand"
)

type Universe [][]bool

const (
	width = 80
	height = 15
	live = "#"
	dead = " "
	percentAlive = 0.25
)

func (univers Universe) createSize(width int, height int) Universe{
	// Я еще не работал с ошибками в go поэтому простй отладчик
	if width <= 0 || height <= 0 {
		return univers
	}

	newUnivers := make(Universe, height)
	for i := 0; i < height; i++ {
		newUnivers[i] = make([]bool, width)
	}
	return Universe(newUnivers)
}

func (univers Universe) randSeed(percentAlive float64) Universe {
	// Я еще не работал с ошибками в go поэтому простй отладчик
	if percentAlive > 1 || percentAlive < 0 {
		return univers
	}

	width, height := len(univers[0]), len(univers)
	var liveCells int = int(float64(width) * float64(height) * percentAlive)// По условию задачи нужно заполнить 25%
	for i := 0; i < liveCells; i++ {
		randHeihgt, randWidth := rand.Intn(height), rand.Intn(width)
		for univers[randHeihgt][randWidth]{
			randHeihgt, randWidth = rand.Intn(height), rand.Intn(width)
		}
		univers[randHeihgt][randWidth] = true
	}
	return univers
}

func (univers Universe) Show() {
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
	univers.randSeed(0.25)
	univers.Show()

	// fmt.Println(univers)
}