package sudokuLib

import (
	. "launchpad.net/gocheck"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type UtilSuite struct{}

var _ = Suite(&UtilSuite{})

func (s *UtilSuite) Test_getNextOpenCel(c *C) {
	var grid = SudokuGrid{}
	found, row, col := grid.getNextOpenCel()
	c.Assert(found, Equals, true)
	c.Assert(row, Equals, 0)
	c.Assert(col, Equals, 0)

	grid[0][0] = 6
	found, row, col = grid.getNextOpenCel()
	c.Log(col)
	c.Assert(found, Equals, true)
	c.Assert(col, Equals, 1)
	c.Assert(row, Equals, 0)

	//Make sure it fails as expected when the grid is full
	for i := range grid {
		for j := range grid[i] {
			grid[i][j] = 1
		}
	}
	found, row, col = grid.getNextOpenCel()
	c.Assert(found, Equals, false)
	c.Assert(row, Equals, -1)
	c.Assert(col, Equals, -1)
}
