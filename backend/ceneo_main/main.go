package main

import (
	"ceneo"
	"fmt"
)

func main() {
	m := ceneo.SearchForItem("xiaomi redmi note 8t")
	for name, url := range m {
		fmt.Println(name + "|" + url)
	}
}
