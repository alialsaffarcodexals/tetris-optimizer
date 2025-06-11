package main

// Point represents a coordinate on the board
// x is column, y is row

type Point struct {
	X int
	Y int
}

// Tetromino represents a single Tetris piece normalized to its
// minimal bounding box.
type Tetromino struct {
	Blocks []Point // list of 4 blocks relative to top-left
	Width  int     // width of bounding box
	Height int     // height of bounding box
}
