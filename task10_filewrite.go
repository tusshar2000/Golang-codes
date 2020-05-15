package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data := []byte("Hello from go")
	err := ioutil.WriteFile("ourFile1.txt", data, 0644) //replaces all text with this.
	checkError(err)
	fmt.Println("Done writing.")

	f, er := os.Create("ourFile2.txt")
	checkError(er)
	bytesWritten, e := f.WriteString("Hello")
	checkError(e)
	fmt.Println("Bytes written are:", bytesWritten)
	fmt.Println("Done writing.")
}

func checkError(err error) {
	if err != nil {
		//panic(err)
		fmt.Println("Error occured is:", err) //this doesnt stop the program.
	}
}
