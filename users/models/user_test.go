package models_test

import (
	"demo/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateUserSuccess(t *testing.T) {
	user := &models.User{Name: "Jiten", Email: "Jp@gmail.com", Contact: "9618558500"}
	err := user.Validate()
	assert.NoError(t, err)
	// if err != nil {
	// 	t.Fail()
	// 	t.Errorf("test failed")
	// }
}
func TestValidateUserFail(t *testing.T) {
	user := &models.User{Email: "Jp@gmail.com", Contact: "9618558500"}
	err := user.Validate()
	if err == nil {
		t.Fail()
		t.Errorf("test failed")
	}

	user = &models.User{Name: "Jiten", Contact: "9618558500"}
	err = user.Validate()
	if err == nil {
		t.Fail()
		t.Errorf("test failed")
	}
	user = &models.User{Name: "Jiten", Email: "Jitenp@outlook.com"}
	err = user.Validate()
	if err == nil {
		t.Fail()
		t.Errorf("test failed")
	}
}
