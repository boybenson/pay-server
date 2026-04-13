package auth

import (
	"errors"

	"github.com/boybenson/pay-server/pkg"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s *Service) CreateUser(body CreateUserRequest) (*User, error) {

	user := &User{
		Name:  body.Name,
		Email: body.Email,
		Phone: body.Phone,
	}

	hashedPassword := pkg.HashPassword(body.Password)
	user.Password = hashedPassword

	err := s.db.Create(user).Error

	if err != nil {
		return nil, err
	}

	return &User{
		ID:    "123",
		Name:  body.Name,
		Email: body.Email,
		Phone: body.Phone,
	}, nil
}

func (s *Service) SignIn(body SignInRequest) (*SignInUser, error) {
	var user = User{Email: body.Email}

	err := s.db.Where("email = ?", body.Email).First(&user).Error

	if err != nil {
		return nil, err
	}

	if !pkg.CheckPasswordHash(body.Password, user.Password) {
		return nil, errors.New("incorrect password")
	}

	return &SignInUser{
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone,
	}, nil
}
