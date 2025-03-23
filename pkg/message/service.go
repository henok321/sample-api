package message

type MessageService interface {
	Create(message *Message) error
	FindAll() ([]*Message, error)
	FindByID(id int) (*Message, error)
	Update(message *Message) error
	Delete(id int) error
}

type messageService struct {
	repo MessageRepository
}

func newMessageService(repo MessageRepository) MessageService {
	return &messageService{repo: repo}
}

func (s *messageService) Create(message *Message) error {
	return s.repo.Create(message)
}

func (s *messageService) FindAll() ([]*Message, error) {

	return s.repo.FindAll()
}

func (s *messageService) FindByID(id int) (*Message, error) {
	return s.repo.FindByID(id)
}

func (s *messageService) Update(message *Message) error {
	return s.repo.Update(message)
}

func (s *messageService) Delete(id int) error {

	return s.repo.Delete(id)
}
