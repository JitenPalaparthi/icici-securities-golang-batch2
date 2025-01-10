package handlers

import (
	"demo/database"
	"demo/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type IUserHandler interface {
	CreateUser(*gin.Context)
	GetUserByID(ctx *gin.Context)
	DeleteByID(ctx *gin.Context)
}

type UserHandler struct {
	IUser database.IUser
}

func NewUserHandler(user database.IUser) IUserHandler {
	return &UserHandler{
		IUser: user,
	}
}

func (u *UserHandler) CreateUser(ctx *gin.Context) {
	user := new(models.User)
	err := ctx.Bind(user)
	if err != nil {
		ctx.String(400, err.Error())
		ctx.Abort()
		return
	}

	err = user.Validate()
	if err != nil {
		ctx.String(400, err.Error())
		ctx.Abort()
		return
	}
	user.Status = "active"
	user.LastModified = time.Now().Unix()

	user, err = u.IUser.CreateUser(user)
	if err != nil {
		ctx.String(400, err.Error())
		ctx.Abort()
		return
	}
	data, _ := user.ToBytes()
	database.ChAudit <- &models.Audit{Action: "create", Table: "user", RecordId: user.Id, Status: "active", LastModified: time.Now().Unix(), Data: string(data)}
	ctx.JSON(201, user)
}

func (u *UserHandler) GetUserByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.String(400, err.Error())
		ctx.Abort()
		return
	}

	user, err := u.IUser.GetUserByID(uint(id))
	if err != nil {
		ctx.String(400, err.Error())
		ctx.Abort()
		return
	}
	ctx.JSON(200, user)
}

func (u *UserHandler) DeleteByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.String(400, err.Error())
		ctx.Abort()
		return
	}

	r, err := u.IUser.DeleteByID(uint(id))
	if err != nil {
		ctx.String(400, err.Error())
		ctx.Abort()
		return
	}
	ctx.JSON(202, r)
}
