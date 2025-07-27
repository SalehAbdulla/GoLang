package main

import "fmt"

func main() {
	// Step 1: How the board actualy function
	// var board [3][3]int
	// // Assign some values
	// board[0][0] = 1
	// board[1][1] = 1
	// board[2][2] = 1

	// // Extraordinary Print
	// for i := 0; i < 3; i++ {
	// 	for j := 0; j < 3; j++ {
	// 		fmt.Print(board[i][j], " ")
	// 	}
	// 	fmt.Println()
	// }

	// fmt.Println(board) // Ordinary Print

	// -----

	// Step 2: How to validate a row, column, and 3x3 box

	var board [9][9]int

	board[1][0] = 9
	// fmt.Println(board)

	// in a row Validation
	isNumInRow := func(board [9][9]int, row int, num int) bool {
		for col := 0; col < 9; col++ {
			if board[row][col] == num {
				return true
			}
		}
		return false
	}

	if isNumInRow(board, 1, 9) {
		fmt.Println("This number is exist in a row")
	}

	// in a column validation
	isNumInColumn := func(board [9][9]int, column int, num int) bool {
		for row := 0; row < 9; row++ {
			if board[row][column] == num {
				return true
			}
		}
		return false
	}

	if isNumInColumn(board, 0, 9) {
		fmt.Println("This number is exist in a column")
	}
	// in 3x3 box validation
	isNumInBox := func(board [9][9]int, col int, row int, num int) bool {
		startRow := (row / 3) * 3
		startCol := (col / 3) * 3

		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if board[startRow+i][startCol+j] == num {
					return true
				}
			}
		}
		return false
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Print(board[i][j])
			fmt.Print("  ")
		}
		fmt.Println()
	}

	if isNumInBox(board, 1, 0, 9) {
		fmt.Println("9 is exists in the 3x3 box")
	}

	// Step 3) Compine all isNumFunction in one function for validation
	// isValidPlacement := func(board [9][9]int, row, col, num int) bool {
	// 	return !isNumInRow(board, row, num) &&
	// 		!isNumInColumn(board, col, num) &&
	// 		!isNumInBox(board, row, col, num)
	// }

}
