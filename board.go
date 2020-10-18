package main

import (
	"errors"
	"strings"
)


// Board struct
type Board struct {
	grid  [3][3]string
	cells [9]uint
	won bool
}

// Update update a board with a marker
func (b *Board) Update(cell uint, marker string) error {
	row := cell / 3
	col := cell % 3

	b.grid[row][col] = marker

	return errors.New("Invalid turn")
}

// Draw prints out the board as a 3 x 3 grid
func (b *Board) Draw() string {
	bs := ""

	for _, row := range b.grid {
		rs := row[0:]

		bs += " "
		bs += strings.Join(rs, " | ")
		bs += "\n-----------\n"
	}

	return bs
}

// New initalize a new 3 x 3 game board
func New() *Board {
	b := Board{}
	b.grid[0] = [3]string{"0", "1", "2"}
	b.grid[1] = [3]string{"3", "4", "5"}
	b.grid[2] = [3]string{"6", "7", "8"}

	b.cells = [9]uint{0, 1, 2, 3, 4, 5, 6, 7, 8}

	return &b
}
