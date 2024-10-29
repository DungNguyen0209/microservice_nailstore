package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	entity_api "github.com/minhdung/nailstore/internal/domain/api"
	"github.com/minhdung/nailstore/internal/domain/entity"
	interfaceObject "github.com/minhdung/nailstore/internal/interface"
	token "github.com/minhdung/nailstore/internal/usecase/token"
	"github.com/minhdung/nailstore/internal/util"
	"gorm.io/gorm"
)

type AccountHandler struct {
	config         util.Config
	tokenMaker     interfaceObject.Maker
	accountUseCase interfaceObject.UserUsecase
}

func NewAccountHandler(config util.Config, service interfaceObject.UserUsecase) (*AccountHandler, error) {
	token, err := token.NewJWTMaker(config.TokenSymetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token marker : %w", &err)
	}
	return &AccountHandler{
		accountUseCase: service,
		tokenMaker:     token,
		config:         config,
	}, nil
}

func (handler *AccountHandler) CreateAccount(ctx *gin.Context) {
	var req entity_api.UserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrResponse(err))
		return
	}

	err := handler.accountUseCase.CreateUser(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, req)
}

func (handler *AccountHandler) GetAccount(ctx *gin.Context) {
	var req entity_api.GetUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, _ := uuid.Parse(req.Id)

	account, err := handler.accountUseCase.FindUserById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func (handler *AccountHandler) LoginUser(ctx *gin.Context) {
	var req entity_api.LogInUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrResponse(err))
		return
	}

	user, err := handler.accountUseCase.GetUserByName(req.Username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, util.ErrResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, util.ErrResponse(err))
		return
	}

	err = util.CheckPassWord(req.Password, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, util.ErrResponse(err))
		return
	}

	accessToken, accessPayload, err := handler.tokenMaker.CreateToken(req.Username, handler.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrResponse(err))
		return
	}

	refreshtoken, refreshPayload, err := handler.tokenMaker.CreateToken(
		user.Username,
		handler.config.RefreshTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrResponse(err))
		return
	}
	sessionId := uuid.New()
	session, err := handler.accountUseCase.CreateSession(entity.Session{
		Id:           sessionId,
		UserId:       user.Id,
		RefreshToken: refreshtoken,
		UserAgent:    ctx.Request.UserAgent(),
		ClientIP:     ctx.ClientIP(),
		IsBlocked:    false,
		ExpiredAt:    refreshPayload.ExpiredAt,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrResponse(err))
		return
	}

	rsp := entity_api.LogInUserResponse{
		SessionID:             session.Id,
		AccessToken:           accessToken,
		AccessTokenExpiredAt:  accessPayload.ExpiredAt,
		RefreshToken:          refreshtoken,
		RefreshTokenExpiredAt: refreshPayload.ExpiredAt,
		User:                  newUserResponse(*user),
	}
	ctx.JSON(http.StatusOK, rsp)
}
func newUserResponse(user entity.User) entity_api.UserRespone {
	return entity_api.UserRespone{
		Username:  user.Username,
		FullName:  user.FullName,
		Email:     user.Email,
		Phone:     user.Phone,
		Note:      user.Note,
		CreatedBy: user.CreatedBy.String(),
	}
}
