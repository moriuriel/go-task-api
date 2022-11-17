package domain

import (
	"context"
	"errors"
	"time"
)

var (
	ErrCreateTask      = errors.New("error creating task")
	ErrFindTaskByOnwer = errors.New("error to find task by onwer")
)

type (
	ID   string
	Task struct {
		id          ID
		title       string
		description string
		priority    string
		completed   bool
		owner       *Owner
		completedAt time.Time
		createdAt   time.Time
	}
	TaskRepository interface {
		Create(Task, context.Context) (Task, error)
		FindByOwner(string, context.Context) ([]Task, error)
	}
)

func NewTask(id ID, title string, description string, priority string, completed bool, createdAt time.Time, owner *Owner) Task {
	return Task{
		id:          id,
		title:       title,
		description: description,
		priority:    priority,
		completed:   completed,
		owner:       owner,
		createdAt:   createdAt,
	}
}

func UpdateCompleteDate(completedAt time.Time) Task {
	return Task{
		completedAt: completedAt,
	}
}

func (Id ID) String() string {
	return string(Id)
}

func (t Task) ID() ID {
	return t.id
}

func (t Task) Title() string {
	return t.title
}

func (t Task) Description() string {
	return t.description
}

func (t Task) Priority() string {
	return t.priority
}

func (t Task) Completed() bool {
	return t.completed
}

func (t Task) CreatedAt() time.Time {
	return t.createdAt
}

func (t Task) Owner() *Owner {
	return t.owner
}

func (t Task) CompletedAt() time.Time {
	return t.completedAt
}
