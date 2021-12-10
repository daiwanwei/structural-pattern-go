package main

import (
	"fmt"
	"structural-pattern-go/facade/service"
)

func main() {
	billing := service.BillingService{}
	product := service.ProductService{}
	customer := service.CustomerService{
		Billing: &billing, Product: &product,
	}
	reply, err := customer.ReplyForProduct("product fuck")
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)
	reply, err = customer.ReplyForBilling("billing fuck")
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)
}
