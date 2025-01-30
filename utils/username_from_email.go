package utils

import (
	"errors"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type ParamGenerateUsername struct {
	Username     *string
	Email        *string
	Exists       *bool
	LastSequence *uint
}

func GenerateUsername(param *ParamGenerateUsername) (string, error) {
	// If email is not provided, generate a random username
	if param.Email == nil {
		return generateRandomUsername(), nil
	}

	// Validate input email
	if !isValidEmail(*param.Email) {
		return "", errors.New("invalid email provided")
	}

	// Generate base username by cleaning the email (lowercase and removing special characters)
	email := *param.Email
	baseUsername := cleanUsername(email)

	// Check if a username already exists
	if param.Exists != nil && *param.Exists {
		username := baseUsername
		lastSequence := uint(0)
		if param.LastSequence != nil {
			lastSequence = *param.LastSequence
		}

		// Append the next sequence to the username
		lastSequence++
		username = baseUsername + "." + strconv.Itoa(int(lastSequence))
		return username, nil
	}

	// If no existing username and no sequence, return the base username
	return baseUsername, nil
}

// Helper function to validate email
func isValidEmail(email string) bool {
	// Regex for basic email validation
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

// Helper function to clean username by removing special characters
func cleanUsername(email string) string {
	// Extract the part before "@" from the email
	parts := strings.Split(email, "@")
	localPart := parts[0]

	// Remove special characters and lowercase the string
	re := regexp.MustCompile(`[^a-zA-Z0-9]`)
	return re.ReplaceAllString(strings.ToLower(localPart), "")
}

// Helper function to generate a random username
func generateRandomUsername() string {
	const letters = "abcdefghijklmnopqrstuvwxyz"
	const alphanumeric = "abcdefghijklmnopqrstuvwxyz0123456789"

	rand.Seed(time.Now().UnixNano())

	// Generate first 4 alphabet characters
	firstPart := make([]byte, 4)
	for i := 0; i < 4; i++ {
		firstPart[i] = letters[rand.Intn(len(letters))]
	}

	// Generate remaining part (length 4 to 16, total length 8-20)
	remainingLength := rand.Intn(13) + 4 // Min 4, Max 16
	remainingPart := make([]byte, remainingLength)
	for i := 0; i < remainingLength; i++ {
		remainingPart[i] = alphanumeric[rand.Intn(len(alphanumeric))]
	}

	return string(firstPart) + string(remainingPart)
}
