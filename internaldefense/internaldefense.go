package internaldefense

import (
	"fmt"
	"os"
	"strings"

	filestore "chandanbsd.com/fileCRUD/filestore"
)

const SUSPENDED_USERS_FILE = "suspended_users.txt"
const MAX_FAILED_ATTEMPTS = 3
const USER_SUSPENDED = 941

func CaptureEmail() string {
	var email string
	fmt.Print("Please enter your email to continue: ")
	fmt.Scan(&email)

	if strings.Trim(email, " ") == "" || !strings.Contains(email, "@") {
		fmt.Print("\nCannot proceed with invalid email\n")
		os.Exit(1)
	}

	return email
}

func LogBadActors(email string) {
	filestore.WriteToFile(SUSPENDED_USERS_FILE, email)
}
