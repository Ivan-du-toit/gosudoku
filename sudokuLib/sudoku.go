package sudoku

import (
	"fmt"
)

type SudokuGrid struct {
	grid [9][9]int;
}

func LoadGrid() (SudokuGrid) {
	var grid = SudokuGrid{};
	return grid;
}

func (S SudokuGrid) Solve() (SudokuGrid){
	return S;
}

func (S SudokuGrid) Print() {
	fmt.Println("Printing 2 come");
}

func (S SudokuGrid) IsSolved() bool {
	return false;
}

func (S SudokuGrid) isValid() bool {
	var rowCheck [10]bool;
	var colChecks [9][10]bool;
	var blockChecks [3][10]bool;
	//Check rows
	for row := range S.grid {
		//Clear row
		for i := range rowCheck {rowCheck[i]=false;}
		//Clear blocks every 3rd row
		if row%3==0 { for b:=range blockChecks { for i:=range blockChecks[b] { blockChecks[b][i] =false;}}}

		for col := range S.grid[row] {
			celValue := S.grid[row][col];
			//Skip the checks if the cel is unset
			if celValue == 0 { continue; }

			if rowCheck[celValue] {
				return false;
			} else { 
				rowCheck[celValue] = true;
			}

			//Check cols
			if colChecks[col][celValue] {
				return false
			} else { colChecks[col][celValue] = true; }

			//Check blocks
			if blockChecks[col%3][celValue] {
				return false;
			} else {
				blockChecks[col%3][celValue] = true;
			}
		}
	}

	return true;
}
