package user

import (
	"context"

	"github.com/soerjadi/apollo/internal/model"
)

//go:generate mockgen -package=mocks -mock_names=Repository=MockUserRepository -destination=../../mocks/user_repository_mock.go -source=type.go
type Repository interface {
	GetUserByID(ctx context.Context, id int64) (model.User, error)
}

type userRepository struct {
	query prepareQuery
}
