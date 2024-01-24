package user

import (
	"context"

	"github.com/soerjadi/apollo/internal/model"
)

func (u *userUsecase) GetUserByID(ctx context.Context, id int64) (model.User, error) {
	return model.User{}, nil
}
