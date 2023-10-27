package user

import (
	"context"
	"restapi-lesson/pkg/logging"
)

type Service struct {
	storage Storage
	logger  *logging.Logger
}

func (s *Service) Create(ctx context.Context, dto CreateUserDTO) (user User, err error) {
	//TODO next
	panic("implement me")
}
