package gender

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

// AddGender repository implementation
func (r *postgresRepository) AddGender(ctx context.Context, name string, description string) error {
	a := &Gender{
		Name:        name,
		Description: description,
	}
	_, err := r.db.Query(ctx, "INSERT INTO genders(name, description) VALUES($1, $2)", a.Name, a.Description)
	return err
}

// GetGenders repository implementation
func (r *postgresRepository) GetGenders(ctx context.Context) ([]Gender, error) {
	rows, err := r.db.Query(
		ctx,
		"SELECT id, name, description FROM genders",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var genders []Gender
	for rows.Next() {
		a := &Gender{}
		if err = rows.Scan(&a.ID, &a.Name, &a.Description); err == nil {
			genders = append(genders, *a)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return genders, nil
}

// GetGenderByID repository implementation
func (r *postgresRepository) GetGenderByID(ctx context.Context, id string) (*Gender, error) {
	a := &Gender{}
	row := r.db.QueryRow(ctx, `SELECT id, name, description FROM genders WHERE id=$1;`, id)

	err := row.Scan(&a.ID, &a.Name, &a.Description)

	if err != nil {
		return nil, err
	}

	return a, nil

}
