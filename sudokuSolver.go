package main

import (
	"fmt"
	"github.com/Ivan-du-toit/gosudoku/sudokuLib"
)

func main() {
	grid := sudokuLib.LoadGrid()
	fmt.Println("Solving:")
	grid.Print()
	found, solution := grid.Solve()
	if found {
		fmt.Println("Solution found:")
		solution.Print()
	} else {
		fmt.Println("No solution found.")
	}
}
