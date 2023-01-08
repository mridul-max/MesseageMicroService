package Services

import (
	"MesseageMicroService/restApi/Domain"
)

type MessageService interface {
	GetAllMessages() ([]Domain.Message, error)
}

// "constructor" like function
// whereby we pass in the repo (interface) as a dependency
type DefaultMessageService struct {
	repo Domain.MessageRepository
}

// receiver function -attaches it as a method to a class
func (s DefaultMessageService) GetAllMessages() ([]Domain.Message, error) {

	//Once again we talk to the interface
	return s.repo.FindAll()
}

// Helper function to instantiate customer service
func NewMessageService(repo Domain.MessageRepository) DefaultMessageService {
	return DefaultMessageService{repo}
}
