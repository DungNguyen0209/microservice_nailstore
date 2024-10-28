package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/minhdung/nailstore/internal/domain/request"
	interfaceObject "github.com/minhdung/nailstore/internal/interface"
	"github.com/minhdung/nailstore/internal/util"
)

type AccountController struct {
	accountUseCase interfaceObject.UserUsecase
}

func NewAccountController(service interfaceObject.UserUsecase) *AccountController {
	return &AccountController{accountUseCase: service}
}

func (controller *AccountController) CreateAccount(ctx *gin.Context) {
	var req request.UserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrResponse(err))
		return
	}

	err := controller.accountUseCase.CreateUser(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, req)
}

func (controller *AccountController) GetAccount(ctx *gin.Context) {
	var req request.GetUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, _ := uuid.Parse(req.Id)

	account, err := controller.accountUseCase.FindUserById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, account)
}
