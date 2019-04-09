package main

import "fmt"

// The Actor struct
type Actor struct {
	actionChan chan string
}

func (actor *Actor) update() {
	fmt.Println("called into update")
	actor.actionChan <- "update"
}

func (actor *Actor) actorLoop(ch <-chan string) {
	fmt.Println("in actor loop")
	for {
		action := <-ch
		fmt.Println("the action is ", action)
	}
}
