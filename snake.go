package main

import (
	"fmt"
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

// what direction is to the right/left?
/*
	start north(1)
	right is east(0) -1
	left is west(2) +1

	start east(0)
	right is south(3) +3
	left is north(1) +1

	start south(3)
	right is west(2) -1
	left is east(0) -3

	start west(2)
	right is north(1) -1
	left is south(3) +1
*/

type Section struct {
	Position  pixel.Vec
	Direction int
}

// Snake is the player character
type Snake struct {
	NextDirection  int
	GrowNextUpdate bool
	Head           *Section
	Sections       []Section
	Sprites        map[string]*pixel.Sprite
}

// NewSnake creates the player character
func NewSnake() *Snake {
	sections := []Section{
		Section{Position: pixel.V(0.0, 0.0), Direction: north},
		Section{Position: pixel.V(0.0, -tileSize), Direction: north},
		Section{Position: pixel.V(0.0, -tileSize*2), Direction: north},
	}
	sprites := make(map[string]*pixel.Sprite)
	spritesheet, err := loadPicture("snake.png")
	if err != nil {
		log.Println(err)
		return nil
	}

	sprites["tail"] = pixel.NewSprite(spritesheet, pixel.R(0, 24, 8, 32))
	sprites["straight"] = pixel.NewSprite(spritesheet, pixel.R(8, 24, 16, 32))
	sprites["rcorner"] = pixel.NewSprite(spritesheet, pixel.R(16, 24, 24, 32))
	sprites["head"] = pixel.NewSprite(spritesheet, pixel.R(24, 24, 32, 32))
	sprites["lcorner"] = pixel.NewSprite(spritesheet, pixel.R(8, 16, 16, 24))

	return &Snake{NextDirection: north, GrowNextUpdate: false, Head: &sections[0], Sections: sections, Sprites: sprites}
}

// Update the snakes position
func (s *Snake) Update() {
	// move all sections forward one
	// each sections gets the position and direction of the section ahead of it
	// the head section gets position infront of it
	// if the snake should grow, do it after moving positions forward one
	var newSection *Section
	if s.GrowNextUpdate {
		newSection = &Section{Position: s.Sections[len(s.Sections)-1].Position, Direction: s.Sections[len(s.Sections)-1].Direction}
	}
	for i := len(s.Sections) - 1; i > 0; i-- {
		s.Sections[i].Position = s.Sections[i-1].Position
		s.Sections[i].Direction = s.Sections[i-1].Direction
	}
	if newSection != nil {
		s.Sections = append(s.Sections, *newSection)
		s.GrowNextUpdate = false
		s.Head = &s.Sections[0]
	}
	s.Sections[len(s.Sections)-1].Direction = s.Sections[len(s.Sections)-2].Direction
	s.Head.Direction = s.NextDirection
	switch s.Head.Direction {
	case north:
		s.Head.Position = s.Head.Position.Add(pixel.V(0.0, tileSize))
	case east:
		s.Head.Position = s.Head.Position.Add(pixel.V(tileSize, 0.0))
	case south:
		s.Head.Position = s.Head.Position.Add(pixel.V(0.0, -tileSize))
	case west:
		s.Head.Position = s.Head.Position.Add(pixel.V(-tileSize, 0.0))
	}

	// fill slice of open spots

	// bounds check
	if s.Head.Position.X > win.Bounds().Max.X ||
		s.Head.Position.Y > win.Bounds().Max.Y ||
		s.Head.Position.X < win.Bounds().Min.X ||
		s.Head.Position.Y < win.Bounds().Min.Y {
		fmt.Println("Game Over! (out of bounds)")
	}

	// apple check
	if s.Head.Position.X == apple.Position.X && s.Head.Position.Y == apple.Position.Y {
		s.Eat()
	}

	// suicide check
	for _, section := range s.Sections[1:] {
		if s.Head.Position.X == section.Position.X && s.Head.Position.Y == section.Position.Y {
			fmt.Println("Game Over! (suicide)")
		}
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
			// todo: refactor this to not be so big
			switch s.Sections[i].Direction {
			case north:
				if s.Sections[i-1].Direction == east {
					s.Sprites["rcorner"].Draw(win, mat)
				} else {
					s.Sprites["lcorner"].Draw(win, mat)
				}
			case east:
				if s.Sections[i-1].Direction == south {
					s.Sprites["rcorner"].Draw(win, mat)
				} else {
					s.Sprites["lcorner"].Draw(win, mat)
				}
			case south:
				if s.Sections[i-1].Direction == west {
					s.Sprites["rcorner"].Draw(win, mat)
				} else {
					s.Sprites["lcorner"].Draw(win, mat)
				}
			case west:
				if s.Sections[i-1].Direction == north {
					s.Sprites["rcorner"].Draw(win, mat)
				} else {
					s.Sprites["lcorner"].Draw(win, mat)
				}
			}
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
	mat = mat.Rotated(pixel.ZV, float64(s.Head.Direction)*math.Pi/2)
	mat = mat.ScaledXY(pixel.ZV, pixel.V(pixelScale, pixelScale))
	mat = mat.Moved(s.Head.Position)
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
	fmt.Println("apple eaten!")
	apple.Regen()
	s.Grow()
}

func (s *Snake) Grow() {
	s.GrowNextUpdate = true
}
