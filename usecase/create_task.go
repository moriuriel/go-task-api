package usecase

import (
	"context"
	"time"

	"github.com/moriuriel/go-task-api/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	CreateTaskInput struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Priority    string `json:"priority"`
		OwnerID     string `json:"owner_id"`
	}

	CreateTaskOutput struct {
		Id          string                 `json:"_id"`
		Title       string                 `json:"title"`
		Description string                 `json:"description"`
		Priority    string                 `json:"priority"`
		Completed   bool                   `json:"completed"`
		Owner       domain.TaskOwnerOutput `json:"owner"`
		CreatedAt   string                 `json:"created_at"`
	}

	CreateTaskUsecase interface {
		Execute(CreateTaskInput, context.Context) (CreateTaskOutput, error)
	}

	CreateTaskPresenter interface {
		Output(domain.Task) CreateTaskOutput
	}

	CreateTaskContainer struct {
		pre        CreateTaskPresenter
		repo       domain.TaskRepository
		userRepo   domain.UserRepository
		ctxTimeout time.Duration
	}
)

func NewCreateTaskContainer(p CreateTaskPresenter, r domain.TaskRepository, ur domain.UserRepository, t time.Duration) CreateTaskContainer {
	return CreateTaskContainer{
		pre:        p,
		repo:       r,
		userRepo:   ur,
		ctxTimeout: t,
	}
}

func (uc CreateTaskContainer) Execute(input CreateTaskInput, ctx context.Context) (CreateTaskOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.ctxTimeout)
	defer cancel()

	taskID := primitive.NewObjectID().Hex()

	user, err := uc.userRepo.FindById(input.OwnerID, ctx)
	if err != nil {
		return uc.pre.Output(domain.Task{}), err
	}

	owner := domain.NewOwner(
		domain.ID(input.OwnerID),
		user.Name(),
	)

	var task = domain.NewTask(
		domain.ID(taskID),
		input.Title,
		input.Description,
		input.Priority,
		false,
		time.Now(),
		owner,
	)

	task, err = uc.repo.Create(task, ctx)
	if err != nil {
		return uc.pre.Output(domain.Task{}), err
	}

	return uc.pre.Output(task), nil

}
