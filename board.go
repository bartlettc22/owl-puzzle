package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

type board struct {
	Solution        []*Square
	Squares         [9]*Square
	CurrentPosition int8
}

// Recursive function that takes the current solution and unused squares
// and determines the position/orientation of the next empty position
func (b *board) solve() bool {

	// Our breakout condition (board is solved!)
	if b.CurrentPosition == 8 {
		return true
	}

	// Looping through each unused square and trying it out
	for i := 0; i < 9; i++ {
		if !b.Squares[i].Used {
			b.addSquare(i)
			// Here we loop through and flip the squares around if the don't fit
			for j := 0; j < 4; j++ {
				if b.isLastSquareValid() {
					log.Infof("Position %d is valid", b.CurrentPosition)
					solved := b.solve()
					if solved {
						// the rest of the puzzle is solved, we're done here
						return true
					}
					log.Infof("Square %d in position %d cannot be solved", i, b.CurrentPosition)
					break
				} else {
					log.Infof("Position %d is invalid", b.CurrentPosition)
				}
				// Dont' really need to rotate on the last loop (save ourselves some time)
				if j < 3 {
					b.Squares[i].rotate()
				}
			}
			b.removeSquare()
		}
	}

	return false
}

// AddSquare adds a square to the next open position on the board and marks it used
func (b *board) addSquare(i int) {
	b.Solution = append(b.Solution, b.Squares[i])
	b.CurrentPosition++
	b.Squares[i].Used = true
	log.Infof("Added square %d to position %d", i, b.CurrentPosition)
}

// RemoveSquare removes a square in the current position on the board and marks it unused
func (b *board) removeSquare() {
	var squareRemoved *Square
	squareRemoved, b.Solution = b.Solution[b.CurrentPosition], b.Solution[:b.CurrentPosition]
	log.Infof("Removed square %s in position %d", squareRemoved.Id, b.CurrentPosition)
	b.CurrentPosition--
	squareRemoved.Used = false

}

// Checks that the last piece placed on the board is valid
// |0|1|2|
// |3|4|5|
// |6|7|8|
func (b *board) isLastSquareValid() bool {
	iterations++
	position := len(b.Solution) - 1

	switch position {
	case 0:
		return true
	case 1:
		return isMatch(b.Solution[0].right(), b.Solution[1].left())
	case 2:
		return isMatch(b.Solution[1].right(), b.Solution[2].left())
	case 3:
		return isMatch(b.Solution[0].bottom(), b.Solution[3].top())
	case 4:
		return isMatch(b.Solution[3].right(), b.Solution[4].left()) && isMatch(b.Solution[1].bottom(), b.Solution[4].top())
	case 5:
		return isMatch(b.Solution[4].right(), b.Solution[5].left()) && isMatch(b.Solution[2].bottom(), b.Solution[5].top())
	case 6:
		return isMatch(b.Solution[3].bottom(), b.Solution[6].top())
	case 7:
		return isMatch(b.Solution[6].right(), b.Solution[7].left()) && isMatch(b.Solution[4].bottom(), b.Solution[7].top())
	case 8:
		// If one side matches, they all match, no need to check the final side
		return isMatch(b.Solution[7].right(), b.Solution[8].left())
	default:
		log.Fatal("Should never get here")
	}

	// Should never get here
	return false
}

// Print prints the current board position and orientation
func (b *board) print() {

	fmt.Println()
	fmt.Println("First Piece Orientation")
	fmt.Println("-------------------------------------------")
	fmt.Printf("| %-11s | %-11s | %-11s |\n", "  ", b.Solution[0].Sides[1].Color+"-"+b.Solution[0].Sides[1].End, "  ")
	fmt.Println("-------------------------------------------")
	fmt.Printf("| %-11s | %-11s | %-11s |\n", b.Solution[0].Sides[0].Color+"-"+b.Solution[0].Sides[0].End, b.Solution[0].Id, b.Solution[0].Sides[2].Color+"-"+b.Solution[0].Sides[2].End)
	fmt.Println("-------------------------------------------")
	fmt.Printf("| %-11s | %-11s | %-11s |\n", "  ", b.Solution[0].Sides[3].Color+"-"+b.Solution[0].Sides[3].End, "  ")
	fmt.Println("-------------------------------------------")
	fmt.Println()
	fmt.Println("Piece Positions")
	fmt.Println("-------------")
	fmt.Printf("| %s | %s | %s |\n", b.Solution[0].Id, b.Solution[1].Id, b.Solution[2].Id)
	fmt.Println("-------------")
	fmt.Printf("| %s | %s | %s |\n", b.Solution[3].Id, b.Solution[4].Id, b.Solution[5].Id)
	fmt.Println("-------------")
	fmt.Printf("| %s | %s | %s |\n", b.Solution[6].Id, b.Solution[7].Id, b.Solution[8].Id)
	fmt.Println("-------------")
	fmt.Println()
}

// Determines if two sides match (same color, different end)
func isMatch(side1 *Side, side2 *Side) bool {
	return side1.Color == side2.Color && side1.End != side2.End
}
