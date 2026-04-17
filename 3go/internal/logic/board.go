// Mantem e gerencia o estado do tabuleiro
package logic

type CellInfo struct {
	Liberties int

	// 0 empty
	// 1 black
	// 2 white
	State int
}

type Board struct {
	Board [19][19]CellInfo
}

func NewBoard() *Board {
	board := &Board{}
	for i := range 19 {
		for j := range 19 {
			liberties := 8
			if i == 0 || i == 18 {
				liberties -= 1
			}
			if j == 0 || j == 18 {
				liberties -= 1
			}
			board.Board[i][j] = CellInfo{liberties, 0}
		}
	}

	return board
}

func (b *Board) SetCell(x, y, player int) {
	b.Board[y][x] = CellInfo{8, player}
}
