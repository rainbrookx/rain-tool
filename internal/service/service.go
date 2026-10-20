package service

type AllService struct {
	Url        *UrlService
	Identifier *IdentifierService
}

func NewAllService() *AllService {
	var s AllService
	s.Url = NewUrlService()
	s.Identifier = NewIdentifierService()

	return &s
}
