package main

import "strconv"

func main() {
	var officers []Police

	for i := 0; i < 10; i++ {
		officers = append(officers, &PoliceOfficer{
			"ann" + strconv.Itoa(i),
		})
	}
	manager := PoliceManager{
		Name:    "wei",
		Members: officers,
	}

	manager2 := PoliceManager{
		Name: "ann",
		Members: []Police{
			&manager,
		},
	}
	manager2.Executive()
}
