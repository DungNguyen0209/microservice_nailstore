package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	entity_api "github.com/minhdung/nailstore/internal/domain/api"
	interfaceObject "github.com/minhdung/nailstore/internal/interface"
	token "github.com/minhdung/nailstore/internal/usecase/token"
	"github.com/minhdung/nailstore/internal/util"
)

type UserHandler struct {
	config         util.Config
	tokenMaker     interfaceObject.Maker
	accountUseCase interfaceObject.UserUsecase
}

func NewUserHandler(config util.Config, service interfaceObject.UserUsecase) (*UserHandler, error) {
	token, err := token.NewJWTMaker(config.TokenSymetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token marker : %w", &err)
	}
	return &UserHandler{
		accountUseCase: service,
		tokenMaker:     token,
		config:         config,
	}, nil
}
func (handler *UserHandler) CreateUser(ctx *gin.Context) {
	var req entity_api.UserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrResponse(err))
		return
	}

	hashedPassword, err := util.HashPassWord(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrResponse(err))
	}
	req.Password = hashedPassword

	err = handler.accountUseCase.CreateUser(req)
	if err != nil {
		fmt.Println("can not create User")
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	rsp := entity_api.UserRespone{
		Username:  req.Username,
		FullName:  req.FullName,
		Email:     req.Email,
		Phone:     req.Phone,
		Note:      req.Note,
		CreatedBy: req.CreatedBy.String(),
	}
	ctx.JSON(http.StatusOK, rsp)
}
