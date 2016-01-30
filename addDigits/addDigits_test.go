package addDigits_test

import (
	"fmt"
	"testing"
	"leetcodeGo/addDigits"
)

func Test1(t *testing.T) {
    res, err  := addDigits.AddDigits(1)

    if (err == nil) {
        fmt.Println("Answer :", res, "\n")
    } else {
		t.Fail()
    }   
}

func Test1357(t *testing.T) {
    res, err  := addDigits.AddDigits(1357)

    if (err == nil) {
        fmt.Println("Answer :", res, "\n")
    } else {
		t.Fail()
    }   
}

func Test2468(t *testing.T) {
    res, err  := addDigits.AddDigits(2468)

    if (err == nil) {
        fmt.Println("Answer :", res, "\n")
    } else {
		t.Fail()
    }   
}

