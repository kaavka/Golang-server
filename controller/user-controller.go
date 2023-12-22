package controller

import (
	"golang/server/entity"
	"golang/server/service"
	"golang/server/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UserController defines the interface for user-related operations.
type UserController interface {
	FindAll(context *gin.Context)
	Save(context *gin.Context)
	Edit(context *gin.Context)
	Delete(context *gin.Context)
}

type controller struct {
	service service.UsersService
}

// New creates a new UserController with the provided UsersService.
func New(service service.UsersService) UserController {
	return &controller{
		service: service,
	}
}

// FindAll retrieves all users and sends the response to the client.
func (c *controller) FindAll(context *gin.Context) {
	// Get all users
	users := c.service.FindAll()

	// Check if there are no users
	if len(users) == 0 {
		context.JSON(http.StatusOK, []entity.User{})
	}

	// Send the users in the response
	context.JSON(http.StatusOK, users)
}

// Save creates a new user based on the JSON request and sends the response to the client.
func (c *controller) Save(context *gin.Context) {
	var user entity.User

	// Generate a safe ID using UUID
	user.Id = uuid.New().String()

	// Convert user from JSON request to the User type
	convertError := context.BindJSON(&user)

	// Handle conversion or validation errors
	validationError := utils.ValidateUser(user)
	if validationError != nil || convertError != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": convertError.Error()})
		return
	}

	// Save the user
	c.service.Save(user)

	// Send the created user in the response
	context.JSON(http.StatusCreated, user)
}

// Edit updates an existing user based on the JSON request and sends the response to the client.
func (c *controller) Edit(context *gin.Context) {
	var user entity.User

	// Convert user from JSON request to the User type
	convertError := context.BindJSON(&user)

	// Handle conversion error
	if convertError != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": convertError.Error()})
		return
	}

	// Edit the user
	editedUser, err := c.service.Edit(user)

	// Handle not found error
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Send the edited user in the response
	context.JSON(http.StatusOK, editedUser)
}

// Delete removes a user based on the provided ID and sends the response to the client.
func (c *controller) Delete(context *gin.Context) {
	// Retrieve the user ID from the query parameter "id"
	id := context.Query("id")

	// Check if the "id" parameter is missing
	if id == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Missing 'id' parameter"})
		return
	}

	// Delete the user
	err := c.service.Delete(id)

	// Handle not found error
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Send status code
	context.Status(http.StatusNoContent)
}