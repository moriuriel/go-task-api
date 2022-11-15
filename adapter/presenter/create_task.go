package presenter

import (
	"time"

	"github.com/moriuriel/go-task-api/domain"
	"github.com/moriuriel/go-task-api/usecase"
)

type createTaskPresenter struct{}

func NewCreateTaskPresenter() usecase.CreateTaskPresenter {
	return createTaskPresenter{}
}

func (p createTaskPresenter) Output(task domain.Task) usecase.CreateTaskOutput {
	return usecase.CreateTaskOutput{
		Id:          task.ID().String(),
		Title:       task.Title(),
		Description: task.Description(),
		Priority:    task.Priority(),
		Completed:   task.Completed(),
		Owner: usecase.TaskOwnerOutput{
			Id:   task.Owner().ID().String(),
			Name: task.Owner().Name(),
		},
		CreatedAt: task.CreatedAt().Format(time.RFC3339),
	}
}
