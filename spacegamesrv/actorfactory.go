package main

//StartActor returns an Actor instance
func StartActor() *Actor {
	ch := make(chan string, 1)
	actor := &Actor{actionChan: ch}
	go actor.actorLoop(ch)
	return actor
}
