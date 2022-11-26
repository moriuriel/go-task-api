package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/moriuriel/go-task-api/adapter/api/handlers"
	"github.com/moriuriel/go-task-api/adapter/presenter"
	"github.com/moriuriel/go-task-api/adapter/repository"
	"github.com/moriuriel/go-task-api/infrastructure/database"
	"github.com/moriuriel/go-task-api/infrastructure/providers"
	"github.com/moriuriel/go-task-api/usecase"
	"go.mongodb.org/mongo-driver/mongo"
)

type Routes struct {
	router *mux.Router
	db     *mongo.Database
}

func NewRoutes() *Routes {
	return &Routes{
		router: NewGorillaMux(),
		db:     database.NewMongodbConnection(),
	}
}

func (r Routes) BuildRoutes() *mux.Router {
	routes := r.router.PathPrefix("/api").Subrouter()
	//Health Routes
	routes.HandleFunc("/v1/health", healthCheck).Methods(http.MethodGet)
	//Task Routes
	routes.HandleFunc("/v1/tasks", r.buildCreateTaskHandler()).Methods(http.MethodPost)
	routes.HandleFunc("/v1/tasks/{owner_id}", r.buildFindTaskByOwnerHandler()).Methods(http.MethodGet)
	//User Routes
	routes.HandleFunc("/v1/users", r.buildCreateUserHandler()).Methods(http.MethodPost)
	return routes
}

func healthCheck(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		Status string `json:"status"`
	}{Status: http.StatusText(http.StatusOK)})
}

func (r Routes) buildCreateTaskHandler() http.HandlerFunc {
	uc := usecase.NewCreateTaskContainer(
		presenter.NewCreateTaskPresenter(),
		repository.NewTaskRepository(r.db),
		repository.NewUserRepository(r.db),
		15*time.Second,
	)

	return handlers.NewCreateAccountHandler(uc).Handle
}

func (r Routes) buildFindTaskByOwnerHandler() http.HandlerFunc {
	uc := usecase.NewTaskByOwnerContainer(
		presenter.NewFindTaskByOnwerPresneter(),
		repository.NewTaskRepository(r.db),
		15*time.Second,
	)

	return handlers.NewFindTaskByOwnerHandler(uc).Handle
}

func (r Routes) buildCreateUserHandler() http.HandlerFunc {
	uc := usecase.NewUserContainer(
		providers.NewHashProvider(),
		presenter.NewCreateUserPresenter(),
		repository.NewUserRepository(r.db),
		15*time.Second,
	)

	return handlers.NewCreateUserHandler(uc).Handle
}
