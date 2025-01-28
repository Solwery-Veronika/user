package service

import (
	"context"

	"github.com/Solwery-Veronika/user/pkg/user"
)

type Service struct {
	user.UnimplementedUserServiceServer
}

func New() *Service {
	return &Service{}
}

func (s *Service) CreateUser(context.Context, *user.CreateUserIn) (*user.CreateUserOut, error) {
	return &user.CreateUserOut{
		Success: true,
	}, nil
}
