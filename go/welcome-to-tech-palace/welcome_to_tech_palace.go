package techpalace

import (
	"regexp"
	"strings"
)

const (
	welcomeMessage  = "Welcome to the Tech Palace, "
	borderCharacter = "*"
	cleanupRegex    = `[*\n\t]+`
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return welcomeMessage + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	return strings.Repeat(borderCharacter, numStarsPerLine) + "\n" + welcomeMsg + "\n" + strings.Repeat(borderCharacter, numStarsPerLine)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	re := regexp.MustCompile(cleanupRegex)
	return strings.TrimSpace(re.ReplaceAllString(oldMsg, ""))

}
