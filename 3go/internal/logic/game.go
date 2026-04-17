package logic

import (
	"3go/3go/internal/textures"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	Board         *Board
	Groups        *Groups
	CurrentPlayer int // 1 = black, 2 = white
	Sprites       *textures.Textures
}

func NewGame() *Game {
	game := &Game{
		Board:         NewBoard(),
		Groups:        NewGroups(),
		CurrentPlayer: 1,
		Sprites:       textures.LoadTextures(),
	}

	return game
}

func IsMouseOverBoard() bool {
	x := rl.GetMouseX()
	y := rl.GetMouseY()
	return (y >= 29 && y <= 570) && (x >= 29 && x <= 570)
}

func (game *Game) RunLogic() {
	mouse_x := rl.GetMouseX()
	mouse_y := rl.GetMouseY()

	if IsMouseOverBoard() {
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			point_x := ((mouse_x + 15) / 30) - 1
			point_y := ((mouse_y + 15) / 30) - 1

			cell := game.Board.Board[point_x][point_y]
			if cell.State == 0 {
				game.Board.Board[point_x][point_y].State = game.CurrentPlayer
				if game.CurrentPlayer == 1 {
					game.CurrentPlayer = 2
				} else {
					game.CurrentPlayer = 1
				}
			}
		}
	}
}
