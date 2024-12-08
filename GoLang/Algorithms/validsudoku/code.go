package validsudoku

import "fmt"

func Solve(board [][]byte) {
	fmt.Println(isValidSudoku(board))
}

func isValidSudoku(board [][]byte) bool {
	rows := make([]int, 9)
	cols := make([]int, 9)
	squares := make([]int, 9)

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == '.' {
				continue
			}

			val := board[r][c] - '1'
			bit := 1 << val
			squareIdx := (r/3)*3 + c/3

			if rows[r]&bit != 0 || cols[c]&bit != 0 ||
				squares[squareIdx]&bit != 0 {
				return false
			}

			rows[r] |= bit
			cols[c] |= bit
			squares[squareIdx] |= bit
		}
	}

	return true
}

// func isValidSudoku(board [][]byte) bool {
// 	r := 0

// 	for i, row := range board {
// 		c := 0
// 		if c != len(row) {
// 			fmt.Println(row, board[i])
// 			fmt.Println(row[r], board[i][c])
// 		}
// 		c++
// 	}

// 	return false
// }
