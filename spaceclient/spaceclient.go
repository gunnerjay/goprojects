package main

import (
	"fmt"

	zmq "github.com/pebbe/zmq4"
	"github.com/veandco/go-sdl2/sdl"
)

const windowX int32 = 1024
const windowY int32 = 768

func main() {
	requester, _ := zmq.NewSocket(zmq.REQ)
	defer requester.Close()
	requester.Connect("tcp://localhost:5555")
	msg := fmt.Sprintf("Hello ")
	requester.Send(msg, 0)
	reply, _ := requester.Recv(0)
	fmt.Println("reply ", reply)
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("SpaceClient 0.01", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		windowX, windowY, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("Initializing render:", err)
		return
	}

	defer renderer.Destroy()

	img, err := sdl.LoadBMP("assets/corona.bmp")
	if err != nil {
		fmt.Println("loading background sprite: ", err)
		return
	}

	backgroundTex, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		fmt.Println("creating background texture: ", err)
		return
	}

	img, err = sdl.LoadBMP("assets/playership.bmp")
	if err != nil {
		fmt.Println("loading ship sprite: ", err)
		return
	}
	pf, _ := sdl.AllocFormat(sdl.PIXELFORMAT_RGB24)
	img.SetColorKey(true, sdl.MapRGB(pf, 255, 255, 255))

	playerTex, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		fmt.Println("creating ship texture: ", err)
		return
	}

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
		}

		renderer.Copy(backgroundTex,
			&sdl.Rect{X: 0, Y: 0, W: 1024, H: 1024},
			&sdl.Rect{X: 0, Y: 0, W: 1024, H: 1024})

		renderer.Copy(playerTex,
			&sdl.Rect{X: 0, Y: 0, W: 51, H: 53},
			&sdl.Rect{X: (windowX - 51) / 2, Y: (windowY - 53) / 2, W: 51, H: 53})

		renderer.Present()
	}
}
