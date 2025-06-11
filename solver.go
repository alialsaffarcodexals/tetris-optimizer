package main

import "math"

// Solve tries to fit all tetromino pieces into the smallest square possible.
func Solve(pieces []Tetromino) *Board {
	if len(pieces) == 0 {
		return nil
	}
	size := int(math.Ceil(math.Sqrt(float64(len(pieces) * 4))))
	for {
		board := NewBoard(size)
		if placePieces(board, pieces, 0) {
			return board
		}
		size++
	}
}

func placePieces(b *Board, pieces []Tetromino, index int) bool {
	if index == len(pieces) {
		return true
	}
	piece := pieces[index]
	letter := rune('A' + index)
	for r := 0; r <= b.Size-piece.Height; r++ {
		for c := 0; c <= b.Size-piece.Width; c++ {
			if b.Fit(piece, r, c) {
				b.Place(piece, r, c, letter)
				if placePieces(b, pieces, index+1) {
					return true
				}
				b.Remove(piece, r, c)
			}
		}
	}
	return false
}
