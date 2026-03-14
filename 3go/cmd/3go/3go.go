package main

import (
	logic "3go/3go/internal/logic"
	render "3go/3go/internal/render"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 600, "Hello, Raylib!")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	game := logic.NewGame()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.EndDrawing()
		render.RenderGame(game)
	}
}
