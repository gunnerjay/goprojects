package main

import (
	"fmt"
	"time"

	zmq "github.com/pebbe/zmq4"
	"github.com/veandco/go-sdl2/sdl"
)

const windowX int32 = 1024
const windowY int32 = 768
const targetTicksPerSecond = 60

var delta float64
var entityList []*entity

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

	imgStars, err := sdl.LoadBMP("assets/stars.bmp")
	if err != nil {
		fmt.Println("loading stars bitmap: ", err)
		return
	}
	pf, _ := sdl.AllocFormat(sdl.PIXELFORMAT_RGB24)
	imgStars.SetColorKey(true, sdl.MapRGB(pf, 0, 0, 0))

	starsTex, err := renderer.CreateTextureFromSurface(imgStars)
	if err != nil {
		fmt.Println("creating stars texture: ", err)
		return
	}

	var p = newPlayer(renderer)
	entityList = append(entityList, p)

	p.position.x = float64((windowX - 51) / 2)
	p.position.y = float64((windowY - 53) / 2)

	var starPos = 4096.0 - float64(windowY)
	running := true
	for running {
		frameStartTime := time.Now()
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
		}

		for _, ent := range entityList {
			ent.update(delta)
		}

		renderer.Copy(backgroundTex,
			&sdl.Rect{X: 0, Y: 0, W: windowX, H: windowY},
			&sdl.Rect{X: 0, Y: 0, W: windowX, H: windowY})

		starPos -= 0.06
		var intStarPos = int32(starPos)
		fmt.Println(delta, starPos, intStarPos)
		renderer.Copy(starsTex,
			&sdl.Rect{X: 0, Y: int32(starPos), W: windowX, H: windowY},
			&sdl.Rect{X: 0, Y: 0, W: windowX, H: windowY})
		if starPos < 0 {
			starPos = 4096.0 - float64(windowY)
		}

		for _, ent := range entityList {
			ent.draw(renderer, delta)
		}

		renderer.Present()

		delta = time.Since(frameStartTime).Seconds() * targetTicksPerSecond
	}
}
