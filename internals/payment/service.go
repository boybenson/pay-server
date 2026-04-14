package payment

import (
	"errors"

	"github.com/boybenson/pay-server/internals/auth"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s *Service) CreatePaymentCode(body CreatePaymentCodeRequest) (*PaymentCode, error) {
	var existing PaymentCode

	err := s.db.First(&existing, "code = ?", body.Code).Error

	if err == nil {
		return nil, errors.New("payment code already exists")
	}

	paymentCode := PaymentCode{
		Code: body.Code,
		User: auth.User{ID: body.User},
	}

	if err := s.db.Create(&paymentCode).Error; err != nil {
		return nil, err
	}

	return &paymentCode, nil
}
