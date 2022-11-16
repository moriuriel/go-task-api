package repository

import (
	"context"
	"time"

	"github.com/moriuriel/go-task-api/domain"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	TaskRepository struct {
		db         *mongo.Database
		collection string
	}

	TaskBSON struct {
		Id          string    `bson:"_id"`
		Title       string    `bson:"title"`
		Description string    `bson:"description"`
		Priority    string    `bson:"priority"`
		Completed   bool      `bson:"completed"`
		Owner       OwnerBSON `bson:"owner"`
		CompletedAt time.Time `bson:"completed_at,omitempty"`
		CreatedAt   time.Time `bson:"created_at"`
	}

	OwnerBSON struct {
		Id   string `bson:"_id"`
		Name string `bson:"name"`
	}
)

func NewTaskRepository(db *mongo.Database) TaskRepository {
	return TaskRepository{
		db:         db,
		collection: "tasks",
	}
}

func (r TaskRepository) Create(task domain.Task, ctx context.Context) (domain.Task, error) {
	var taskBson = TaskBSON{
		Id:          task.ID().String(),
		Title:       task.Title(),
		Description: task.Description(),
		Priority:    task.Priority(),
		Completed:   task.Completed(),
		Owner: OwnerBSON{
			Id:   task.Owner().ID().String(),
			Name: task.Owner().Name(),
		},
		CreatedAt: task.CreatedAt(),
	}
	_, err := r.db.Collection(r.collection).InsertOne(ctx, taskBson)

	if err != nil {
		return domain.Task{}, errors.Wrap(err, domain.ErrCreateTask.Error())
	}

	return task, nil
}

func (r TaskRepository) FindByOwner(id string, ctx context.Context) ([]domain.Task, error) {
	tasksBson := make([]TaskBSON, 0)

	cur, err := r.db.Collection(r.collection).Find(ctx, bson.M{
		"owner._id": id,
	})
	if err != nil {
		return []domain.Task{}, errors.Wrap(err, domain.ErrFindTaskByOnwer.Error())
	}
	defer cur.Close(ctx)

	if err = cur.All(ctx, &tasksBson); err != nil {
		return []domain.Task{}, errors.Wrap(err, domain.ErrFindTaskByOnwer.Error())
	}

	var tasks = make([]domain.Task, 0)

	for _, taskBson := range tasksBson {
		task := domain.NewTask(
			domain.ID(taskBson.Id),
			taskBson.Title,
			taskBson.Description,
			taskBson.Priority,
			taskBson.Completed,
			taskBson.CompletedAt,
			domain.NewOwner(
				domain.ID(taskBson.Owner.Id),
				taskBson.Owner.Name,
			),
		)

		tasks = append(tasks, task)
	}

	return tasks, nil

}
