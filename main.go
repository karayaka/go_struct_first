package main

import "fmt"

func main() {

	human := Human{
		Liver: Liver{
			talk: "Konuş",
			walk: "yürü",
		},
		Name:   "Çağrı",
		Age:    33,
		wiht:   10,
		height: 20,
	}
	print(human.calculate("de"))
	fmt.Println(human.walk)
}
