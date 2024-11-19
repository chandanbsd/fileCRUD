package filestore

import (
	"fmt"
	"os"
	"path/filepath"
)

const READ_FILE = 1
const CREATE_FILE = 2
const WRITE_FILE = 3
const DELETE_FILE = 4
const LOG_DIR = "logs"
const SUSPENDED_USERS_FILE = "suspended_users.txt"

func createFile(filename string) error {
	file, err := os.Create(filename)

	if err != nil {
		return err
	}

	err = file.Close()

	if err != nil {
		return err
	}

	return err
}

func readFromFile(filename string) error {
	fileContent, err := os.ReadFile(filename)

	if err != nil {
		return err
	}

	_, err = fmt.Printf("\nContents of %s:\n %s", filename, string(fileContent))
	return err

}

func WriteToFile(filename, content string) error {
	file, err := os.OpenFile(filename, os.O_APPEND, 0644)

	if err != nil {
		err = os.WriteFile(filename, []byte(content), 0644)
		return err
	} else {
		_, err = file.WriteString(content)
		file.Close()
		return err
	}
}

func deleteFile(filename string) error {
	return os.Remove(filename)
}

func CreateOrCaptureBaseDir() (string, error) {
	var err error
	var baseDir string
	fmt.Print("\nEnter base directory: ")
	fmt.Scan(&baseDir)

	_, err = os.Stat(baseDir)

	if err != nil {
		err = os.Mkdir(baseDir, 0644)
		if err != nil {
			return "", err
		}
	}
	return baseDir, err
}

func CreateLogsDirectoryAndFileIfItDoesNotExist() error {
	var err error

	_, err = os.Stat(LOG_DIR)

	if err != nil {
		err = os.Mkdir(LOG_DIR, 0644)
		if err != nil {
			return err
		}
	}

	filePath := filepath.Join(LOG_DIR, SUSPENDED_USERS_FILE)

	_, err = os.Stat(filePath)
	if err != nil {
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		file.Close()
	}

	return nil
}

func ProcessFileOperation(baseDir string, choice int) error {
	var fileName string
	var filePath string

	fmt.Print("\nEnter file name: ")
	fmt.Scan(&fileName)

	filePath = filepath.Join(baseDir, fileName)

	fmt.Print("\n\nPerforming requested operation on:", filePath)

	switch choice {
	case READ_FILE:
		return readFromFile(filePath)
	case CREATE_FILE:
		return createFile(filePath)
	case WRITE_FILE:
		var content string
		fmt.Print("\nEnter content to append:")
		fmt.Scan(&content)
		return WriteToFile(filePath, content)
	case DELETE_FILE:
		return deleteFile(filePath)
	default:
		panic("\nFatal error in UI package")
	}
}
