package main

import (
	"fmt"
)

type debris struct {
	gActor  gameActor
	spawner *spawner
	health  uint32
}

func newDebris(s *spawner, v vector) *debris {
	d := &debris{spawner: s}
	d.Init()

	//d.addComponent(newEntityRendererComp(player, renderer, "assets/playership.bmp"))
	//d.addComponent(newKeyboardInput(player))
	return d
}

func (actor *debris) Mailbox() chan payload {
	return actor.gActor.base.mailbox
}

func (actor *debris) Init() {
	actor.gActor.base.mailbox = make(chan payload, 1)
	go actor.receive(actor.gActor.base.mailbox)
}

func (actor *debris) attacked(damage uint32) {
	actor.health -= damage
	if actor.health <= 0 {
		actor.health = 0
		actor.spawner.Mailbox() <- payload{msg: "dead"}
	}
}

func (actor *debris) receive(ch <-chan payload) {
	for {
		payload := <-ch
		switch payload.msg {
		case "startup":
			// todo: get this data from some external source
			fmt.Println("sector startup message")
		case "update":
			fmt.Println("update message")
		case "attack":
			fmt.Println("attack message")
		default:
			fmt.Println("unhandled message")
		}
	}
}
