package main

import (
	"errors"
	"fmt"
	"strings"
)

type cellFiledSudoku struct {
	data         int8
	accessChange bool
}

func (cell *cellFiledSudoku) setCell(value int8) error {
	if value > 9 && value < 0 {
		return errors.New("Вышел за диапазон чисел")
	}
	if cell.accessChange != true {
		return errors.New("Данное поле не изменяется")
	}
	cell.data = value
	return nil
}

func (cell *cellFiledSudoku) getCell() int8 {
	return cell.data
}

type areaSudoku struct {
	area [9][9]cellFiledSudoku
}

func (area *areaSudoku) setCell(X, Y, value int8) {
	if X < 0 && X > 9 || Y < 0 && Y > 9 {
		panic("Координаты X или Y выходят за диапазон")
	}
	err := area.area[X][Y].setCell(value)

	if err != nil {
		panic(err)
	}
}

func (area *areaSudoku) setCellZero(X, Y int8) {
	if X < 0 && X > 9 || Y < 0 && Y > 9 {
		panic("Координаты X или Y выходят за диапазон")
	}
	err := area.area[X][Y].setCell(0)

	if err != nil {
		panic(err)
	}
}

func (area *areaSudoku) returnArea() {
	var mainLine = strings.Repeat("-", 41)
	var otherLine = "[]---|---|---[]---|---|---[]---|---|---[]"

	for y, yData := range area.area {
		if y%3 == 0 {
			fmt.Println(mainLine)
		} else {
			fmt.Println(otherLine)
		}

		for x, xData := range yData {
			value := xData.getCell()

			if x%3 == 0 {
				fmt.Printf("[] %v ", value)
			} else {
				fmt.Printf("| %v ", value)
			}

			/*	Without zero
					if x%3 == 0 {
						if value == 0 {
							fmt.Print("[]   ")
						} else {
							fmt.Printf("[] %v ", value)
						}
					} else {
						if value == 0 {
							fmt.Print("|   ")
						} else {
							fmt.Printf("| %v ", value)
						}
					}
				}
			*/
		}
		fmt.Print("[]\n")
	}
	fmt.Println(mainLine)
}

func NewSudoku(area [9][9]int8) areaSudoku {
	var newAreaSudoku [9][9]cellFiledSudoku

	for y, yArea := range area {
		for x, xArea := range yArea {
			if xArea != 0 {
				newAreaSudoku[y][x] = cellFiledSudoku{
					data:         xArea,
					accessChange: false,
				}
			} else {
				newAreaSudoku[y][x] = cellFiledSudoku{
					data:         xArea,
					accessChange: true,
				}
			}
		}
	}
	return areaSudoku{area: newAreaSudoku}
}

func main() {
	game1 := NewSudoku([9][9]int8{
		{5, 2, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	})
	game1.setCell(0, 2, 5)
	game1.returnArea()
	game1.setCell(0, 2, 2)
	game1.setCell(0, 1, 2)
	game1.returnArea()
}
