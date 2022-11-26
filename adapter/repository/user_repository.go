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
	UserRepository struct {
		db         *mongo.Database
		collection string
	}

	UserBSON struct {
		Id        string    `bson:"_id"`
		Name      string    `bson:"name"`
		Email     string    `bson:"email"`
		Password  string    `bson:"password"`
		CreatedAt time.Time `bson:"created_at,"`
	}
)

func NewUserRepository(db *mongo.Database) UserRepository {
	return UserRepository{
		db:         db,
		collection: "users",
	}
}

func (r UserRepository) Create(user domain.User, ctx context.Context) (domain.User, error) {
	var userBson = UserBSON{
		Id:        user.ID().String(),
		Name:      user.Name(),
		Email:     user.Email(),
		Password:  user.Password(),
		CreatedAt: user.CreatedAt(),
	}
	_, err := r.db.Collection(r.collection).InsertOne(ctx, userBson)

	if err != nil {
		return domain.User{}, errors.Wrap(err, domain.ErrCreateUser.Error())
	}

	return user, nil
}

func (r UserRepository) FindById(id string, ctx context.Context) (domain.User, error) {
	var user = &UserBSON{}

	err := r.db.Collection(r.collection).FindOne(ctx, bson.M{"_id": id}).Decode(user)

	if err != nil {
		return domain.User{}, errors.Wrap(err, domain.ErrFindByID.Error())
	}

	return domain.NewUser(
		domain.ID(user.Id),
		user.Name,
		user.Email,
		user.Email,
		user.CreatedAt,
	), nil
}
