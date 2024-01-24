package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/soerjadi/apollo/internal/delivery/rest"
	"github.com/soerjadi/apollo/internal/usecase/user"
)

func NewHandler(usecase user.Usecase) rest.API {
	return &Handler{
		usecase: usecase,
	}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/test", rest.HandlerFunc(h.HelloWorld).Serve).Methods(http.MethodGet)
}
