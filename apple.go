package main

import (
	"log"

	"github.com/faiface/pixel"
)

// Apple is the goal that must be eaten
type Apple struct {
	Position pixel.Vec
	Sprite   *pixel.Sprite
}

// NewApple creates a new apple
func NewApple() *Apple {
	position := pixel.V(32.0, 32.0)

	spritesheet, err := loadPicture("snake.png")
	if err != nil {
		log.Println(err)
		return nil
	}

	sprite := pixel.NewSprite(spritesheet, pixel.R(16, 16, 24, 24))

	return &Apple{Position: position, Sprite: sprite}
}

// Render the apple
func (a *Apple) Render() {
	mat := pixel.IM
	mat = mat.ScaledXY(pixel.ZV, pixel.V(pixelScale, pixelScale))
	mat = mat.Moved(a.Position)
	a.Sprite.Draw(win, mat)
}

func (a *Apple) Regen() {
	a.Position = pixel.V(64.0, 64.0)
}
