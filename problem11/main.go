package main

import "fmt"

type SubrectangleQueries struct {
	row1 int
	row2 int
	col1 int
	col2 int
}

func Constructor(rectangle [][]int) SubrectangleQueries {
	return SubrectangleQueries{}
}

func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {

}

func (this *SubrectangleQueries) GetValue(row int, col int) int {
	return 0
}

func main() {
	fmt.Print("Hello World")
	s1 := SubrectangleQueries{}
	s1.row1 = 1
	s1.row2 = 2
	s1.col1 = 3
	s1.col2 = 4

	//this.UpdateSubrectangle(1, 2, 3, 4)

}

/**
 * Your SubrectangleQueries object will be instantiated and called as such:
 * obj := Constructor(rectangle);
 * obj.UpdateSubrectangle(row1,col1,row2,col2,newValue);
 * param_2 := obj.GetValue(row,col);
 */
