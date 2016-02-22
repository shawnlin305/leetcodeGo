package removeElement_test

import (
	"fmt"
	"leetcodeGo/removeElement"
	"testing"
)

var TestTable = []struct {
	array []int
	val   int
}{
	{[]int{1, 2, 3, 4, 6, 8}, 3},
	{[]int{3, 3, 3, 3}, 2},
	{[]int{5, 5, 5, 5, 5}, 5},
	{[]int{}, 0},
}

func Test(t *testing.T) {
	/*
		TestTable := [][]int{
			{1, 2, 3, 4},
			{3, 3, 3, 3},
		}
	*/
	for i := 0; i < len(TestTable); i++ {
		fmt.Println(TestTable[i].array, " remove val : ", TestTable[i].val)
		ret_val := removeElement.RemoveElement(TestTable[i].array, len(TestTable[i].array), TestTable[i].val)
		fmt.Println("Updated array : ", TestTable[i].array)
		fmt.Println("Total val : ", ret_val, "\n")
		fmt.Println("================================")
	}
}
