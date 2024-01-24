package user

import "github.com/soerjadi/apollo/internal/repository/user"

func GetUsecase(repo user.Repository) Usecase {
	return &userUsecase{
		repository: repo,
	}
}
