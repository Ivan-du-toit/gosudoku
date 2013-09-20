package sudokuLib

type SudokuGrid [9][9]int8

func (S SudokuGrid) Solve() (bool, SudokuGrid) {
	if !S.isValid() {
		return false, S
	}
	found, row, col := S.getNextOpenCel()
	if !found {
		if S.isValid() {
			//This is the solution!!
			return true, S
		}
		return false, S
	}
	for i := 1; i < 10; i++ {
		S[row][col] = int8(i)
		worked, solvedGrid := S.Solve()
		if worked {
			return worked, solvedGrid
		}
	}

	return false, S
}

// func (s SudokuGrid) findAllSolutions() []SudokuGrid {

// }
