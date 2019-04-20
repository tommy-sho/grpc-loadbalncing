package server

import "fmt"

type BackendRepository interface {
	GetMessageByName(name string) (string, error)
}

func NewUserRepository() BackendRepository {
	return &userRepository{}
}

type userRepository struct{}

func (u *userRepository) GetMessageByName(name string) (string, error) {
	switch name != "" {
	case true:
		return fmt.Sprintf("Hello, %v, Nice to meet you!", name), nil
	default:
		return "", fmt.Errorf("no name error")
	}
}
