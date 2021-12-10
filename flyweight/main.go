package main

import (
	"fmt"
	"structural-pattern-go/flyweight/forests"
)

func main() {
	tree, err := forests.NewTree("Ficus")
	if err != nil {
		panic(err)
	}
	tree.NewLocation(1, 1)
	fmt.Println(tree)
}
