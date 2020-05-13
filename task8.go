package main

import (
	"fmt"
)

func checkType(anyType interface{}) {
	switch anyType.(type) {
	case string:
		fmt.Println("Argument is a string")
	case float64:
		fmt.Println("Argument is a float64")
	case float32:
		fmt.Println("Argument is a float32")
	case int:
		fmt.Println("Argument is an integer")
	default:
		fmt.Println("Other than my case")
	}
}
func main() {
	var i interface{} = 10
	// i=20
	// i=true
	storeString, ok := i.(int)
	fmt.Println(storeString, ok)
	checkType(5.6)
	//var num float32 = 10.2
	checkType(true)
}
