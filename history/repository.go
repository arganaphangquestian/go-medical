package history

import "context"

type Repository interface {
	Close()

	AddHistory(ctx context.Context, userID string, diseaseID string, note string) error
	GetHistories(ctx context.Context) ([]History, error)
	GetHistoryByID(ctx context.Context, id string) (*History, error)
}
