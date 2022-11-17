package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/moriuriel/go-task-api/adapter/api/response"
	"github.com/moriuriel/go-task-api/usecase"
)

type CreateUserHandler struct {
	uc usecase.CreateUserUsecase
}

func NewCreateUserHandler(uc usecase.CreateUserUsecase) CreateUserHandler {
	return CreateUserHandler{
		uc: uc,
	}
}

func (h CreateUserHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateUserInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		response.NewError(err, http.StatusBadRequest, input).Send(w)
		return
	}
	defer r.Body.Close()

	output, err := h.uc.Execute(input, r.Context())
	if err != nil {
		response.NewError(err, http.StatusUnprocessableEntity, input).Send(w)
		return
	}

	response.NewSuccess(output, http.StatusCreated).Send(w)
}
