package sudoku

import (
	"fmt"
    "os"
    "io"
    "encoding/csv"
    "strconv"
)

type SudokuGrid struct {
	grid [9][9]int8;
}

func LoadGrid() (SudokuGrid) {
    file, err := os.Open("grid.csv");
    if err != nil {
        //fmt.Println("Can't read filei grid.csv, error: ", err);
        panic(fmt.Sprintf("%v", err));
    }
    defer file.Close();
    reader := csv.NewReader(file);
	var grid = SudokuGrid{};
    for row:= range grid.grid {
        record, err := reader.Read();
        if err == io.EOF {
            break;
        } else if err != nil {
            //fmt.Println("Error: ", err);
            panic(fmt.Sprintf("%v", err));
        }
        for col := range grid.grid[row] {
            value, er := strconv.ParseInt(record[col], 10, 8);
            if er != nil {
                //fmt.Println("Error: ", err);
                panic(fmt.Sprintf("%v", err));
            }
            grid.grid[row][col] = int8(value);
        }
    }
	return grid;
}

func (S SudokuGrid) Solve() (bool, SudokuGrid){
    /*S.Print();
    i := 0;
    fmt.Scanf("%d", &i);*/
    if !S.isValid() {
        return false, S;
    }
    found, row, col :=  S.getNextOpenCel();
    if !found {
        if S.isValid() {
            //This is the solution!!
            return true, S;
        }
        return false, S;
    }
    for i:=1; i<10; i++ {
        S.grid[row][col] = int8(i);
        worked, solvedGrid := S.Solve();
        if worked {
            return worked, solvedGrid;
        }
    }
    
	return false, S;
}

func (S SudokuGrid) getNextOpenCel() (bool, int, int) {
    for row := range S.grid {
        for col := range S.grid[row] {
            if S.grid[row][col] == 0 {
                return true, row, col;
            }
        }
    }
    return false, 0, 0; 
}

func (S SudokuGrid) Print() {

	fmt.Println("");
    for row := range S.grid {
        if row%3==0 { 
            for i:=0; i<13; i++ { fmt.Print("-"); }
            fmt.Println("");
        }
        fmt.Print("[");
        for col := range S.grid[row] {
            if col%3 == 0 && col != 0 {
                fmt.Print("|");
            }
            fmt.Print(S.grid[row][col]);
        }
        fmt.Println("]");
    }
	fmt.Println("");
}

func (S SudokuGrid) IsSolved() bool {
    if !S.isValid() {
        return false;
    }
    for row := range S.grid {
        for col := range S.grid[row] {
            if S.grid[row][col] == 0 {
                return false;
            }
        }
    }
	return true;
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
		if row%3==0 {
            //fmt.Println("cleaning blocks");
            for b:=range blockChecks { 
                for i:=range blockChecks[b] { 
                    blockChecks[b][i] =false;
                }
            }
        }

		for col := range S.grid[row] {
			celValue := S.grid[row][col];
			//Skip the checks if the cel is unset
			if celValue == 0 { continue; }

			if rowCheck[celValue] {
                //fmt.Println("row fail");
				return false;
			} else { 
				rowCheck[celValue] = true;
			}

			//Check cols
			if colChecks[col][celValue] {
                //fmt.Println("col fail");
				return false
			} else { 
                colChecks[col][celValue] = true;
            }

			//Check blocks
            //fmt.Println("Checking:", celValue, "in block", col/3);
            //print2D(blockChecks);
			if blockChecks[col/3][celValue] {
                //fmt.Println("block fail in block:", col/3, "with value:", celValue);
				return false;
			} else {
				blockChecks[col/3][celValue] = true;
			}
		}
	}

	return true;
}

func print2D(ar [3][10]bool) {
    for i:=range ar {
        for j:=range ar[i] {
            if j%3==0 {
                fmt.Println("");
            }
            fmt.Print(",", ar[i][j]);
        }
        fmt.Println("");
    }
}
