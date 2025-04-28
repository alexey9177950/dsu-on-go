package main

import (
	"dsuOnGo/dsu"
	"fmt"
)

func main() {
	data := dsu.NewDsu(50)
	fmt.Println(data.Unite(5, 7))
}
