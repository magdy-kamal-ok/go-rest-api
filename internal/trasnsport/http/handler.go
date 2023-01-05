package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/magdy-kamal-ok/go-rest-api/internal/comment"
)

type Handler struct {
	Router  *mux.Router
	Service *comment.Service
}

func NewHandler(service *comment.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) SetupRoutes() {
	fmt.Println("start setup handler")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "is a live")
	})
}

func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Unable to parse string to uint")
	}
	comment, err := h.Service.GetComment(i)
	if err != nil {
		fmt.Fprintf(w, "Error retreive comment by ID")
	}
	fmt.Fprintf(w, "%+v", comment)
}
