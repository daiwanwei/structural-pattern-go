package service

import "fmt"

type BillingService struct {
}

func (service *BillingService) Reply(problem string) (answer string, err error) {
	switch problem {
	default:
		return fmt.Sprintf("your problem is %s,answer is %s", problem, "waiting!!!!"), nil
	}
}
