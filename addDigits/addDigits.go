package addDigits

import "fmt"

// Add digital separatelly. For example : AddDigits(123) will return (1 + 2 + 3)
func AddDigits(num int) (ret int, err error) {
	sum := 0
	fmt.Println("Input : ", num)

	for num >= 10 {
		tmp := num % 10
		if sum != 0 {
			fmt.Printf("Process : %d + %d\n", sum, tmp)
		}
		sum += tmp
		num = num / 10
	}
	fmt.Printf("Process : %d + %d\n", sum, num)
	sum += num

	return sum, err
}
