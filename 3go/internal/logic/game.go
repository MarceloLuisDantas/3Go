package logic

import "3go/3go/internal/textures"

type Game struct {
	Board         *Board
	Groups        *Groups
	CurrentPlayer int
	BlackScore    int
	WhiteScore    int
	Sprites       *textures.Textures
}

func NewGame() *Game {
	game := &Game{
		Board:         NewBoard(),
		Groups:        NewGroups(),
		CurrentPlayer: 1,
		BlackScore:    0,
		WhiteScore:    0,
		Sprites:       textures.LoadTextures(),
	}

	return game
}
