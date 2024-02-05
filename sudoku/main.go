package main

import (
	"errors"
	"fmt"
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

func (cell *cellFiledSudoku) getCell() (int8, bool, error) {
	return cell.data, cell.accessChange, nil
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

func NewSudoku(area [9][9]int8) areaSudoku {
	var newAreaSudoku [9][9]cellFiledSudoku

	for Y, yArea := range area {
		for X, xArea := range yArea {
			if xArea != 0 {
				newAreaSudoku[Y][X] = cellFiledSudoku{
					data:         xArea,
					accessChange: false,
				}
			} else {
				newAreaSudoku[Y][X] = cellFiledSudoku{
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
	fmt.Println(game1)
}
