//
//   Weather proxy listens to weather server which is constantly
//   emitting weather data
//   Binds SUB socket to tcp://*:5556
//
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	context, _ := zmq.NewContext()
	socket, _ := context.NewSocket(zmq.SUB)
	defer context.Term()
	defer socket.Close()

	var temps []string
	var err error
	var temp int64
	totalTemp := 0
	filter := "59937"

	// find zipcode
	if len(os.Args) > 1 { // ./wuclient 85678
		filter = string(os.Args[1])
	}

	////  Subscribe to just one zipcode (whitefish MT 59937) //
	fmt.Printf("Collecting updates from weather server for %s…\n", filter)
	//socket.SetSubscribe(filter) // <== REQUIRED
	socket.SetSubscribe("ZeusMini:19700") // <== REQUIRED
	socket.Connect("tcp://localhost:5560")

	for i := 0; i < 101; i++ {
		// found temperature point
		fmt.Printf("Found temperature point for %s…\n", filter)
		address, _ := socket.Recv(0)
		fmt.Println("address coming thru", address)
		datapt, _ := socket.Recv(0)
		temps = strings.Split(string(datapt), " ")
		temp, err = strconv.ParseInt(temps[1], 10, 64)
		if err == nil {
			//// Invalid string //
			totalTemp += int(temp)
		}
	}

	fmt.Printf("Average temperature for zipcode %s was %dF \n\n", filter, totalTemp/100)
}
