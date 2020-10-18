package main

import (
	"errors"
	"fmt"
	"strings"
)

// MarkableCells cells that can be marked
const MarkableCells = 9

// Board struct
type Board struct {
	grid           [3][3]string
	markedCells    []uint
	availableCells []uint
}

// Update update a board with a marker
func (b *Board) Update(cell uint, marker string) error {
	if len(b.markedCells) == MarkableCells {
		return errors.New("game over")
	}

	found := false
	for _, ac := range b.availableCells {
		if ac == cell {
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("cell %d is not available", cell)
	}

	row := cell / 3
	col := cell % 3
	b.grid[row][col] = marker

	b.markedCells = append(b.markedCells, cell)

	// remove the markedCell from the array

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

	b.availableCells = []uint{0, 1, 2, 3, 4, 5, 6, 7, 8}

	return &b
}
