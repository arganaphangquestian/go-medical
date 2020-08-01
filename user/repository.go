package user

import "context"

// Repository User
type Repository interface {
	Close()

	AddUser(ctx context.Context, name string, email string, address string, roleID uint32, genderID uint32, bloodID uint32, BirthOfDate string, contact string) error
	GetUsers(ctx context.Context) ([]User, error)
	GetUserByID(ctx context.Context, id string) (*User, error)
}
