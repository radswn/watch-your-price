package main

import (
	"ceneo"
	"fmt"
)

func main() {
	// test searching
	m := ceneo.SearchForItem("xiaomi redmi note 8t", true)
	for name, url := range m {
		fmt.Println(name + "|" + url)
	}

	// test price checking
	price := ceneo.CheckPrice("https://www.ceneo.pl/87825152")
	fmt.Println(price)
}
