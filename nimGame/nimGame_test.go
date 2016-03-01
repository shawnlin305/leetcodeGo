package nimGame_test

import (
	"fmt"
	"leetcodeGo/nimGame"
	"testing"
)

var TestMap = map[int]bool{
	1: true,
	2: true,
	4: false,
}

func TestNimGame(t *testing.T) {
	for key, expect := range TestMap {
		ans := nimGame.CanWinNim(key)
		fmt.Println("Input :", key, " Expect :", expect, " Return :", ans)
		if nimGame.CanWinNim(key) != expect {
			t.Error("Incorrect return :", ans)
		}
		fmt.Println("================================")
	}
}
