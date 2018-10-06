package main

import "fmt"

func main()  {
	jb := []string{"James", "Bond", "Shaken, not stirred"}
	mp := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	ii := [][]string{jb, mp}
	fmt.Println(ii)

	for i,v :=  range ii {
		fmt.Println("Records", i)
		for j,val := range v {
			fmt.Printf("\t index position: %v \t value: \t %v \n", j, val)
		}
	}
}
