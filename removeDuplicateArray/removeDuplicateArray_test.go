package removeDuplicateArray_test

import (
	"fmt"
	"leetcodeGo/removeDuplicateArray"
	"testing"
)

/* Test array should be sorted */
var TestTable = []struct {
	array []int
}{
	{[]int{}},
	{[]int{1}},
	{[]int{2, 2}},
	{[]int{1, 1, 2, 2, 2, 4, 4, 5, 5}},
}

func TestFun(t *testing.T) {
	//func TestBomb(t *testing.T) {
	for i := 0; i < len(TestTable); i++ {
		fmt.Println("Input array : ", TestTable[i].array)
		rt_val := removeDuplicateArray.RemoveDupArray(TestTable[i].array, len(TestTable[i].array))
		fmt.Println("Return array: ", TestTable[i].array[:rt_val], "\nArray size : ", rt_val)
		fmt.Println("=====================================")
	}
}
