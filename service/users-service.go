package service

import (
	"fmt"
	"golang/server/constants"
	"golang/server/entity"
)

// UsersService defines the interface for user-related operations.
type UsersService interface {
	FindAll() []entity.User
	Save(entity.User) entity.User
	Edit(entity.User) (entity.User, error)
	Delete(string) (error)
}

type userService struct {
	users []entity.User
}

// New creates a new UsersService with a default set of users.
func New() UsersService {
	return &userService{users: constants.DEFAULT_USERS}
}

// FindAll returns all users.
func (service *userService) FindAll() []entity.User {
	// Return all users
	return service.users
}

// Save adds a new user to the list of users.
func (service *userService) Save(user entity.User) entity.User {
	// Add user to the list
	service.users = append(service.users, user)

	return user
}

// Edit updates an existing user in the list.
func (service *userService) Edit(editedUser entity.User) (entity.User, error) {
	for index, user := range service.users {
		if user.Id == editedUser.Id {
			// Update the user in the slice
			service.users[index] = editedUser
			return editedUser, nil
		}
	}

	// If no user found
	return entity.User{}, fmt.Errorf("user with ID %v not found", editedUser.Id)
}

// Delete removes a user from the list based on the provided ID.
func (service *userService) Delete(id string) (error) {
	for index, user := range service.users {
		if user.Id == id {
			// Delete the user from the slice using slicing
			service.users = append(service.users[:index], service.users[index+1:]...)
			return nil
		}
	}

	// If no user found
	return fmt.Errorf("user with ID %v not found", id)
}