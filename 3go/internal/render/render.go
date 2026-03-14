// Realiza a renderização de todos os elementos do jogo
package render

import (
	"3go/3go/internal/logic"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func RenderBoard(board *logic.Board) {
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
}

func RenderStoneToBePlaced(game *logic.Game) {
	mouse_x := rl.GetMouseX()
	mouse_y := rl.GetMouseY()

	var stone rl.Texture2D
	if game.CurrentPlayer == 1 {
		stone = game.Sprites.BlackStone
	} else {
		stone = game.Sprites.WhiteStone
	}

	rl.DrawTexture(stone, mouse_x, mouse_y, rl.Beige)
}

func RenderGame(game *logic.Game) {
	RenderBoard(game.Board)
	RenderStoneToBePlaced(game)
}
