package auth

type Repository interface {
    Create(user *User) error
    GetAll() ([]User, error)
    GetByID(id string) (*User, error)
    Update(user *User) error
    Delete(id string) error
}