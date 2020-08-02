package role

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/segmentio/ksuid"
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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_ = r.db.Close(ctx)
}

// AddRole repository implementation
func (r *postgresRepository) AddRole(ctx context.Context, name string, description string) error {
	a := &Role{
		Name:        name,
		Description: description,
	}
	res, err := r.db.Query(ctx, "INSERT INTO roles(id, name, description) VALUES ($1, $2, $3)", ksuid.New().String(), a.Name, a.Description)
	if res != nil {
		res.Close()
	}
	fmt.Println(err)
	return err
}

// GetRoles repository implementation
func (r *postgresRepository) GetRoles(ctx context.Context) ([]Role, error) {
	rows, err := r.db.Query(
		ctx,
		"SELECT id, name, description FROM roles",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []Role
	for rows.Next() {
		a := &Role{}
		if err = rows.Scan(&a.ID, &a.Name, &a.Description); err == nil {
			roles = append(roles, *a)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return roles, nil
}

// GetRoleByID repository implementation
func (r *postgresRepository) GetRoleByID(ctx context.Context, id string) (*Role, error) {
	a := &Role{}
	row := r.db.QueryRow(ctx, `SELECT id, name, description FROM roles WHERE id=$1;`, id)

	err := row.Scan(&a.ID, &a.Name, &a.Description)

	if err != nil {
		return nil, err
	}

	return a, nil

}
