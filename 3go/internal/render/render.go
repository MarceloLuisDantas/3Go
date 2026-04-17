// Realiza a renderização de todos os elementos do jogo
package render

import (
	"3go/3go/internal/logic"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func RenderBoard(game *logic.Game) {
	// canto superior fica faltando 1 pixel sem esta linha
	rl.DrawLine(29, 30, 30, 30, rl.Black)
	for i := range 19 {
		offset := int32(30 + (30 * i))
		// linhas horizontais
		rl.DrawLineEx(
			rl.Vector2{X: 30, Y: float32(offset)},
			rl.Vector2{X: 30 * 19, Y: float32(offset)},
			2, rl.Black,
		)

		// linhas verticais
		rl.DrawLineEx(
			rl.Vector2{X: float32(offset), Y: 30},
			rl.Vector2{X: float32(offset), Y: 30 * 19},
			2, rl.Black,
		)
	}

	for i := range 19 {
		for j := range 19 {
			l := game.Board.Board[i][j].Liberties
			rl.DrawText(fmt.Sprintf("%d", l), int32(i*30)+25, int32(j*30)+21, 20, rl.Red)
		}
	}

	for i := range 19 {
		for j := range 19 {
			switch game.Board.Board[i][j].State {
			case 1:
				rl.DrawTexture(game.Sprites.BlackStone, int32(i*30)+15, int32(j*30)+15, rl.Beige)
			case 2:
				rl.DrawTexture(game.Sprites.WhiteStone, int32(i*30)+15, int32(j*30)+15, rl.Beige)
			}
		}
	}
}

func RenderStoneToBePlaced(game *logic.Game) {
	mouse_x := rl.GetMouseX()
	mouse_y := rl.GetMouseY()

	point_x := ((mouse_x + 15) / 30) - 1
	point_y := ((mouse_y + 15) / 30) - 1
	if game.Board.Board[point_x][point_y].State == 0 {
		var stone rl.Texture2D
		if game.CurrentPlayer == 1 {
			stone = game.Sprites.BlackStone
		} else {
			stone = game.Sprites.WhiteStone
		}

		point_x = (((mouse_x + 15) / 30) * 30)
		point_y = (((mouse_y + 15) / 30) * 30)
		rl.DrawTexture(stone, point_x-15, point_y-15, rl.Beige)
	}
}

func RenderGame(game *logic.Game) {
	mouse_x := rl.GetMouseX()
	mouse_y := rl.GetMouseY()

	RenderBoard(game)

	// Mouse is over the board
	if (mouse_y >= 29 && mouse_y <= (30*19)) &&
		(mouse_x >= 29 && mouse_x <= (30*19)) {
		RenderStoneToBePlaced(game)
	}
}
