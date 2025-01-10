package database

import (
	"demo/models"
	"log"

	"gorm.io/gorm"
)

var ChAudit chan *models.Audit = make(chan *models.Audit, 100)

type AuditDB struct {
	DB *gorm.DB
}

type IAudit interface {
	Audit()
}

func NewAudit(db *gorm.DB) IAudit {
	return &AuditDB{DB: db}
}

func (a *AuditDB) Audit() {
	a.DB.AutoMigrate(&models.Audit{})
	for audit := range ChAudit {
		tx := a.DB.Create(audit)
		if tx.Error != nil {
			log.Println(tx.Error.Error())
		} else {
			log.Println("Audit success with ID: ", audit.Id)
		}
	}
}
