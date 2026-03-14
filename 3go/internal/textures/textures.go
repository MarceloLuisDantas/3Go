package textures

import rl "github.com/gen2brain/raylib-go/raylib"

type Textures struct {
	BlackStone rl.Texture2D
	WhiteStone rl.Texture2D
}

func LoadTextures() *Textures {
	t := &Textures{}

	t.BlackStone = rl.LoadTexture("3go/assets/black_stone.png")
	t.WhiteStone = rl.LoadTexture("3go/assets/white_stone.png")

	return t
}
