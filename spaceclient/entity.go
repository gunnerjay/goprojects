package main

import (
	"fmt"
	"reflect"

	"github.com/veandco/go-sdl2/sdl"
)

type vector struct {
	x, y, z float64
}

type component interface {
	draw(renderer *sdl.Renderer, delta float64) error
	update(delta float64) error
}

type entity struct {
	position   vector
	rotation   float64
	active     bool
	components []component
}

func (ent *entity) addComponent(new component) {
	// TODO: check for duplicates
	ent.components = append(ent.components, new)
}

func (ent *entity) getComponent(ofType component) component {
	findType := reflect.TypeOf(ofType)
	for _, comp := range ent.components {
		if findType == reflect.TypeOf(comp) {
			return comp
		}
	}

	panic(fmt.Sprintf("no component of type %s", reflect.TypeOf(ofType)))
}

func (ent *entity) draw(renderer *sdl.Renderer, delta float64) error {
	for _, comp := range ent.components {
		err := comp.draw(renderer, delta)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ent *entity) update(delta float64) error {
	for _, comp := range ent.components {
		err := comp.update(delta)
		if err != nil {
			return err
		}
	}

	return nil
}
