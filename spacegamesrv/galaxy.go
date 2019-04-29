package main

import "fmt"

type galaxy struct {
	base    baseActor
	sectors []iActor
}

func (actor *galaxy) Mailbox() chan payload {
	return actor.base.mailbox
}

func (actor *galaxy) Init() {
	actor.base.mailbox = make(chan payload, 1)
	go actor.receive(actor.base.mailbox)
}

func (actor *galaxy) startup() {
	// TODO: get this data from some external source
	var a = sector{name: "Sector1"}
	a.Init()
	a.Mailbox() <- payload{msg: "startup", srcActor: int64(actor)}
	actor.sectors = append(actor.sectors, &a)
	var b = sector{name: "Sector2"}
	b.Init()
	b.Mailbox() <- payload{msg: "startup"}
	actor.sectors = append(actor.sectors, &b)
}

func (actor *galaxy) receive(ch <-chan payload) {
	for {
		payload := <-ch
		switch payload.msg {
		case "startup":
			fmt.Println("galaxy startup")
			actor.startup()
		case "update":
			fmt.Println("update message")
		default:
			fmt.Println("unhandled message")
		}
	}
}
