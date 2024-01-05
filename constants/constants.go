package constants

import (
	"golang/server/entity"

	"github.com/google/uuid"
)

// DefaultPort defines the default server port.
var DEFAULT_PORT = "3000"
// UsersRoute defines the route for user-related operations
var USERS_ROUTE = "/users"
var USERS_SAVE_ROUTE = "/users_save"
// Define initial users
var DEFAULT_USERS = []entity.User{
	{
		Id:        uuid.New().String(),
		Name:      "John",
		Surname:   "Doe",
		Age:       30,
		Email:     "john.doe@example.com",
		About:     "Software Engineer",
		IconColor: "blue",
	},
	{
		Id:        uuid.New().String(),
		Name:      "Alice",
		Surname:   "Smith",
		Age:       25,
		Email:     "alice.smith@example.com",
		About:     "Web Developer",
		IconColor: "green",
	},
	{
		Id:        uuid.New().String(),
		Name:      "Bob",
		Surname:   "Johnson",
		Age:       35,
		Email:     "bob.johnson@example.com",
		About:     "Data Scientist",
		IconColor: "red",
	},
	{
		Id:        uuid.New().String(),
		Name:      "Eva",
		Surname:   "Williams",
		Age:       28,
		Email:     "eva.williams@example.com",
		About:     "Graphic Designer",
		IconColor: "purple",
	},
	{
		Id:        uuid.New().String(),
		Name:      "Charlie",
		Surname:   "Brown",
		Age:       40,
		Email:     "charlie.brown@example.com",
		About:     "Project Manager",
		IconColor: "orange",
	},
	{
		Id:        uuid.New().String(),
		Name:      "Grace",
		Surname:   "Miller",
		Age:       33,
		Email:     "grace.miller@example.com",
		About:     "Marketing Specialist",
		IconColor: "pink",
	},
	{
		Id:        uuid.New().String(),
		Name:      "Daniel",
		Surname:   "Clark",
		Age:       28,
		Email:     "daniel.clark@example.com",
		About:     "UI/UX Designer",
		IconColor: "cyan",
	},
	{
		Id:        uuid.New().String(),
		Name:      "Sophia",
		Surname:   "Moore",
		Age:       32,
		Email:     "sophia.moore@example.com",
		About:     "Product Manager",
		IconColor: "yellow",
	},
	{
		Id:        uuid.New().String(),
		Name:      "Michael",
		Surname:   "White",
		Age:       37,
		Email:     "michael.white@example.com",
		About:     "Systems Analyst",
		IconColor: "teal",
	},
}
