package main

import (
	"log"
	"math"

	"github.com/faiface/pixel"
)

const (
	east  = iota
	north = iota
	west  = iota
	south = iota
)

// don't want to turn north<->south or east<->west
// olddir + 2 % 4 = newdir
// east(0) + 2 % 4 = 2(west)
// west(2) + 2 % 4 = 0(east)
// north(1) + 2 % 4 = 3(south)
// south(3) + 2 % 4 = 1(north)

type Section struct {
	Position  pixel.Vec
	Direction int
}

// Snake is the player character
type Snake struct {
	NextDirection int
	Sections      []Section
	Sprites       map[string]*pixel.Sprite
}

// NewSnake creates the player character
func NewSnake() *Snake {
	sections := []Section{
		Section{Position: pixel.V(0.0, 0.0), Direction: north},
		Section{Position: pixel.V(0.0, -32.0), Direction: north},
		Section{Position: pixel.V(0.0, -64.0), Direction: north},
		Section{Position: pixel.V(0.0, -96.0), Direction: north},
		Section{Position: pixel.V(0.0, -128.0), Direction: north},
	}
	sprites := make(map[string]*pixel.Sprite)
	spritesheet, err := loadPicture("snake.png")
	if err != nil {
		log.Println(err)
		return nil
	}

	sprites["tail"] = pixel.NewSprite(spritesheet, pixel.R(0, 24, 8, 32))
	sprites["straight"] = pixel.NewSprite(spritesheet, pixel.R(8, 24, 16, 32))
	sprites["corner"] = pixel.NewSprite(spritesheet, pixel.R(16, 24, 24, 32))
	sprites["head"] = pixel.NewSprite(spritesheet, pixel.R(24, 24, 32, 32))

	return &Snake{NextDirection: north, Sprites: sprites, Sections: sections}
}

// Update the snakes position
func (s *Snake) Update() {
	// move all sections forward one
	// each sections gets the position and direction of the section ahead of it
	// the head section gets position infront of it
	// check if head intersects apple, if so eat it
	for i := len(s.Sections) - 1; i > 0; i-- {
		s.Sections[i].Position = s.Sections[i-1].Position
		s.Sections[i].Direction = s.Sections[i-1].Direction
	}
	s.Sections[len(s.Sections)-1].Direction = s.Sections[len(s.Sections)-2].Direction
	s.Sections[0].Direction = s.NextDirection
	switch s.Sections[0].Direction {
	case north:
		s.Sections[0].Position = s.Sections[0].Position.Add(pixel.V(0.0, tileSize))
	case east:
		s.Sections[0].Position = s.Sections[0].Position.Add(pixel.V(tileSize, 0.0))
	case south:
		s.Sections[0].Position = s.Sections[0].Position.Add(pixel.V(0.0, -tileSize))
	case west:
		s.Sections[0].Position = s.Sections[0].Position.Add(pixel.V(-tileSize, 0.0))
	}
}

// Render the snake
func (s *Snake) Render() {
	// loop through sections and render them
	for i := len(s.Sections) - 2; i > 0; i-- {
		mat := pixel.IM
		mat = mat.Rotated(pixel.ZV, float64(s.Sections[i].Direction)*math.Pi/2)
		mat = mat.ScaledXY(pixel.ZV, pixel.V(pixelScale, pixelScale))
		mat = mat.Moved(s.Sections[i].Position)
		if s.Sections[i].Direction != s.Sections[i-1].Direction {
			s.Sprites["corner"].Draw(win, mat)
		} else {
			s.Sprites["straight"].Draw(win, mat)
		}
	}
	mat := pixel.IM
	mat = mat.Rotated(pixel.ZV, float64(s.Sections[len(s.Sections)-1].Direction)*math.Pi/2)
	mat = mat.ScaledXY(pixel.ZV, pixel.V(pixelScale, pixelScale))
	mat = mat.Moved(s.Sections[len(s.Sections)-1].Position)
	s.Sprites["tail"].Draw(win, mat)

	mat = pixel.IM
	mat = mat.Rotated(pixel.ZV, float64(s.Sections[0].Direction)*math.Pi/2)
	mat = mat.ScaledXY(pixel.ZV, pixel.V(pixelScale, pixelScale))
	mat = mat.Moved(s.Sections[0].Position)
	s.Sprites["head"].Draw(win, mat)
}

// Turn the snake
func (s *Snake) Turn(direction int) {
	if (s.Sections[1].Direction+2)%4 != direction {
		s.NextDirection = direction
	}
}

// Eat an apple
func (s *Snake) Eat() {
	// check if there is an apple at the head location
	// if so, eat it and add a section
}
