package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/Solwery-Veronika/user/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var ErrUserExists = errors.New("user already exists")

type Repository struct {
	conn *sqlx.DB
}

func NewRepository(cfg *config.Config) *Repository {
	connectCmd := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.Database, cfg.Postgres.Host, cfg.Postgres.Port)

	conn, err := sqlx.Connect("postgres", connectCmd)
	if err != nil {
		log.Fatal(err)
	}
	return &Repository{conn: conn}
}

func (r *Repository) CreateUser(ctx context.Context, username string) error {
	query := `SELECT true FROM participants WHERE username = $1`

	var exists bool

	err := r.conn.GetContext(ctx, &exists, query, username)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("failed to get user: %w", err)
	}

	if exists {
		return ErrUserExists
	}

	query = `INSERT INTO participants (username) 
	          VALUES ($1)` // запрос

	_, err = r.conn.ExecContext(ctx, query, username)
	if err != nil {
		return fmt.Errorf("failed to insert new user: %w", err)
	}
	return nil
}
