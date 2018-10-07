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

	m := map[string]Person{
		p1.Last: p1,
		p2.Last: p2,
	}

	for _,v := range m {
		fmt.Println(v.First)
		fmt.Println(v.Last)
		for i, val := range v.FavoriteIceCream {
			fmt.Println(i,val)
		}
		fmt.Println("-------")
	}
}