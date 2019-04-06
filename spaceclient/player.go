package main

import "github.com/veandco/go-sdl2/sdl"

func newPlayer(renderer *sdl.Renderer) *entity {
	player := &entity{}

	player.addComponent(newEntityRendererComp(player, renderer, "assets/playership.bmp"))
	player.addComponent(newKeyboardInput(player))
	return player
}
