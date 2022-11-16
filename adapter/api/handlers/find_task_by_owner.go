package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/moriuriel/go-task-api/adapter/api/response"
	"github.com/moriuriel/go-task-api/usecase"
)

type FindTaskByOwnerHandler struct {
	uc usecase.FindTaskByOwnerUsecase
}

func NewFindTaskByOwnerHandler(uc usecase.FindTaskByOwnerUsecase) FindTaskByOwnerHandler {
	return FindTaskByOwnerHandler{
		uc: uc,
	}
}

func (h FindTaskByOwnerHandler) Handle(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["owner_id"]

	output, err := h.uc.Execute(id, r.Context())
	if err != nil {
		response.NewError(err, http.StatusUnprocessableEntity, id).Send(w)
		return
	}

	response.NewSuccess(output, http.StatusOK).Send(w)
}
