package repositories

import (
	"fmt"

	"github.com/jn0x/reddigo/domain"
	"github.com/jn0x/reddigo/storage"
	uuid "github.com/satori/go.uuid"
)

type UserRepository interface {
	CreateUser(user *domain.User) error
	DeleteUser(id uuid.UUID) error
	UpdateUser(user *domain.User) error
	GetUser(id uuid.UUID) domain.User
	GetUserByUsername(username string) domain.User
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return UserRepository(&userRepository{})
}

func (r *userRepository) CreateUser(user *domain.User) error {
	result := storage.DB.Create(user)
	if result.Error != nil {
		fmt.Println(result.Error)
		return fmt.Errorf("user with this username already exists")
	}
	return nil
}

func (r *userRepository) UpdateUser(user *domain.User) error {
	return nil
}

func (r *userRepository) GetUser(uuid.UUID) domain.User {
	return domain.User{}
}

func (r *userRepository) DeleteUser(id uuid.UUID) error {
	return nil
}

func (r *userRepository) GetUserByUsername(username string) domain.User {
	var user domain.User
	storage.DB.First(&user, "username = ?", username)
	return user
}
