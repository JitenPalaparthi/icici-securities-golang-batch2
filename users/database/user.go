package database

//go:generate mockgen -source=user.go -destination=mocks/mocks_user.go -package=mocks
import (
	"demo/models"
	"errors"

	"gorm.io/gorm"
)

type IUser interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUserByID(id uint) (*models.User, error)
	DeleteByID(id uint) (uint, error)
}

type UserDB struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) IUser {
	return &UserDB{
		DB: db,
	}
}

func (u *UserDB) CreateUser(user *models.User) (*models.User, error) {
	if u.DB == nil {
		return nil, ErrInvalidDBInstance
	}
	err := u.DB.AutoMigrate(&models.User{})
	if err != nil {
		return nil, errors.New("unable to automatically migrate the table")
	}
	tx := u.DB.Create(user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}

func (u *UserDB) GetUserByID(id uint) (*models.User, error) {
	if u.DB == nil {
		return nil, ErrInvalidDBInstance
	}
	user := new(models.User)
	tx := u.DB.First(user, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}

func (u *UserDB) DeleteByID(id uint) (uint, error) {
	if u.DB == nil {
		return 0, ErrInvalidDBInstance
	}

	tx := u.DB.Delete(models.User{}, id)

	if tx.Error != nil {
		return 0, tx.Error
	}

	if tx.RowsAffected == 0 {
		return 0, errors.New("no record to delete")
	}
	return uint(tx.RowsAffected), nil
}
