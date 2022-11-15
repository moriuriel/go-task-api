package routes

import "github.com/gorilla/mux"

func NewGorillaMux() *mux.Router {
	return mux.NewRouter()
}
