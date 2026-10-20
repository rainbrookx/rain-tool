package service

type AllService struct {
	Url        *UrlService
	Identifier *IdentifierService
	Uuid       *UuidService
}

func NewAllService() *AllService {
	var s AllService
	s.Url = NewUrlService()
	s.Identifier = NewIdentifierService()
	s.Uuid = NewUuidService()

	return &s
}
