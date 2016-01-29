package main

import "fmt"
import "leetcodeGo/addDigits"

func main() {
	fmt.Println("Input a number : \n")
	var testValue int
	fmt.Scanf("%d", &testValue)
	res, err  := addDigits.AddDigits(testValue)

	if err == nil {
		fmt.Println(res)
	} else {
		fmt.Println("Error")
	}
}
