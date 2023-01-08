package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/magdy-kamal-ok/go-rest-api/internal/comment"
	log "github.com/sirupsen/logrus"
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

type FuncHttpDefinition = func(w http.ResponseWriter, r *http.Request)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(
			log.Fields{
				"Method": r.Method,
				"Path":   r.URL.Path,
			}).Info("Endpoint hit")
		next.ServeHTTP(w, r)
	})
}

func BasicAuth(original FuncHttpDefinition) FuncHttpDefinition {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("Basic Auth hit endpoint")
		user, pass, ok := r.BasicAuth()
		if user == "admin" && pass == "password" && ok {
			original(w, r)
		} else {
			sendErrorResponse(w, "Not Authorized", errors.New("Not Authorized user"))
		}
	}
}

func validateToken(accessToken string) bool {
	var mySigningKey = []byte("missionimpossible")
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there has been an error")
		}
		return mySigningKey, nil
	})
	if err != nil {
		return false
	}
	return token.Valid
}

func JWTAuth(original FuncHttpDefinition) FuncHttpDefinition {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("jwt Auth hit endpoint")
		authHeader := r.Header["Authorization"]
		if authHeader == nil {
			sendErrorResponse(w, "Not Jwt Authorized", errors.New("Not Jwt Authorized user"))
			return
		}
		authHeaderParts := strings.Split(authHeader[0], " ")
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			sendErrorResponse(w, "Not Jwt Authorized", errors.New("Not Jwt Authorized user"))
			return
		}
		if validateToken(authHeaderParts[1]) {
			original(w, r)
		} else {
			sendErrorResponse(w, "Not Jwt Authorized", errors.New("Not Jwt Authorized user"))
		}
	}
}

func (h *Handler) SetupRoutes() {
	log.Info("start setup handler")
	h.Router = mux.NewRouter()
	h.Router.Use(LoggingMiddleware)
	h.Router.HandleFunc("/api/comment/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/comment", JWTAuth(h.PostComment)).Methods("POST")
	h.Router.HandleFunc("/api/comment", h.GetAllComments).Methods("GET")
	h.Router.HandleFunc("/api/comment/{id}", BasicAuth(h.DeleteComment)).Methods("DELETE")
	h.Router.HandleFunc("/api/comment/{id}", BasicAuth(h.UpdateComment)).Methods("PUT")

	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(Response{Message: "Hello i am alive"}); err != nil {
			panic(err)
		}

	})
}
