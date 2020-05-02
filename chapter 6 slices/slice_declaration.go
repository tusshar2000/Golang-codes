package main

import "fmt"

func main() {
	notes := []string{"do", "re", "mi"}
	primes := []int{
		2,
		3,
		5,
	}
	fmt.Println(notes)
	fmt.Println(primes)

	underlyingarray := [5]int{1, 2, 3, 4, 5}
	slicearray := underlyingarray[1:3]
	fmt.Println(slicearray)
	//here it works as call by reference, so any changes in slice makes same changes in array

	//when you append on a slice changes aren't ref
	slicearray = append(slicearray, 101)
	fmt.Println(underlyingarray)
	fmt.Println(slicearray)
	slicearray[1] = 100
	fmt.Println(underlyingarray)
	fmt.Println(slicearray)

}
