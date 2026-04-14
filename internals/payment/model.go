package payment

import "github.com/boybenson/pay-server/internals/auth"

type PaymentCode struct {
	ID     string    `json:"id" gorm:"primaryKey"`
	UserID string    `json:"userId"`
	User   auth.User `gorm:"foreignKey:UserID"`
	Code   string    `json:"code"`
}
