package utils

import "github.com/aditya13raja/alumni-student-backend/models"

// FormatUserResponse removes sensitive data like password
func FormatUserResponse(user models.User) map[string]interface{} {
	return map[string]interface{}{
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"username":   user.Username,
		"role":       user.Role,
	}
}
