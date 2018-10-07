package main

import "fmt"

type Person struct {
	First            string
	Last             string
	FavoriteIceCream []string
}

func main()  {

	p1 := Person{
		First: "Abhishek",
		Last:  "Singh",
		FavoriteIceCream: []string{
			"chocolate",
			"martini",
			"rum and coke",
		},
	}
	p2 := Person{
		First: "Bobby",
		Last: "Singh",
		FavoriteIceCream: []string{
			"strawberry",
			"vanilla",
			"capuccino",
		},
	}
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	for i,v := range p1.FavoriteIceCream {
		fmt.Println(i,v)
	}

	fmt.Println(p2.First)
	fmt.Println(p2.Last)
	for i,v := range p2.FavoriteIceCream {
		fmt.Println(i,v)
	}
}