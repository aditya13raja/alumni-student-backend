package utils

import "github.com/aditya13raja/alumni-student-backend/models"

// FormatUserResponse removes sensitive data like password
func FormatUserResponse(user models.User) map[string]interface{} {
	return map[string]interface{}{
		"first_name":   user.FirstName,
		"last_name":    user.LastName,
		"age":          user.Age,
		"role":         user.Role,
		"degree":       user.Degree,
		"major":        user.Major,
		"passing_year": user.PassingYear,
		"username":     user.Username,
		"email":        user.Email,
	}
}
