package maths

import (
	//"fmt"
	"testing"
)

//name convention of func matters.
func TestAdd(t *testing.T) {
	actual := add(20, 30)
	var expected int64 = 50
	if actual != expected {
		t.Error("Actual is:", actual, "but expected is:", expected)
	}
}

func TestAddBulk(t *testing.T) {
	var list = []struct {
		first    int
		second   int
		expected int64
	}{
		{
			20, 30, 50,
		},
		{-20, 30, 10},
		{-20, -30, -50},
	}
	//fmt.Println(list)
	for _, str := range list {
		if actual := add(str.first, str.second); actual != str.expected {
			t.Error("Actual is:", actual, "but expected is:", str.expected)
		}

	}
}
