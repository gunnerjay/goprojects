//
//  Weather update server
//  Binds PUB socket to tcp://*:5556
//  Publishes random weather updates
//

package main

import (
	"math/rand"
	"os"
	"time"

	zmq "github.com/pebbe/zmq4"

	"fmt"
)

func main() {
	hostname, _ := os.Hostname()
	fmt.Println("publishing from", hostname, os.Getpid())
	context, _ := zmq.NewContext()
	publisher, _ := zmq.NewSocket(zmq.PUB)
	defer context.Term()
	defer publisher.Close()
	e := publisher.Connect("tcp://localhost:5559")
	if e != nil {
		panic(e)
	}

	// e = publisher.Bind("ipc://weather.ipc")
	// if e != nil {
	// 	panic(e)
	// }

	rand.Seed(time.Now().UnixNano())

	for {
		var zipcode, temperature, relhumidity int
		zipcode = rand.Intn(100000)
		temperature = rand.Intn(125) - 80
		relhumidity = rand.Intn(50) + 10

		var env = fmt.Sprintf("%s:%d", hostname, os.Getpid())
		var msg = fmt.Sprintf("%05d %d %d", zipcode, temperature, relhumidity)
		_, e = publisher.SendMessage([]byte(env), []byte(msg))
		if e != nil {
			panic(e)
		}
	}
}
