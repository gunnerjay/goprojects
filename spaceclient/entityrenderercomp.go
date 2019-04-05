package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type entityRendererComp struct {
	owner   *entity
	texture *sdl.Texture
}

func newEntityRendererComp(owner *entity, renderer *sdl.Renderer, filename string) *entityRendererComp {
	return &entityRendererComp{
		owner:   owner,
		texture: textureFromBMP(renderer, filename),
	}
}

func (erc *entityRendererComp) update(delta float64) error {
	return nil
}

func (erc *entityRendererComp) draw(renderer *sdl.Renderer, delta float64) error {
	_, _, width, height, err := erc.texture.Query()
	if err != nil {
		return fmt.Errorf("querying texture: %v", err)
	}

	renderer.CopyEx(
		erc.texture,
		&sdl.Rect{X: 0, Y: 0, W: width, H: height},
		&sdl.Rect{X: int32(erc.owner.position.x), Y: int32(erc.owner.position.y), W: width, H: height},
		erc.owner.rotation,
		&sdl.Point{X: width / 2, Y: height / 2},
		sdl.FLIP_NONE)

	return nil
}

func textureFromBMP(renderer *sdl.Renderer, filename string) *sdl.Texture {
	img, err := sdl.LoadBMP(filename)
	if err != nil {
		panic(fmt.Errorf("loading %v: %v", filename, err))
	}
	defer img.Free()
	pf, _ := sdl.AllocFormat(sdl.PIXELFORMAT_RGB24)
	img.SetColorKey(true, sdl.MapRGB(pf, 255, 255, 255))

	tex, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		panic(fmt.Errorf("creating texture from %v: %v", filename, err))
	}

	return tex
}
