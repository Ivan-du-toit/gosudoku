package main

import (
	"fmt"
	"./sudokuLib"
)

func main() {
	grid := sudoku.LoadGrid();
	solution := grid.Solve();
	if solution.IsSolved() {
		fmt.Println("solution found!!!");
		solution.Print();
	} else {
		fmt.Println("No solution found");
	}
}
