package main

import (
	"fmt"
	"strings"
	"math/rand"
    "os"
    "os/exec"
	"time"
)

type Universe [][]bool

const (
	width = 80
	height = 15
	live = "#"
	dead = "-"
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
func (univers Universe) Alive(x, y int) bool {
	height, width := len(univers), len(univers[0])

	if x < 0{
		x = height + (x % height)
	}
	if x > 0{
		x = (x % height)
	}
	if y < 0{
		y = width + (y % width)
	}
	if y > 0{
		y = (y % width)
	}
	if univers[x][y] {
		return true
	}
	return false
}

func (univers Universe) Neighbors(x, y int) int {
	height, width := len(univers), len(univers[0])
	// Я еще не работал с ошибками в go поэтому простй отладчик
	if height < x || width < y {
		fmt.Println("cord x more size or cord y more size", x, y, height, width)
		return -1
	}
	var count int = 0
	for h := -1; h < 2; h++{
		for w := -1; w < 2; w++{
			if univers.Alive(x+h, y+w){
				count++
			}
		}
	}
	if univers.Alive(x, y) {
		count--
	}
	return count
}

func (univers Universe) next(x, y int) bool {
	if univers.Alive(x, y) && univers.Neighbors(x, y) < 2 {
		return false
	}
	if univers.Alive(x, y) && (univers.Neighbors(x, y) == 2 || univers.Neighbors(x, y) == 3){
		return true
	}
	if univers.Alive(x, y) && (univers.Neighbors(x, y) > 3) {
		return false
	}
	if !univers.Alive(x, y) && (univers.Neighbors(x, y) == 3) {
		return true
	}

	// Тут желательно внести проверку
	return univers.Alive(x, y)
}

func (univers Universe) NextStep() Universe {
	height, width := len(univers), len(univers[0])
	var newUnivers Universe
	newUnivers = newUnivers.createSize(width, height)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			newUnivers[i][j] = univers.next(i, j)
		}
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			univers[i][j] = newUnivers[i][j]
		}
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
	fmt.Printf("%v\n\n", strings.Repeat("=", len(univers[0]) + 2))
}

func main() {
	var univers Universe
	univers = univers.createSize(width, height)
	univers.randSeed(0.25)

	var 
	for i != 0{
		univers.Show()
	}
	univers.Show()
	univers.NextStep()
	univers.Show()

    cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Run()

	// fmt.Println(univers)
}