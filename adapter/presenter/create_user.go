package presenter

import (
	"time"

	"github.com/moriuriel/go-task-api/domain"
	"github.com/moriuriel/go-task-api/usecase"
)

type createUserPresenter struct{}

func NewCreateUserPresenter() usecase.CreateUserPresenter {
	return createUserPresenter{}
}

func (p createUserPresenter) Output(user domain.User) usecase.CreateUserOutput {
	return usecase.CreateUserOutput{
		Id:        user.ID().String(),
		Name:      user.Name(),
		Email:     user.Email(),
		Password:  user.Password(),
		CreatedAt: user.CreatedAt().Format(time.RFC3339),
	}
}
