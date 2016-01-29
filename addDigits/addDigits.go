package addDigits

import "fmt"

func AddDigits(num int) (ret int, err error) {
	sum := 0

	for num >= 10 {
		fmt.Println("Process : ", num)
		sum += num % 10
		num = num / 10
	}
	sum += num

	return sum, err
}
