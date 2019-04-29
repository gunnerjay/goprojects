package main

import "fmt"

type spawner struct {
	base          baseActor
	currentActors []iActor
}

func (actor *spawner) Mailbox() chan payload {
	return actor.base.mailbox
}

func (actor *spawner) Init() {
	actor.base.mailbox = make(chan payload, 1)
	go actor.receive(actor.base.mailbox)
}

func (actor *spawner) startup() {
	// TODO: get this data from some external source
	numActors := 1
	for i := 0; i < numActors; i++ {
		d := newDebris(actor, vector{0.0, 0.0, 0.0})
		actor.currentActors = append(actor.currentActors, d)
	}
}

func (actor *spawner) receive(ch <-chan payload) {
	for {
		payload := <-ch
		switch payload.msg {
		case "spawn":
			fmt.Println("spawner startup message")
		case "update":
			fmt.Println("update message")
		default:
			fmt.Println("unhandled message")
		}
	}
}
