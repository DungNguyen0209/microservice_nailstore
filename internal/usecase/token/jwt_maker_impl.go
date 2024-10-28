package usecase

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/minhdung/nailstore/internal/domain/entity"
	interfaceObject "github.com/minhdung/nailstore/internal/interface"
)

const minSecretKeySize = 32

// JWTMakerImpl is make token
type JWTMakerImpl struct {
	secretkey string
}

func NewJWTMaker(secretKey string) (interfaceObject.Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("Invalid Key size: must be at least %d character", minSecretKeySize)
	}
	return &JWTMakerImpl{secretKey}, nil
}

func (maker *JWTMakerImpl) CreateToken(username string, duration time.Duration) (string, *entity.PayLoad, error) {
	payload, err := entity.NewPayLoad(username, duration)
	if err != nil {
		return "", payload, err
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(maker.secretkey))
	return token, payload, err
}

// Verify check if token is valid or not
func (maker *JWTMakerImpl) VerifyToken(token string) (*entity.PayLoad, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, entity.InvalidToken
		}
		return []byte(maker.secretkey), nil
	}
	jwtToken, err := jwt.ParseWithClaims(token, &entity.PayLoad{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, entity.ErrExpiredToken) {
			return nil, entity.ErrExpiredToken
		}
		return nil, entity.InvalidToken
	}
	payload, ok := jwtToken.Claims.(*entity.PayLoad)
	if !ok {
		return nil, entity.InvalidToken
	}
	return payload, nil
}
