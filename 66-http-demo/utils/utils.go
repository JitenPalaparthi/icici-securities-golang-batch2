package utils

import (
	"demo/models"
	"log"
	"os"
)

var ChUser chan *models.User

func init() {
	ChUser = make(chan *models.User)
	go Process("users.txt")
}

func Process(filename string) {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err.Error())
	}
	defer f.Close()

	for user := range ChUser {
		bytes, err := user.ToBytes()
		if err != nil {
			log.Println(err.Error())
		}
		_, err = f.WriteString(string(bytes) + "\n")
		if err != nil {
			log.Println(err.Error())
		}
	}
}

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
