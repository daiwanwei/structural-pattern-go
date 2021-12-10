package main

import (
	"errors"
	"fmt"
)

type Ingredient interface {
	GetIngredient() (ingredient string, err error)
}

type Pizza struct {
}

func (i *Pizza) GetIngredient() (ingredient string, err error) {
	return fmt.Sprintf("%s", "pizza"), nil
}

type Meat struct {
	Ingredient Ingredient
}

func (i *Meat) GetIngredient() (ingredient string, err error) {
	if i.Ingredient == nil {
		return "", errors.New("an Ingredient is needed in the Ingredient field of the Meat")
	}
	another, err := i.Ingredient.GetIngredient()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s %s", another, "meat"), nil
}

type Tomato struct {
	Ingredient Ingredient
}

func (i *Tomato) GetIngredient() (ingredient string, err error) {
	if i.Ingredient == nil {
		return "", errors.New("an Ingredient is needed in the Ingredient field of the tomato")
	}
	another, err := i.Ingredient.GetIngredient()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s %s", another, "tomato"), nil
}
