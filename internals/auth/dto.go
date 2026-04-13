package auth

type CreateUserRequest struct {
    Name  string `json:"name" validate:"required"`
    Email string `json:"email" validate:"required,email"`
    Phone string `json:"phone" validate:"required"`
}

type UpdateUserRequest struct {
    Name  *string `json:"name"`
    Email *string `json:"email"`
    Phone *string `json:"phone"`
}