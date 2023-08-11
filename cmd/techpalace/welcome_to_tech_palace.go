package techpalace

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	// => Welcome to the Tech Palace, JUDY
	return fmt.Sprintf("Welcome to the Tech Palace, %v", strings.ToUpper(customer))
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	/**
	**********
	Welcome!
	**********
	*/
	sb := strings.Builder{}
	repeat := strings.Repeat("*", numStarsPerLine)
	sb.WriteString(repeat)
	sb.WriteByte('\n')
	sb.WriteString(welcomeMsg)
	sb.WriteByte('\n')
	sb.WriteString(repeat)
	return sb.String()
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	/**
	message := `
	**************************
	*    BUY NOW, SAVE 10%   *
	**************************
	`

	*/
	rep := strings.ReplaceAll(oldMsg, "*", "")
	space := strings.TrimSpace(rep)
	return space
}
