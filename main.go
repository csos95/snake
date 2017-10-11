package main

import (
	"log"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var (
	win        *pixelgl.Window
	snake      *Snake
	apple      *Apple
	dt         float64
	moveDelay  int64 = 300000000
	tileSize         = 32.0
	pixelScale       = 4.0
	openSpots  []pixel.Vec
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Snake by Christopher Silva",
		Bounds: pixel.R(-tileSize*12, -tileSize*12, tileSize*12, tileSize*12),
		VSync:  true,
	}

	var err error
	win, err = pixelgl.NewWindow(cfg)
	if err != nil {
		log.Fatal(err)
	}

	snake = NewSnake()
	apple = NewApple()

	var background = colornames.Forestgreen
	last := time.Now()
	for !win.Closed() {
		win.Clear(background)

		// 1 000 000 000ns = 1s
		if time.Since(last).Nanoseconds() > moveDelay {
			last = time.Now()
			snake.Update()
		}

		if win.JustPressed(pixelgl.KeyUp) {
			snake.Turn(north)
		}
		if win.JustPressed(pixelgl.KeyRight) {
			snake.Turn(east)
		}
		if win.JustPressed(pixelgl.KeyDown) {
			snake.Turn(south)
		}
		if win.JustPressed(pixelgl.KeyLeft) {
			snake.Turn(west)
		}
		if win.JustPressed(pixelgl.KeyEscape) {
			win.SetClosed(true)
		}

		snake.Render()
		apple.Render()

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
