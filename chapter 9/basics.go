package main

import "fmt"

type litres float64
type gallons float64

func main() {
	var carFuel gallons
	var busFuel litres
	carFuel = gallons(20.0)
	busFuel = litres(240.0)
	fmt.Println(carFuel, busFuel)
	// //can't do this
	// carFuel:=litres(20.0) type gallons assigned
	// busFuel:=gallons(240.0)

	// carFuel:=gallons(litres(20.0)) //legal but incorrect
	carFuel = gallons(litres(20.0) * 0.264)
	busFuel = litres(gallons(63.0) * 3.785)
	fmt.Printf("Gallons:%0.1f Litres:%0.1f\n", carFuel, busFuel)

	//can use all underlying arithmetic
	fmt.Println(litres(10) + litres(20))

	// 	But defined types cannot be used in operations together with values of a
	//different type, even if the other type has the same underlying type.
	//Again, this is to protect developers from accidentally mixing the two types.
	//fmt.Println(litres(10)+gallons(20))
}
