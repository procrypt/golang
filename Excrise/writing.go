package main

import (
	"io/ioutil"
	"os"
	"fmt"
	"bufio"
)

func check(e error) {
	if e!= nil {
		panic(e)
	}
}

func main() {

	data := []byte("Hello fron Golang!")
	err := ioutil.WriteFile("sample.txt", data, 0644)
	check(err)

	file, err := os.Create("Sample2.txt")
	check(err)
	defer file.Close()
	data2 := []byte{65, 98, 104, 105, 115, 104, 101, 107, 32, 80, 114, 97, 116, 97, 112, 32, 83, 105, 110, 103, 104,10}
	n,err:= file.Write(data2)
	check(err)
	fmt.Printf("%d bytes[], %s\n", n, string(data2))

	// Write String
	n1, err := file.WriteString("Hello from golang!\n")
	check(err)
	fmt.Printf("%d bytes wrote\n", n1)
	file.Sync()

	// Buffered writer
	w := bufio.NewWriter(file)
	n2, err := w.WriteString("Hello again from Golang!")
	fmt.Printf("%d bytes wrote\n", n2)

	w.Flush()
}
