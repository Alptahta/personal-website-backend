package service

import "github.com/Alptahta/personal-website-backend/internal/user"

type UserRepository interface {
	Create(username string, password []byte) (user.User, error)
}

type User struct {
	userRepository UserRepository
}

func NewUser(userRepository UserRepository) *User {
	return &User{userRepository: userRepository}
}

func (u User) Create(username string, password []byte) (user.User, error) {
	return u.userRepository.Create(username, password)
}
