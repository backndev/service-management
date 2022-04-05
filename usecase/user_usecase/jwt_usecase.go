package user_usecase

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type CustomClaim struct {
	jwt.StandardClaims
	Role   string `json:"role"`
	UserID string `json:"user_id"`
}

func (jwtAuth *jwtUseCase) GenerateToken(userId string, roleId string) (string, error) {
	roleUUID, _ := uuid.Parse(roleId)

	data, error := jwtAuth.userRepo.GetRoleByRoleId(roleUUID)
	if error != nil {
		return "role not found", error
	}

	claim := CustomClaim{
		UserID: userId,
		Role:   data.Title,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
			Issuer:    os.Getenv("APP_NAME"),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claim)
	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

func (jwtAuth *jwtUseCase) ValidateToken(token string) (*jwt.Token, error) {

	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
}

func (jwtAuth *jwtUseCase) ValidateTokenAndGetUserId(token string) (string, error) {
	validatedToken, err := jwtAuth.ValidateToken(token)
	if err != nil {
		return "", err
	}

	claims, ok := validatedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("failed to claim token")
	}

	return claims["user_id"].(string), nil
}
func (jwtAuth *jwtUseCase) ValidateTokenAndGetRole(token string) (string, error) {
	validatedToken, err := jwtAuth.ValidateToken(token)
	if err != nil {
		return "", err
	}

	claims, ok := validatedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("failed to claim token")
	}

	return claims["role"].(string), nil
}
