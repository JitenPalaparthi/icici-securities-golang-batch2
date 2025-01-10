package models

import (
	"encoding/json"
	"errors"
)

type User struct {
	//*gorm.Model // promoted field
	Id           uint   `json:"id" gorm:"primaryKey"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Contact      string `json:"contact"`
	Status       string `json:"status"`
	LastModified int64  `json:"last_modified"`
}

func (u *User) Validate() error {
	if u.Email == "" {
		return errors.New("invalid email")
	}
	if u.Name == "" {
		return errors.New("invalid name")
	}
	if u.Contact == "" {
		return errors.New("invalid contact")
	}
	return nil
}

func (u *User) ToBytes() ([]byte, error) {
	if u == nil {
		return nil, errors.New("nil user")
	}
	return json.Marshal(u)
}

// json, protocol buffers, xmlx
// reflection..
