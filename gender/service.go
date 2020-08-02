package gender

import "context"

// Service for user Module
type Service interface {
	AddGender(ctx context.Context, name string, description string) error
	GetGenders(ctx context.Context) ([]Gender, error)
	GetGenderByID(ctx context.Context, id string) (*Gender, error)
}

// Gender struct Model
type Gender struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
