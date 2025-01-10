package database

import (
	"errors"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	RETRY_TIMES    = 5
	RETRY_DURATION = time.Second * 5
)

var (
	ErrInvalidDBInstance = errors.New("invalid db instancer")
)

func GetConnection(dsn string) (*gorm.DB, error) {
	count := 0
try:
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	count++
	if err != nil {
		log.Println(err.Error())
		log.Println("Retrying to connect to database..." + string(count))
		time.Sleep(RETRY_DURATION)
		if count < RETRY_TIMES {
			goto try
		}
	}
	return db, err
}

func GetMySQLConnection(dsn string) (*gorm.DB, error) {
	count := 0
try:
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	count++
	if err != nil {
		log.Println(err.Error())
		log.Println("Retrying to connect to database..." + string(count))
		time.Sleep(RETRY_DURATION)
		if count < RETRY_TIMES {
			goto try
		}
	}
	return db, err
}
