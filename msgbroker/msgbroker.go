//
//

package main

import (
	"os"

	zmq "github.com/pebbe/zmq4"

	"fmt"
)

func main() {
	hostname, _ := os.Hostname()
	fmt.Println("brokering", hostname, os.Getpid())

	context, _ := zmq.NewContext()
	defer context.Term()

	frontend, _ := zmq.NewSocket(zmq.XSUB)
	defer frontend.Close()
	e := frontend.Bind("tcp://*:5559")
	if e != nil {
		panic(e)
	}

	backend, _ := zmq.NewSocket(zmq.XPUB)
	defer backend.Close()
	e = backend.Bind("tcp://*:5560")
	if e != nil {
		panic(e)
	}

	zmq.Proxy(frontend, backend, nil)
}
