package main

import (
	"./sudokuLib"
	"fmt"
)

func main() {
	grid := sudoku.LoadGrid()
	fmt.Println("Solving:")
	grid.Print()
	found, solution := grid.Solve()
	if found {
		fmt.Println("solution found:")
		solution.Print()
	} else {
		fmt.Println("No solution found")
	}
}
