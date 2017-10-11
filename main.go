package main

import (
	"fmt"
	"log"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

var (
	win         *pixelgl.Window
	spritesheet pixel.Picture
	snake       *Snake
	apple       *Apple
	dt          float64
	moveDelay   int64 = 300000000
	tileSize          = 32.0
	pixelScale        = 4.0
	openSpots   []pixel.Vec
	gameOver    bool
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Snake by Christopher Silva",
		Bounds: pixel.R(-tileSize*12.5, -tileSize*12.5, tileSize*12.5, tileSize*12.5),
		VSync:  true,
	}

	var err error
	win, err = pixelgl.NewWindow(cfg)
	if err != nil {
		log.Fatal(err)
	}

	spritesheet, err = loadPicture("snake.png")
	if err != nil {
		log.Fatal(err)
	}

	snake = NewSnake()
	apple = NewApple()

	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	basicTxt := text.New(pixel.V(-124.0, 0.0), basicAtlas)
	fmt.Fprintln(basicTxt, "Game Over")
	fmt.Fprintln(basicTxt, "[R]etry?")

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

		if gameOver {
			basicTxt.Draw(win, pixel.IM.Scaled(basicTxt.Orig, 4))
			if win.JustPressed(pixelgl.KeyR) {
				gameOver = false
				snake = NewSnake()
				apple.Regen()
				moveDelay = 300000000
			}
		}

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
