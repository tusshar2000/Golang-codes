package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	content, err := ioutil.ReadFile("ourFile.txt")
	checkError(err)
	//fmt.Println("Content of out file are",content)we get utf-8 values in slices
	result := string(content)
	fmt.Println("Content of the file is", result)

	f, er := os.Open("ourFile.txt")
	defer f.Close()
	checkError(er)
	bucket := make([]byte, 1000) //we just set an arbitarily huge number
	bytesRead, e := f.Read(bucket)
	checkError(e)
	fmt.Println("Content of the file(linited) is", string(bucket[:bytesRead])) //we didn't
	//wanted spaces so just printed the slice upto the no of bytes read.
	fmt.Println("Bytes Read:", bytesRead)
}
func checkError(err error) {
	if err != nil {
		//panic(err)
		fmt.Println("Error occured is:", err) //this doesnt stop the program.
	}
}
