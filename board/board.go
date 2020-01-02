// Package board provides an implementation of a Boggle board. Representing
// the dice.
package board

import (
	"fmt"
	"math/rand"
)

// HEIGHT is the universal height of a boggle board
const HEIGHT = 4

// WIDTH is the universal width of a boggle board
const WIDTH = 4

// Error for the board package.
type Error struct {
	message string
}

func (e *Error) Error() string {
	return e.message
}

// Board represents an individual instance of a Boggle board.
type Board [HEIGHT][WIDTH]string

// There are 16 dice in a standard Boggle game, these are the
// individual dice.
var dice = [16][6]string{
	{"a", "a", "c", "i", "o", "t"}, // AACIOT
	{"a", "h", "m", "o", "r", "s"}, // AHMORS
	{"e", "g", "k", "l", "u", "y"}, // EGKLUY
	{"a", "b", "i", "l", "t", "y"}, // ABILTY
	{"a", "c", "d", "e", "m", "p"}, // ACDEMP
	{"e", "g", "i", "n", "t", "v"}, // EGINTV
	{"g", "i", "l", "r", "u", "w"}, // GILRUW
	{"e", "l", "p", "s", "t", "u"}, // ELPSTU
	{"d", "e", "n", "o", "s", "w"}, // DENOSW
	{"a", "c", "e", "l", "r", "s"}, // ACELRS
	{"a", "b", "j", "m", "o", "q"}, // ABJMOQ
	{"e", "e", "f", "h", "i", "y"}, // EEFHIY
	{"e", "h", "i", "n", "p", "s"}, // EHINPS
	{"d", "k", "n", "o", "t", "u"}, // DKNOTU
	{"a", "d", "e", "n", "v", "z"}, // ADENVZ
	{"b", "i", "f", "o", "r", "x"}} // BIFORX

// New creates a random boggle board.
func New() *Board {
	board := &Board{}

	order := rand.Perm(16)
	pos := 0
	for x := 0; x < HEIGHT; x++ {
		for y := 0; y < WIDTH; y++ {
			board[x][y] = dice[order[pos]][rand.Int()%6]
			pos++
		}
	}

	return board
}

// Height returns the height of the board
func (b *Board) Height() int {
	return HEIGHT
}

// Width returns the width of the board.
func (b *Board) Width() int {
	return WIDTH
}

// GetAt returns the character as a speciifc position.
// And Error is returned if an invalid coordinate is specfified.
func (b *Board) GetAt(x, y int) (string, error) {
	if x < 0 || x >= HEIGHT {
		return "", &Error{fmt.Sprintf("X value of %v is out of bounds must be 0 =< x < %v", x, HEIGHT)}
	}
	if y < 0 || y >= WIDTH {
		return "", &Error{fmt.Sprintf("Y value of %v is out of bounds must be 0 =< y < %v", y, WIDTH)}
	}
	return b[x][y], nil
}

// PrintBoard writes the board to stdout.
func (b *Board) PrintBoard() {
	fmt.Println("Board:")
	for x := 0; x < HEIGHT; x++ {
		for y := 0; y < WIDTH; y++ {
			fmt.Printf("%s ", b[x][y])
		}
		fmt.Printf("\n")
	}
}
