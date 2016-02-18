package moveZeros_test

import (
	"fmt"
	"leetcodeGo/moveZeros"
	"testing"
)

func TestMvZeros(t *testing.T) {
	_sampleArray := [][]int{
		{0, 1, 2, 3, 4, 0, 5, 6},
		{7, 6, 4, 0, 2, 2, 0, 1},
	}

	for i := 0; i < len(_sampleArray); i++ {
		fmt.Println("Input array as :", _sampleArray[i])
		moveZeros.MvZeros(_sampleArray[i])
		fmt.Println("Output array as:", _sampleArray[i])
		fmt.Println("--------------------------------")
	}
}
