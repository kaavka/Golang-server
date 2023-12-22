package utils

import (
	"fmt"
	"golang/server/entity"
)

func ValidateUser(user entity.User) error {
	if user.Name == "" {
		return fmt.Errorf("Name is required")
	}

	if user.Surname == "" {
		return fmt.Errorf("Surname is required")
	}

	if user.Age < 0 {
		return fmt.Errorf("Age must be a non-negative value")
	}

	if user.Email == "" {
		return fmt.Errorf("Email is required")
	}

	if user.IconColor == "" {
		return fmt.Errorf("IconColor is required")
	}

	return nil
}