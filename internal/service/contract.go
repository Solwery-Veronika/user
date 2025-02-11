package service

import (
	"context"
)

type DbRepo interface {
	CreateUser(ctx context.Context, username string) error
}
