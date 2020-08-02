package history

import (
	"context"
	"github.com/jackc/pgx/v4"
	"time"
)

type postgresRepository struct {
	db *pgx.Conn
}

// NewPostgres methods
func NewPostgres(url string) (Repository, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db, err := pgx.Connect(context.Background(), url)
	if err != nil {
		return nil, err
	}
	err = db.Ping(ctx)
	if err != nil {
		return nil, err
	}
	return &postgresRepository{db}, nil
}

func (r *postgresRepository) Close() {
	_ = r.db.Close(context.Background())
}

func (r *postgresRepository) AddHistory(ctx context.Context, userID string, diseaseID string, note string) error {
	a := &History{
		UserID:    userID,
		DiseaseID: diseaseID,
		Note:      note,
	}
	_, err := r.db.Query(ctx, "INSERT INTO histories(user_id, disease_id, note) VALUES($1, $2, $3)", a.UserID, a.DiseaseID, a.Note)
	return err
}

func (r *postgresRepository) GetHistories(ctx context.Context) ([]History, error) {
	rows, err := r.db.Query(
		ctx,
		"SELECT id, user_id, disease_id, note FROM histories",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var histories []History
	for rows.Next() {
		a := &History{}
		if err = rows.Scan(&a.ID, &a.UserID, &a.DiseaseID, &a.Note); err == nil {
			histories = append(histories, *a)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return histories, nil
}

func (r *postgresRepository) GetHistoryByID(ctx context.Context, id string) (*History, error) {
	a := &History{}
	row := r.db.QueryRow(ctx, `SELECT id, user_id, disease_id, note FROM histories WHERE id=$1;`, id)

	err := row.Scan(&a.ID, &a.UserID, &a.DiseaseID, &a.Note)

	if err != nil {
		return nil, err
	}

	return a, nil
}
