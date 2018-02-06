package main

import (
	"io/ioutil"
	"fmt"
	"os"
	//"io"
	"bufio"
)

func check(e error) {
	if e != nil{
		panic(e)
	}
}

func main() {
	bs, err := ioutil.ReadFile("pointers.go")
	check(err)
	fmt.Print(string(bs))

	file, err := os.Open("pointers.go")
	check(err)
	b1 := make([]byte,5)
	n1, err := file.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	//o2, err := file.Seek(6, 0)
	//check(err)
	//b2 := make([]byte, 2)
	//n2, err := file.Read(b2)
	//check(err)
	//fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))
	//
	//o3, err := file.Seek(6, 0)
	//check(err)
	//b3 := make([]byte, 2)
	//n3, err := io.ReadAtLeast(file, b3, 2)
	//check(err)
    //fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))
	//
	//
    //r4 :=  bufio.NewReader(file)
    //b4, err := r4.Peek(100)
    //check(err)
    //fmt.Printf("100 bytes: %s\n", string(b4))

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
    	line := scanner.Text()
    	fmt.Println(">>>", line)
	}

}