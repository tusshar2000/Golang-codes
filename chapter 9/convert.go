package main

import (
	"fmt"
)

type litres float64
type gallons float64
type millilitres float64

func (l litres) ToGallons() gallons {
	return gallons(l * 0.264)
}
func (m millilitres) ToGallons() gallons {
	return gallons(m * 0.000264)
}
func (g gallons) ToLitres() litres {
	return litres(g * 3.785)
}
func (g gallons) ToMillilitres() millilitres {
	return millilitres(g * 3785.41)
}
func main() {
	soda := litres(2)
	fmt.Printf("%0.3f liters=%0.3f gallons\n", soda, soda.ToGallons())
	water := millilitres(500)
	fmt.Printf("%0.3f milliliters=%0.3f gallons\n", water, water.ToGallons())
	milk := gallons(2)
	fmt.Printf("%0.3f gallons=%0.3f Litres\n", milk, milk.ToLitres())
	fmt.Printf("%0.3f gallons=%0.3f Millilitres\n", milk, milk.ToMillilitres())
}
