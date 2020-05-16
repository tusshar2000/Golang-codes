package main

import "fmt"

type byteSize float64

const (
	_           = iota
	KB byteSize = 1 << (10 * iota)
	MB
	GB
	TB
)

func main() {
	var fileSize byteSize = 5000000.0
	fmt.Println(KB, MB, GB, TB)
	fmt.Printf("%0.2fKB", fileSize/KB)

}
