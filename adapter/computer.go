package main

import "fmt"

type Mac struct {
}

func (computer *Mac) ConnectWithTypeC() {
	fmt.Println("connect to mac with type C")
}
