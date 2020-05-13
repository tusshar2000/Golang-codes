// I want to know what do u conclude, use pointers for all the types.. even maps,arrays,slices. share your result like this :
// Change in original variable(inside main)
// int-> value : no change , pointer: change
package main

import (
	"fmt"
)

func changeByValue(num int, floatNum float64, string1 string, map1 map[string]int, array1 [6]int, slice1 []int) {
	num = 99
	floatNum = 99.99
	string1 = "Laddha"
	map1["aa"] = 99
	array1[0] = 99
	slice1[0] = 99
}
func changeByPointer(num *int, floatNum *float64, string1 *string, map1 *map[string]int, array1 *[6]int, slice1 *[]int) {
	*num = 99
	*floatNum = 99.99
	*string1 = "Laddha"
	(*map1)["gg"] = 100
	array1[0] = 100
	(*slice1)[0] = 100
}

func main() {
	var num int = 10
	var floatNum float64 = 10.5
	var string1 string = "Tusshar"
	map1 := map[string]int{
		"aa": 3,
		"gg": 2,
	}
	array1 := [6]int{2, 3, 5, 7, 11, 13}
	slice1 := []int{2, 3, 5, 7, 11, 13}

	changeByValue(num, floatNum, string1, map1, array1, slice1)
	fmt.Println(num)
	fmt.Println(floatNum)
	fmt.Println(string1)
	fmt.Println(map1)
	fmt.Println(array1)
	fmt.Println(slice1)

	changeByPointer(&num, &floatNum, &string1, &map1, &array1, &slice1)
	fmt.Println(num)
	fmt.Println(floatNum)
	fmt.Println(string1)
	fmt.Println(map1)
	fmt.Println(array1)
	fmt.Println(slice1)

	fmt.Println()
}
