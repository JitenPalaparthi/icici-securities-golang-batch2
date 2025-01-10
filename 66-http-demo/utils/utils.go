package utils

import (
	"demo/models"
	"os"
)

func SaveToFile(filename string, user *models.User) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	bytes, err := user.ToBytes()
	if err != nil {
		return err
	}
	_, err = f.WriteString(string(bytes) + "\n")
	if err != nil {
		return err
	}
	return nil

}
