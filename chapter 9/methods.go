package main

import (
	"fmt"
)

type MyType string

func (m MyType) sayHi() {
	fmt.Println("Hi from", m)
}

func (m MyType) withReturn() int {
	return len(m)
}
func (m MyType) method() {
	fmt.Println("Hi from method")
}
func (m *MyType) pointermethod() {
	fmt.Println("Hi from pointer method")
}
func main() {
	value := MyType("Tusshar")
	value.sayHi()
	fmt.Println(value.withReturn())
	anotherValue := MyType("New")
	anotherValue.sayHi()

	value1 := MyType("Hi")
	pointer := &value1
	value1.method()
	value1.pointermethod()
	pointer.method()
	pointer.pointermethod()

}
