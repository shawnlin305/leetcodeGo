package moveZeros

import "fmt"

func MvZeros(nums []int) {
	zeroIndex := 0
	numIndex := 0
	arryLen := len(nums)
	fmt.Println("arryLen : ", arryLen)

	for (zeroIndex < arryLen && numIndex < arryLen) {
		if nums[zeroIndex] != 0 {
			zeroIndex++
			numIndex = zeroIndex
			continue
		}
		if nums[numIndex] == 0 {
			numIndex++
			continue
		}
		nums[zeroIndex] = nums[numIndex];
		nums[numIndex] = 0;
		numIndex++;
		zeroIndex++;
	}

	for i := 0; i < arryLen; i++ {
		fmt.Println(nums[i])
	}
}


