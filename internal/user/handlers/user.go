package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Alptahta/personal-website-backend/internal/user"
)

type UserService interface {
	Create(username string, password []byte) (user.User, error)
}

type User struct {
	l           *log.Logger
	userService UserService
}

func NewUser(l *log.Logger, us UserService) *User {
	return &User{l: l, userService: us}
}

func (r User) Register(mux *http.ServeMux) {
	mux.HandleFunc("/register", r.SignUp)
}

// CreateUsersRequest defines the request used for creating users.
type CreateUsersRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u User) SignUp(rw http.ResponseWriter, r *http.Request) {
	var req CreateUsersRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return
	}
	if r.Method == http.MethodPost {
		user, _ := u.userService.Create(req.Username, []byte(req.Password))
		content, _ := json.Marshal(user)
		rw.Write(content)

	}
}
