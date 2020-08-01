package gender

import "context"

// Repository Gender
type Repository interface {
	Close()

	AddGender(ctx context.Context, name string, description string) error
	GetGenders(ctx context.Context) ([]Gender, error)
	GetGenderByID(ctx context.Context, id string) (*Gender, error)
}
