package user

import (
	"context"

	"github.com/soerjadi/apollo/internal/model"
)

func (r *userRepository) GetUserByID(ctx context.Context, id int64) (model.User, error) {
	var (
		res model.User
		err error
	)

	err = r.query.getUserByID.GetContext(ctx, &res, id)
	if err != nil {
		return model.User{}, err
	}

	return res, nil
}
