package removeElement

func RemoveElement(array []int, length int, val int) (rt_length int) {
	if length == 0 {
		return 0
	}

	cur := 0
	for i := 0; i < length; i++ {
		if array[i] == val {
			continue
		}
		array[cur] = array[i]
		cur++
	}

	return cur
}
