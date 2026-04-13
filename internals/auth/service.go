package auth

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) CreateUser(req CreateUserRequest) (*User, error) {
	return &User{
		ID:   "123",
		Name: req.Name,
		Email: req.Email,
		Phone: req.Phone,
	}, nil
}