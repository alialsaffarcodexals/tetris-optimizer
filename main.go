package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR")
		return
	}
	pieces, err := ParseFile(os.Args[1])
	if err != nil {
		fmt.Println("ERROR")
		return
	}
	board := Solve(pieces)
	if board == nil {
		fmt.Println("ERROR")
		return
	}
	fmt.Print(board.String())
}
