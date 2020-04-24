package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	fmt.Println("Hello World")

	fmt.Println(2 > 7)
	fmt.Println(2.5 / 7)

	fmt.Println(strings.Title("shi shuan"))

	fmt.Println(reflect.TypeOf(4))
	fmt.Println(reflect.TypeOf(4.9))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf("jaja"))

	var quantity int
	quantity = 4
	//can be done this way also
	var length, width = 1.2, 6.4
	var custname string = "Tusshar"
	fmt.Println(custname, "has", quantity, "of tiles with dimension", length, width)
	//when variable name is initialized and no value is assigned, int,float var has
	// 0 value, string has empty string, bool has false.

	//Another way of performing same above stuff, this does initializing and assigning
	// at the same time, cant be done on predeclared variables.
	quantit := 4
	lengt, widt := 1.2, 6.4
	cusname := "Tusshar"
	fmt.Println(cusname, "has", quantit, "of tiles with dimension", lengt, widt)

	inte := 2
	flot := 4.0
	//we need type conversions if we want to apply some functions in which they
	// interact directly.
	ans := int(flot) * inte
	fmt.Println(ans)

	// A package is a group of related functions and other code.

	// Before you can use a package’s functions within a Go file, you need to import that
	//  package.

	// A string is a series of bytes that usually represent text characters.

	// A rune represents a single text character.

	// Go’s two most common numeric types are int, which holds integers, and float64,
	//  which holds floating-point numbers.

	// The bool type holds Boolean values, which are either true or false.

	// A variable is a piece of storage that can contain values of a specified type.

	// If no value has been assigned to a variable, it will contain the zero value for its type.
	//  Examples of zero values include 0 for int or float64 variables, or "" for string
	//  variables.

	// You can declare a variable and assign it a value at the same time using a := short
	//  variable declaration.

	// A variable, function, or type can only be accessed from code in other packages if its
	//  name begins with a capital letter.

	// The go fmt command automatically reformats source files to use Go standard

	// formatting. You should run go fmt on any code that you plan to share with others.

	// The go build command compiles Go source code into a binary format that computers
	//  can execute.

	// The go run command compiles and runs a program without saving an executable file in
	//  the current directory.

}
