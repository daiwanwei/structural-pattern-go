package main

import "fmt"

func main() {
	pizza := Meat{&Tomato{&Pizza{}}}
	ingredient, err := pizza.GetIngredient()
	if err != nil {
		panic(err)
	}
	fmt.Println(ingredient)
}
