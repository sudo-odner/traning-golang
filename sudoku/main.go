package main

import (
	"fmt"
)

type cellFiledSudoku struct {
	data         int8
	accessChange bool
}

type areaSudoku struct {
	area [9][9]cellFiledSudoku
}

func NewSudoku(area [9][9]int8) areaSudoku {
	var newAreaSudoku [9][9]cellFiledSudoku
	// if len(area) != 9 {
	// 	panic("Арена имеет не стандартый вид по y координатам")
	// }
	// for _, yArea := range area {
	// 	if len(yArea) != 9 {
	// 		panic("Арена имеет не стандартый вид по x координатам")
	// 	}
	// }

	for idY, yArea := range area {
		for idX, xArea := range yArea {
			if xArea != 0 {
				newAreaSudoku[idY][idX] = cellFiledSudoku{
					data:         xArea,
					accessChange: true,
				}
			} else {
				newAreaSudoku[idY][idX] = cellFiledSudoku{
					data:         xArea,
					accessChange: false,
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
	fmt.Println(game1)
}
