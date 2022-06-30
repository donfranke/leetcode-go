// https://leetcode.com/problems/subrectangle-queries/

package main

type SubrectangleQueries struct {
	rectangle [][]int
	size      int
	empty     bool
}

func Constructor(rectangle [][]int) SubrectangleQueries {
	size := len(rectangle)

	return SubrectangleQueries{
		rectangle: rectangle,
		size:      size,
		empty:     size <= 0,
	}
}

func (s SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {

}

func (s SubrectangleQueries) GetValue(row int, col int) int {
	return 0
}

func main() {

	// ["SubrectangleQueries","getValue","updateSubrectangle","getValue","getValue","updateSubrectangle","getValue","getValue"]
	// [[[[1,2,1],[4,3,4],[3,2,1],[1,1,1]]],[0,2],[0,0,3,2,5],[0,2],[3,1],[3,0,3,2,10],[3,1],[0,2]]

	//this.UpdateSubrectangle(1, 2, 3, 4)

}

/**
 * Your SubrectangleQueries object will be instantiated and called as such:
 * obj := Constructor(rectangle);
 * obj.UpdateSubrectangle(row1,col1,row2,col2,newValue);
 * param_2 := obj.GetValue(row,col);
 */
