package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ParseFile reads a file and returns a slice of Tetromino pieces.
func ParseFile(path string) ([]Tetromino, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var pieces []Tetromino
	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		line := strings.TrimRight(scanner.Text(), "\r")
		if line == "" {
			if len(lines) > 0 {
				t, err := parsePiece(lines)
				if err != nil {
					return nil, err
				}
				pieces = append(pieces, t)
				lines = nil
			}
			continue
		}
		lines = append(lines, line)
		if len(lines) == 4 {
			// expect either blank line or EOF after this
			// but allow continuing
		}
	}
	if len(lines) > 0 {
		t, err := parsePiece(lines)
		if err != nil {
			return nil, err
		}
		pieces = append(pieces, t)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	if len(pieces) == 0 {
		return nil, fmt.Errorf("no pieces")
	}
	return pieces, nil
}

// parsePiece validates and converts four lines into a Tetromino.
func parsePiece(lines []string) (Tetromino, error) {
	if len(lines) != 4 {
		return Tetromino{}, fmt.Errorf("invalid piece line count")
	}
	grid := make([][]rune, 4)
	hashCount := 0
	for i, line := range lines {
		if len(line) != 4 {
			return Tetromino{}, fmt.Errorf("invalid line length")
		}
		row := make([]rune, 4)
		for j, r := range line {
			if r != '#' && r != '.' {
				return Tetromino{}, fmt.Errorf("invalid character")
			}
			row[j] = r
			if r == '#' {
				hashCount++
			}
		}
		grid[i] = row
	}
	if hashCount != 4 {
		return Tetromino{}, fmt.Errorf("invalid number of blocks")
	}
	if !connected(grid) {
		return Tetromino{}, fmt.Errorf("piece not connected")
	}

	// compute bounding box
	minX, minY := 3, 3
	maxX, maxY := 0, 0
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			if grid[y][x] == '#' {
				if x < minX {
					minX = x
				}
				if y < minY {
					minY = y
				}
				if x > maxX {
					maxX = x
				}
				if y > maxY {
					maxY = y
				}
			}
		}
	}

	var blocks []Point
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if grid[y][x] == '#' {
				blocks = append(blocks, Point{X: x - minX, Y: y - minY})
			}
		}
	}

	t := Tetromino{
		Blocks: blocks,
		Width:  maxX - minX + 1,
		Height: maxY - minY + 1,
	}
	return t, nil
}

// connected checks whether all '#' cells are orthogonally connected.
func connected(grid [][]rune) bool {
	visited := make([][]bool, 4)
	for i := range visited {
		visited[i] = make([]bool, 4)
	}
	var queue []Point
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			if grid[y][x] == '#' {
				queue = append(queue, Point{X: x, Y: y})
				visited[y][x] = true
				break
			}
		}
		if len(queue) > 0 {
			break
		}
	}
	if len(queue) == 0 {
		return false
	}
	dirs := []Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	count := 0
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		count++
		for _, d := range dirs {
			nx, ny := p.X+d.X, p.Y+d.Y
			if nx >= 0 && nx < 4 && ny >= 0 && ny < 4 && !visited[ny][nx] && grid[ny][nx] == '#' {
				visited[ny][nx] = true
				queue = append(queue, Point{X: nx, Y: ny})
			}
		}
	}
	return count == 4
}
