package service

type AllService struct {
	Url *UrlService
}

func NewAllService() *AllService {
	var s AllService
	s.Url = NewUrlService()

	return &s
}
