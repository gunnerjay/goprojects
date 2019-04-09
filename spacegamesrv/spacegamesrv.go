//
//  Hello World server.
//  Binds REP socket to tcp://*:5555
//  Expects "Hello" from client, replies with "World"
//

package main

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
	"time"
)

func main() {
	a := StartActor()
	time.Sleep(time.Second)
	a.update()

	fmt.Println("listening")
	//  Socket to talk to clients
	responder, _ := zmq.NewSocket(zmq.ROUTER)
	defer responder.Close()
	responder.Bind("tcp://*:5555")

	for {
		//  Wait for next request from client
		msg, metaData, _ := responder.RecvBytesWithMetadata(0)
		fmt.Println("Received ", msg, metaData)

		//  Do some 'work'
		time.Sleep(time.Second)

		//  Send reply back to client
		reply := "World"
		responder.Send(reply, 0)
		fmt.Println("Sent ", reply)
	}
}
