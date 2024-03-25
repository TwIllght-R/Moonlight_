package core

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func HashedPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password %v", err)
	}
	return string(hashPassword), nil
}

func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(password), []byte(hashedPassword))
}

func ValidatePriority(priority string) error {
	validPriorities := []string{"low", "medium", "high", "critical"}
	isValidPriority := false
	for _, p := range validPriorities {
		if priority == p {
			isValidPriority = true
			break
		}
	}
	if !isValidPriority {
		return fmt.Errorf("invalid priority. Priority should be one of: low, medium, high, critical")
	}
	return nil
}

func ValidateAssignedTo(assignedTo string) error {
	if assignedTo == "" {
		return fmt.Errorf("assignedTo is required")
	}
	return nil
}
func ConvertObjectIDToString(id primitive.ObjectID) string {
	return id.Hex()
}

func ConvertStringToObjectID(id string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(id)
}

// func isValidPriority(priority string) bool {
//     switch priority {
//     case "low", "medium", "high", "critical":
//         return true
//     default:
//         return false
//     }
// }
