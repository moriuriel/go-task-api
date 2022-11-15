package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/moriuriel/go-task-api/adapter/api/response"
	"github.com/moriuriel/go-task-api/usecase"
)

type CreateTaskHandler struct {
	uc usecase.CreateTaskUsecase
}

func NewCreateAccountHandler(uc usecase.CreateTaskUsecase) CreateTaskHandler {
	return CreateTaskHandler{
		uc: uc,
	}
}

func (h CreateTaskHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateTaskInput

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
