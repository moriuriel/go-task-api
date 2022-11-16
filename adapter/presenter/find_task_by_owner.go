package presenter

import (
	"time"

	"github.com/moriuriel/go-task-api/domain"
	"github.com/moriuriel/go-task-api/usecase"
)

type findTaskByOnwerPresenter struct{}

func NewFindTaskByOnwerPresneter() usecase.FindTaskByOwnerPresenter {
	return findTaskByOnwerPresenter{}
}

func (p findTaskByOnwerPresenter) Output(tasks []domain.Task) []usecase.FindTaskByOwnerOutput {
	var o = make([]usecase.FindTaskByOwnerOutput, 0)

	for _, task := range tasks {
		o = append(o, usecase.FindTaskByOwnerOutput{
			Id:          task.ID().String(),
			Title:       task.Title(),
			Description: task.Description(),
			Priority:    task.Priority(),
			Completed:   task.Completed(),
			Owner: domain.TaskOwnerOutput{
				Id:   task.Owner().ID().String(),
				Name: task.Owner().Name(),
			},
			CreatedAt: task.CreatedAt().Format(time.RFC3339),
		})
	}

	return o
}
