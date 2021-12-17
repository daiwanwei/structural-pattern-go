package main

import "fmt"

type Police interface {
	Executive()
}

type PoliceManager struct {
	Name    string
	Members []Police
}

func (police *PoliceManager) Executive() {
	fmt.Printf("manager(%s) executive\n", police.Name)
	for _, m := range police.Members {
		m.Executive()
	}
}

type PoliceOfficer struct {
	Name string
}

func (police *PoliceOfficer) Executive() {
	fmt.Printf("officer(%s) executive\n", police.Name)
}
