package usecase

import (
	"context"
	"time"

	"github.com/moriuriel/go-task-api/domain"
	hashprovider "github.com/moriuriel/go-task-api/infrastructure/providers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	CreateUserInput struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	CreateUserOutput struct {
		Id        string `json:"_id"`
		Name      string `json:"name"`
		Email     string `json:"email"`
		Password  string `json:"password"`
		CreatedAt string `json:"created_at"`
	}

	CreateUserUsecase interface {
		Execute(CreateUserInput, context.Context) (CreateUserOutput, error)
	}

	CreateUserPresenter interface {
		Output(domain.User) CreateUserOutput
	}

	CreateUserContainer struct {
		hashProvider hashprovider.HashProvider
		pre          CreateUserPresenter
		repo         domain.UserRepository
		ctxTimeout   time.Duration
	}
)

func NewUserContainer(hashProvider hashprovider.HashProvider, p CreateUserPresenter, r domain.UserRepository, t time.Duration) CreateUserContainer {
	return CreateUserContainer{
		hashProvider: hashProvider,
		pre:          p,
		repo:         r,
		ctxTimeout:   t,
	}
}

func (uc CreateUserContainer) Execute(input CreateUserInput, ctx context.Context) (CreateUserOutput, error) {
	hashedPassword := uc.hashProvider.GenerateHash(input.Password)

	var user = domain.NewUser(
		domain.ID(primitive.NewObjectID().Hex()),
		input.Name,
		input.Email,
		hashedPassword,
		time.Now(),
	)

	user, err := uc.repo.Create(user, ctx)
	if err != nil {
		return uc.pre.Output(domain.User{}), err
	}

	return uc.pre.Output(user), nil
}
