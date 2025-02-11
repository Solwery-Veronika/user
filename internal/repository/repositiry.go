package repository

import (
	"context"
	"fmt"
)

type Repository struct {
	dataMap map[string]string
}

func New() *Repository {
	return &Repository{
		dataMap: make(map[string]string),
	}
}

func (r *Repository) CreateUser(ctx context.Context, username string) error {
	_, ok := r.dataMap[username] // проверяем есть ли имя в мапе
	if ok {
		return fmt.Errorf("username %s already exists")
	}
	return nil
}
