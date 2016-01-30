package addDigits_test

import (
	"fmt"
	"testing"
	"leetcodeGo/addDigits"
)

var TestTable = []struct {
	n int
	expected int
}{
	{0, 0},
	{1, 1},
	{12, 3},
	{135, 9},
	{246, 12},
	{987, 24},
	{962, 17},
}

func TestBomb(t *testing.T) {


	for _, num := range TestTable {
		res, err  := addDigits.AddDigits(num.n)

		if (err == nil && res == num.expected) {
			fmt.Println("Answer :", res, "\n")
		} else {
			t.Error("Wrong Answer pair : {", num.n, " : ", num.expected, "}")
		}
	}
}

