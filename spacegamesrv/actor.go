package main

import "os"

type payload struct {
	msg      string
	srcNode  string
	srcActor uint64
	//srcActor  *baseActor
	destNode  string
	destActor *baseActor
}

type iActor interface {
	Mailbox() chan payload
	Init()
	receive(<-chan payload)
}

// The baseActor struct
type baseActor struct {
	mailbox chan payload
}

type vector struct {
	x, y, z float64
}

type component interface {
	update(delta float64) error
}

type gameActor struct {
	base       baseActor
	position   vector
	active     bool
	components []component
}

func (actor *baseActor) receive(ch <-chan payload) {
	panic("receive not implemented in actor")
}

func (actor *baseActor) send(msg string) {
	p := payload{}
	p.srcActor = actor
	hostname, _ := os.Hostname()
	p.srcNode = hostname
	p.msg = msg
}
