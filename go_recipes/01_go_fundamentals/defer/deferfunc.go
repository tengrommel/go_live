package main

import (
	"os"
	"io/ioutil"
	"fmt"
)

func ReadFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil{
		panic(err)
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}

func main() {
	f, _ := ReadFile("test.txt")
	fmt.Printf("%s", f)
	fmt.Println(string(f))
}
