package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/magdy-kamal-ok/go-rest-api/internal/comment"
)

type Response struct {
	Message string
	Error   string
}

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
	h.Router.HandleFunc("/api/comment/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/comment", h.PostComment).Methods("POST")
	h.Router.HandleFunc("/api/comment", h.GetAllComments).Methods("GET")
	h.Router.HandleFunc("/api/comment/{id}", h.DeleteComment).Methods("DELETE")
	h.Router.HandleFunc("/api/comment/{id}", h.UpdateComment).Methods("PUT")

	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(Response{Message: "Hello i am alive"}); err != nil {
			panic(err)
		}

	})
}

func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Unable to parse string to uint", err)
		return
	}
	comment, err := h.Service.GetComment(uint(i))
	if err != nil {
		sendErrorResponse(w, "Error retreive comment by ID", err)
		return
	}
	if err := sendOkResponse(w, comment); err != nil {
		panic(err)
	}
}

func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	comments, err := h.Service.GetAllComments()
	if err != nil {
		sendErrorResponse(w, "Error retreive comments", err)
		return
	}
	if err := sendOkResponse(w, comments); err != nil {
		panic(err)
	}
}

func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	var comment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		sendErrorResponse(w, "Failed to decode comment json body", err)
		return
	}
	comment, err := h.Service.PostComment(comment)
	if err != nil {
		sendErrorResponse(w, "Error save comment", err)
		return
	}
	if err := sendOkResponse(w, comment); err != nil {
		panic(err)
	}
}

func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	var comment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		sendErrorResponse(w, "Failed to decode comment json body", err)
		return
	}
	vars := mux.Vars(r)
	id := vars["id"]
	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Unable to parse string to uint", err)
		return
	}
	comment, err = h.Service.UpdateComment(uint(i), comment)
	if err != nil {
		sendErrorResponse(w, "Error update comment", err)
		return
	}
	if err = sendOkResponse(w, comment); err != nil {
		panic(err)
	}
}

func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Unable to parse string to uint")
	}
	err = h.Service.DeleteComment(uint(i))
	if err != nil {
		sendErrorResponse(w, "Error retreive comment by ID", err)
	}
	if err := sendOkResponse(w, Response{Message: "Successfully deleted"}); err != nil {
		panic(err)
	}
}

func sendOkResponse(w http.ResponseWriter, resp interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(resp)
}

func sendErrorResponse(w http.ResponseWriter, message string, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(Response{Message: message, Error: err.Error()}); err != nil {
		panic(err)
	}
}
