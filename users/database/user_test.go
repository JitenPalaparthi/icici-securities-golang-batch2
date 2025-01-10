package database_test

import (
	"demo/database/mocks"
	"demo/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	usermock := mocks.NewMockIUser(ctrl)
	user := &models.User{Name: "jiten", Email: "JitenP@Outlook.com", Contact: "9618558500", Status: "active"}
	usermock.EXPECT().CreateUser(gomock.Any()).Return(user, nil).Times(1)
	actuakluser, err := usermock.CreateUser(user)

	// if err != nil {
	// 	t.Fail()
	// }
	// if user != actuakluser {
	// 	t.Fail()
	// }

	assert.NoError(t, err)
	assert.Equal(t, user, actuakluser)
}
