package piscine

import (
	"ft"
)

func PrintStr(str string) {
	for _, c := range str {
		ft.PrintRune(c)
	}
}

func PrintCheck(isSuccess bool) {
	if isSuccess {
		PrintStr("Success\n")
	} else {
		PrintStr("Fail\n")
	}
}

// King K
// Pawn P
// Bishop B
// Rook R
// Queen Q
func Checkmate(board []string) {
	isSuccess := false
	for y, line := range board {
		for x, c := range line {
			switch c {
			case 'P':
				isSuccess = CheckPawn(x, y, board)
			case 'B':
				isSuccess = CheckBishop(x, y, board)
			case 'R':
				isSuccess = CheckRook(x, y, board)
			case 'Q':
				isSuccess = CheckQueen(x, y, board)
			}
			if isSuccess {
				PrintCheck(isSuccess)
				return
			}
		}
	}
	PrintCheck(isSuccess)
}

/*
Pawn (P):
. . . . . . .
. . . . . . .
. . X . X . .
. . . P . . .
. . . . . . .
. . . . . . .
. . . . . . .
*/

func CheckPawn(x, y int, board []string) bool {
	p1 := [2]int{x + 1, y - 1}
	p2 := [2]int{x - 1, y - 1}
	if IsInBoarder(board, p1[0], p1[1]) && board[p1[0]][p1[1]] == 'K' {
		return true
	}
	if IsInBoarder(board, p2[0], p2[1]) && board[p2[0]][p2[1]] == 'K' {
		return true
	}
	return false
}

/*
Bishop (B):
X . . . . . X
. X . . . X .
. . X . X . .
. . . B . . .
. . X . X . .
. X . . . X .
X . . . . . X
*/

func CheckBishop(x, y int, board []string) bool {
	// check left up
checkLeftUp:
	for i := x - 1; i >= 0; i-- {
		for j := y - 1; j >= 0; j-- {
			if board[j][i] == 'K' {
				return true
			}
			if board[j][i] != '.' {
				break checkLeftUp
			}
		}
	}
	// check left down
checkLeftDown:
	for i := x - 1; i >= 0; i-- {
		for j := y + 1; j < len(board); j++ {
			if board[j][i] == 'K' {
				return true
			}
			if board[j][i] != '.' {
				break checkLeftDown
			}
		}
	}
	// check right up
checkRightUp:
	for i := x + 1; i < len(board[0]); i++ {
		for j := y - 1; j >= 0; j-- {
			if board[j][i] == 'K' {
				return true
			}
			if board[j][i] != '.' {
				break checkRightUp
			}
		}
	}
	// check right down
checkRightDown:
	for i := x + 1; i < len(board[0]); i++ {
		for j := y + 1; j < len(board); j++ {
			if board[j][i] == 'K' {
				return true
			}
			if board[j][i] != '.' {
				break checkRightDown
			}
		}
	}
	return false
}

/*
Rook (R):
. . . X . . .
. . . X . . .
. . . X . . .
X X X R X X X
. . . X . . .
. . . X . . .
. . . X . . .
*/

func CheckRook(x, y int, board []string) bool {
	// check left
	for i := x + 1; i < len(board[0]); i++ {
		if board[y][i] == 'K' {
			return true
		}
		if board[y][i] != '.' {
			break
		}
	}
	// check right
	for i := x - 1; i >= 0; i-- {
		if board[y][i] == 'K' {
			return true
		}
		if board[y][i] != '.' {
			break
		}
	}
	// check up
	for j := y + 1; j < len(board); j++ {
		if board[j][x] == 'K' {
			return true
		}
		if board[j][x] != '.' {
			break
		}
	}
	// check down
	for j := y - 1; j >= 0; j-- {
		if board[j][x] == 'K' {
			return true
		}
		if board[j][x] != '.' {
			break
		}
	}
	return false
}

func CheckQueen(x, y int, board []string) bool {
	if CheckBishop(x, y, board) || CheckRook(x, y, board) {
		return true
	}
	return false
}

func IsInBoarder(board []string, width, height int) bool {
	return (width >= 0 && width < len(board[0])) && (height >= 0 && height < len(board))
}
