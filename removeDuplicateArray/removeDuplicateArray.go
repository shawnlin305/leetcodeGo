package removeDuplicateArray

func RemoveDupArray(array []int, len int) (rt_val int) {
	if len == 0 || len == 1 {
		return len
	}

	index := 1
	cur := 0

	for ; index < len; index++ {
		if array[cur] != array[index] {
			cur++
			array[cur] = array[index]
		}
	}
	return (cur + 1)
}
