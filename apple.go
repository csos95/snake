package main

import (
	"log"
	"math/rand"

	"github.com/faiface/pixel"
)

// Apple is the goal that must be eaten
type Apple struct {
	Position pixel.Vec
	Sprite   *pixel.Sprite
}

// NewApple creates a new apple
func NewApple() *Apple {
	x := float64(rand.Intn(24)-12) * tileSize
	y := float64(rand.Intn(24)-12) * tileSize
	position := pixel.V(x, y)

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

// Regen the apple (give it a new location)
func (a *Apple) Regen() {
	x := float64(rand.Intn(24)-12) * tileSize
	y := float64(rand.Intn(24)-12) * tileSize
	a.Position = pixel.V(x, y)
}
