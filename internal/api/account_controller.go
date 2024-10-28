package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/minhdung/nailstore/internal/domain/request"
	"github.com/minhdung/nailstore/internal/interface/usecases"
)

type AccountController struct {
	accountUseCase usecases.UserUsecase
}

func NewAccountController(service usecases.UserUsecase) *AccountController {
	return &AccountController{accountUseCase: service}
}

func (controller *AccountController) CreateAccount(ctx *gin.Context) {
	var req request.UserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	err := controller.accountUseCase.CreateUser(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
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
