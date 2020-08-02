package history

import "context"

type historyService struct {
	repository Repository
}

// NewService methods
func NewService(r Repository) Service {
	return &historyService{r}
}

func (h historyService) AddHistory(ctx context.Context, userID string, diseaseID string, note string) error {
	return h.repository.AddHistory(ctx, userID, diseaseID, note)
}

func (h historyService) GetHistories(ctx context.Context) ([]History, error) {
	return h.repository.GetHistories(ctx)
}

func (h historyService) GetHistoryByID(ctx context.Context, id string) (*History, error) {
	return h.repository.GetHistoryByID(ctx, id)
}
