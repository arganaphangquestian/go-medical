package history

import "context"

// Service for History Module
type Service interface {
	AddHistory(ctx context.Context, userID string, diseaseID string, note string) error
	GetHistories(ctx context.Context) ([]History, error)
	GetHistoryByID(ctx context.Context, id string) (*History, error)
}

// History struct Model
type History struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	DiseaseID string `json:"disease_id"`
	Note      string `json:"note"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
