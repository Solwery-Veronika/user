package service

import (
	"context"

	"github.com/Solwery-Veronika/user/pkg/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	user.UnimplementedUserServiceServer
	dbR DbRepo
}

func New(repo DbRepo) *Service {
	return &Service{
		dbR: repo,
	}
}

func (s *Service) CreateUser(ctx context.Context, in *user.CreateUserIn) (*user.CreateUserOut, error) {
	err := s.dbR.CreateUser(ctx, in.Username)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &user.CreateUserOut{
		Success: true,
	}, nil
}
