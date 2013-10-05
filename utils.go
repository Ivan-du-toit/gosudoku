package sudokuLib

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func LoadGrid() SudokuGrid {
	file, err := os.Open("grid.csv")
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}
	defer file.Close()
	reader := csv.NewReader(file)
	var grid = SudokuGrid{}
	for row := range grid {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(fmt.Sprintf("%v", err))
		}
		for col := range grid[row] {
			value, er := strconv.ParseInt(record[col], 10, 8)
			if er != nil {
				panic(fmt.Sprintf("%v", err))
			}
			grid[row][col] = int8(value)
		}
	}
	return grid
}

func (S SudokuGrid) getNextOpenCel() (bool, int, int) {
	for row := range S {
		for col := range S[row] {
			if S[row][col] == 0 {
				return true, row, col
			}
		}
	}
	return false, -1, -1
}

//This function should build the string to allow printing to different channels
func (S SudokuGrid) Print() {

	fmt.Println("")
	for row := range S {
		if row%3 == 0 {
			for i := 0; i < 13; i++ {
				fmt.Print("-")
			}
			fmt.Println("")
		}
		fmt.Print("[")
		for col := range S[row] {
			if col%3 == 0 && col != 0 {
				fmt.Print("|")
			}
			fmt.Print(S[row][col])
		}
		fmt.Println("]")
	}
	fmt.Println("")
}

func (S SudokuGrid) IsSolved() bool {
	if !S.isValid() {
		return false
	}
	for row := range S {
		for col := range S[row] {
			if S[row][col] == 0 {
				return false
			}
		}
	}
	return true
}

func (S SudokuGrid) isValid() bool {
	var rowCheck [10]bool
	var colChecks [9][10]bool
	var blockChecks [3][10]bool
	//Check rows
	for row := range S {
		//Clear row
		for i := range rowCheck {
			rowCheck[i] = false
		}
		//Clear blocks every 3rd row
		if row%3 == 0 {
			for b := range blockChecks {
				for i := range blockChecks[b] {
					blockChecks[b][i] = false
				}
			}
		}

		for col := range S[row] {
			celValue := S[row][col]
			//Skip the checks if the cel is unset
			if celValue == 0 {
				continue
			}

			if rowCheck[celValue] {
				return false
			} else {
				rowCheck[celValue] = true
			}

			//Check cols
			if colChecks[col][celValue] {
				return false
			} else {
				colChecks[col][celValue] = true
			}

			//Check blocks
			if blockChecks[col/3][celValue] {
				return false
			} else {
				blockChecks[col/3][celValue] = true
			}
		}
	}

	return true
}
