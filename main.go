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
	tileSize   = 32.0
	pixelScale = 4.0
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Snake by Christopher Silva",
		Bounds: pixel.R(-400.0, -400.0, 400.0, 400.0),
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
	var moveDelay int64 = 300000000
	for !win.Closed() {
		// dt = time.Since(last).Seconds()
		// last = time.Now()

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
		if win.JustPressed(pixelgl.KeyQ) {
			moveDelay = moveDelay - 10000000
		}
		if win.JustPressed(pixelgl.KeyEscape) {
			win.SetClosed(true)
		}
		if win.JustPressed(pixelgl.KeyG) {
			snake.Grow()
		}

		snake.Render()

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
