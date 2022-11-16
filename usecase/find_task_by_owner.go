package usecase

import (
	"context"
	"time"

	"github.com/moriuriel/go-task-api/domain"
)

type (
	FindTaskByOwnerInput struct {
		ID string
	}

	FindTaskByOwnerOutput struct {
		Id          string                 `json:"_id"`
		Title       string                 `json:"title"`
		Description string                 `json:"description"`
		Priority    string                 `json:"priority"`
		Completed   bool                   `json:"completed"`
		Owner       domain.TaskOwnerOutput `json:"owner"`
		CreatedAt   string                 `json:"created_at"`
	}

	FindTaskByOwnerPresenter interface {
		Output([]domain.Task) []FindTaskByOwnerOutput
	}

	FindTaskByOwnerUsecase interface {
		Execute(string, context.Context) ([]FindTaskByOwnerOutput, error)
	}

	FindTaskByOwnerContainer struct {
		pre        FindTaskByOwnerPresenter
		repo       domain.TaskRepository
		ctxTimeout time.Duration
	}
)

func NewTaskByOwnerContainer(p FindTaskByOwnerPresenter, r domain.TaskRepository, t time.Duration) FindTaskByOwnerContainer {
	return FindTaskByOwnerContainer{
		pre:        p,
		repo:       r,
		ctxTimeout: t,
	}
}

func (uc FindTaskByOwnerContainer) Execute(id string, ctx context.Context) ([]FindTaskByOwnerOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.ctxTimeout)
	defer cancel()

	tasks, err := uc.repo.FindByOwner(id, ctx)
	if err != nil {
		return uc.pre.Output([]domain.Task{}), err
	}

	return uc.pre.Output(tasks), nil
}
