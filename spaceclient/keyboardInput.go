package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type keyboardInput struct {
	owner *entity
	speed float64
}

func newKeyboardInput(owner *entity) *keyboardInput {
	return &keyboardInput{
		owner: owner,
		speed: 4,
	}
}

func (kInput *keyboardInput) update(delta float64) error {
	keys := sdl.GetKeyboardState()

	owner := kInput.owner

	if keys[sdl.SCANCODE_A] == 1 {
		owner.position.x = owner.position.x - (kInput.speed * delta)
		owner.position.x = math.Max(0, owner.position.x)
	} else if keys[sdl.SCANCODE_D] == 1 {
		owner.position.x = owner.position.x + (kInput.speed * delta)
		owner.position.x = math.Min(1024-53, owner.position.x)
	}

	return nil
}

func (kInput *keyboardInput) draw(renderer *sdl.Renderer, delta float64) error {
	return nil
}
