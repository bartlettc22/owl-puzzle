package main

import (
	log "github.com/sirupsen/logrus"
)

// Square describes one of the puzzle pieces
type Square struct {
	Id string `json:"id"`

	// Side 0 is considered to be the leftmost side, 1 is the top, 2 is the right and 3 is the bottom
	Sides [4]*Side `json:"sides"`

	// Is set to true if the square is currently being used in the solution
	Used bool
}

// Side describes a size of a square
type Side struct {
	Color string `json:"color"`
	End   string `json:"end"`
}

// Rotate rotates the sides of the square (counterclockwise)
func (s *Square) rotate() {
	log.Infof("Rotating square %s", s.Id)
	x := s.Sides[0]
	s.Sides = [4]*Side{s.Sides[1], s.Sides[2], s.Sides[3], x}
}

func (s *Square) left() *Side {
	return s.Sides[0]
}

func (s *Square) top() *Side {
	return s.Sides[1]
}

func (s *Square) right() *Side {
	return s.Sides[2]
}

func (s *Square) bottom() *Side {
	return s.Sides[3]
}
