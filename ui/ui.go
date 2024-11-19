package ui

import (
	"fmt"

	internaldefense "chandanbsd.com/fileCRUD/internaldefense"
)

const CLOSE = 5
const INVALID_CHOICE = -1

func displayMenu() int {
	var choice int
	fmt.Printf(`
		       Menu
		--------------------
		1. Read a file
		2. Create a file
		3. Write to a file
		4. Delete a file
		5. Close

		Enter your choice: `)

	fmt.Scan(&choice)
	return choice
}

func validateMenuSelection(choice int) bool {
	if choice < 1 || choice > 5 {
		return false
	} else {
		return true
	}
}

func determineUserSuspension(failedAttempts int, email string) bool {
	if failedAttempts >= internaldefense.MAX_FAILED_ATTEMPTS {
		internaldefense.LogBadActors(email)
		return true
	}
	return false
}

func RenderUI(email string, failedAttempts *int) int {
	var choice int
	var isSelectionValid bool

	choice = displayMenu()
	isSelectionValid = validateMenuSelection(choice)

	if !isSelectionValid {
		*failedAttempts++
		isUserSuspended := determineUserSuspension(*failedAttempts, email)
		if isUserSuspended {
			return internaldefense.USER_SUSPENDED
		}

		return INVALID_CHOICE

	} else if choice == 5 {
		return choice
	} else {
		return choice
	}
}
