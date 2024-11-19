package main

import (
	"fmt"
	"os"
	"path/filepath"

	filestore "chandanbsd.com/fileCRUD/filestore"
	internaldefense "chandanbsd.com/fileCRUD/internaldefense"
	ui "chandanbsd.com/fileCRUD/ui"
)

func main() {
	fmt.Print("Welcome to fileCRUD an open source interactive tool to CREATE, READ, UPDATE and DELETE files:\n\n")

	var email string
	var baseDirErr error
	var baseDir string = ""
	var failedAttempts int = 0

	baseDirForLogErr := filestore.CreateLogsDirectoryAndFileIfItDoesNotExist()
	if baseDirForLogErr != nil {
		fmt.Print("Failed to create directory for logs")
		os.Exit(1)
	}

	email = internaldefense.CaptureEmail()

	for {

		if baseDir == "" {
			baseDir, baseDirErr = filestore.CreateOrCaptureBaseDir()
		}

		if baseDirErr != nil {
			print("Failed to find or create directory")
			baseDir = ""
			continue
		}

		var choice int = ui.RenderUI(email, &failedAttempts)

		if choice == ui.CLOSE {
			break
		} else if choice == internaldefense.USER_SUSPENDED {
			filestore.WriteToFile(filepath.Join(filestore.LOG_DIR, internaldefense.SUSPENDED_USERS_FILE), email)
		} else if choice == ui.INVALID_CHOICE {
			continue
		} else {
			err := filestore.ProcessFileOperation(baseDir, choice)

			fmt.Print("\n", err)
		}
	}
}
