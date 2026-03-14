// Mantem e gerencia o estado do tabuleiro
package logic

type Board struct {
	board [19][19]int
}

func NewBoard() *Board {
	return &Board{}
}

func (b *Board) SetCell(x, y, player int) {
	b.board[y][x] = player
}

// [x1, y1][x2, y2][x3, y3]...[xn, yn]
func (b *Board) KillGroup(cells [][]int) {
	for _, cell := range cells {
		x, y := cell[0], cell[1]
		b.board[y][x] = 0
	}
}
