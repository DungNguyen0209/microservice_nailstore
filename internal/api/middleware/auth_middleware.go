package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	interfaceObject "github.com/minhdung/nailstore/internal/interface"
	"github.com/minhdung/nailstore/internal/util"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func AuthMiddleware(tokenMaker interfaceObject.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, util.ErrResponse(err))
			return
		}

		feilds := strings.Fields(authorizationHeader)
		if len(feilds) < 2 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, util.ErrResponse(err))
			return
		}

		authorizationType := strings.ToLower(feilds[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsupport authorization type %s", authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, util.ErrResponse(err))
			return
		}

		accessToken := feilds[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, util.ErrResponse(err))
			return
		}
		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()

	}
}
