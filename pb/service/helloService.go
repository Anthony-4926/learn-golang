package service

type HelloService struct {
}

func NewHelloService() *HelloService {
	return new(HelloService)
}
