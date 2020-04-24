package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var now time.Time = time.Now()
	var year int = now.Year()
	fmt.Println(year)

	bro := "G# R#CKS"
	replacer := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(bro)
	fmt.Println(fixed)

}
