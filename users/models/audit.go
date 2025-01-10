package models

import (
	"encoding/json"
	"errors"
)

type Audit struct {
	Id           uint   `json:"id" gorm:"primaryKey"`
	RecordId     uint   `json:"record_id"`
	Table        string `json:"table"`
	Action       string `json:"action"`
	Data         string `json:"email"`
	Status       string `json:"status"`
	LastModified int64  `json:"last_modified"`
}

func (a *Audit) ToBytes() ([]byte, error) {
	if a == nil {
		return nil, errors.New("nil audit")
	}
	return json.Marshal(a)
}

// json, protocol buffers, xmlx
// reflection..
