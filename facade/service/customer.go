package service

type CustomerService struct {
	Billing Service
	Product Service
}

func (service *CustomerService) ReplyForBilling(problem string) (answer string, err error) {
	return service.Billing.Reply(problem)
}

func (service *CustomerService) ReplyForProduct(problem string) (answer string, err error) {
	return service.Product.Reply(problem)
}

type Service interface {
	Reply(problem string) (answer string, err error)
}
