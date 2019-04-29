package main

import "fmt"

type sector struct {
	base          baseActor
	name          string
	sectorSpawner spawner
}

func (actor *sector) Mailbox() chan payload {
	return actor.base.mailbox
}

func (actor *sector) Init() {
	actor.base.mailbox = make(chan payload, 1)
	go actor.receive(actor.base.mailbox)
}

func (actor *sector) receive(ch <-chan payload) {
	for {
		payload := <-ch
		switch payload.msg {
		case "startup":
			// todo: get this data from some external source
			fmt.Println("sector startup message")
		case "update":
			fmt.Println("update message")
		default:
			fmt.Println("unhandled message")
		}
	}
}
