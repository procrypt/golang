package main

import "fmt"

func main()  {

	m := make(map[string][]string)
	m["bond_james"]      = []string{`Shaken, not stirred`, `Martinis`, `Women`}
	m["moneypenny_miss"] = []string{ `James Bond`, `Literature`, `Computer Science`}
	m["no_dr"]           = []string{`Being evil`, `Ice cream`, `Sunsets`}

	for _,v := range m {
		for i, val := range v {
			fmt.Println(i, val)
		}
	}
}
