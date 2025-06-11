package main

// Board represents the placement grid for tetrominoes.
type Board struct {
	Size  int
	Cells [][]rune
}

// NewBoard creates a new empty board of given size.
func NewBoard(size int) *Board {
	b := &Board{Size: size, Cells: make([][]rune, size)}
	for i := 0; i < size; i++ {
		row := make([]rune, size)
		for j := range row {
			row[j] = '.'
		}
		b.Cells[i] = row
	}
	return b
}

// Fit returns true if the piece fits at the given position.
func (b *Board) Fit(t Tetromino, row, col int) bool {
	if row+t.Height > b.Size || col+t.Width > b.Size {
		return false
	}
	for _, p := range t.Blocks {
		r := row + p.Y
		c := col + p.X
		if b.Cells[r][c] != '.' {
			return false
		}
	}
	return true
}

// Place writes the piece letter at the given position.
func (b *Board) Place(t Tetromino, row, col int, letter rune) {
	for _, p := range t.Blocks {
		b.Cells[row+p.Y][col+p.X] = letter
	}
}

// Remove clears the piece from the board.
func (b *Board) Remove(t Tetromino, row, col int) {
	for _, p := range t.Blocks {
		b.Cells[row+p.Y][col+p.X] = '.'
	}
}

// String returns the board as printable string.
func (b *Board) String() string {
	out := make([]byte, 0, b.Size*(b.Size+1))
	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			out = append(out, byte(b.Cells[i][j]))
		}
		out = append(out, '\n')
	}
	return string(out)
}
