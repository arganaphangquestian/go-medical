package user

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Service for user Module
type Service interface {
	AddUser(ctx context.Context, name string, email string, address string, roleID uint32, genderID uint32, bloodID uint32, birthOfDate string, contact string) error
	GetUsers(ctx context.Context) ([]User, error)
	GetUserByID(ctx context.Context, id string) (*User, error)
}

// User struct Model
type User struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name"`
	Email       string             `json:"email"`
	Address     string             `json:"address"`
	RoleID      uint32             `json:"role_id"`
	GenderID    uint32             `json:"gender_id"`
	BloodID     uint32             `json:"blood_id"`
	BirthOfDate string             `json:"birth_of_date"`
	Contact     string             `json:"contact"`
	CreatedAt   string             `json:"created_at"`
	UpdatedAt   string             `json:"updated_at"`
}
