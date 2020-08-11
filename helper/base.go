package helper

import (
	"github.com/google/uuid"
	"os"
)

func GetProjectRoot() (path string, err error) {
	path, err = os.Getwd()
	return
}

func GetUuid() string {
	return uuid.Must(uuid.NewRandom()).String()
}
