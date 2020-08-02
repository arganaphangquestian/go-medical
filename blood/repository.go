package blood

import "context"

// Repository Blood
type Repository interface {
	Close()

	AddBlood(ctx context.Context, name string, rhesus bool, description string) error
	GetBloods(ctx context.Context) ([]Blood, error)
	GetBloodByID(ctx context.Context, id string) (*Blood, error)
}
