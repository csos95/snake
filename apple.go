package main

import (
	"github.com/faiface/pixel"
)

// Apple is the goal that must be eaten
type Apple struct {
	Position pixel.Vec
	Sprite   *pixel.Sprite
}

// NewApple creates a new apple
func NewApple() *Apple {
	position := pixel.V(0.0, 0.0)
	return &Apple{Position: position}
}

// Render the apple
func (a *Apple) Render() {

}
