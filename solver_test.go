package main

import "testing"

func TestSolveSimple(t *testing.T) {
	pieces := []Tetromino{
		{Blocks: []Point{{0, 0}, {0, 1}, {0, 2}, {0, 3}}, Width: 1, Height: 4},
		{Blocks: []Point{{0, 0}, {1, 0}, {0, 1}, {1, 1}}, Width: 2, Height: 2},
	}
	b := Solve(pieces)
	if b.Size == 0 {
		t.Fatalf("no solution")
	}
	if b.Size != 4 {
		t.Fatalf("expected board size 4 got %d", b.Size)
	}
}
