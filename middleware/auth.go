package middleware

import (
	"backend-onboarding/repository/user_repository"
	"backend-onboarding/usecase/user_usecase"
	"backend-onboarding/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func JWTAuth(jwtUsecase user_usecase.JwtUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		// fmt.Println(authHeader)
		userId, err := jwtUsecase.ValidateTokenAndGetUserId(authHeader)
		if err != nil {
			resp := utils.ResponseError("You are unathorized", err, 403)
			c.AbortWithStatusJSON(resp.StatusCode, resp)
			return
		}
		c.Set("user_id", userId)
	}
}

func JWTAuthAdmin(jwtUsecase user_usecase.JwtUseCase, userRepo user_repository.UserRepository) gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		userId, err := jwtUsecase.ValidateTokenAndGetUserId(authHeader)
		if err != nil {
			errorRes := utils.ResponseError("You are unathorized", err, 403)
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorRes)
			return
		}

		userUUID, err := uuid.Parse(userId)

		if err != nil {
			errorRes := utils.ResponseError("Data not found", err, 404)
			c.AbortWithStatusJSON(http.StatusNotFound, errorRes)
			return
		}
		user, err := userRepo.UserDetailById(userUUID)
		if err != nil {
			errorRes := utils.ResponseError("Internal Server Error", err, 500)
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorRes)
			return
		}

		if user.Title != "admin" {
			errorRes := utils.ResponseError("You are unathorized", err, 401)
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorRes)
			return
		}
		c.Set("user_id", userId)
	}
}

func JWTAuthChecker(jwtUsecase user_usecase.JwtUseCase, userRepo user_repository.UserRepository) gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		userId, err := jwtUsecase.ValidateTokenAndGetUserId(authHeader)
		if err != nil {
			errorRes := utils.ResponseError("You are unathorized", err, 401)
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorRes)
			return
		}

		userUUID, err := uuid.Parse(userId)

		if err != nil {
			errorRes := utils.ResponseError("Data not found", err, 404)
			c.AbortWithStatusJSON(http.StatusNotFound, errorRes)
			return
		}
		user, err := userRepo.UserDetailById(userUUID)
		if err != nil {
			errorRes := utils.ResponseError("Internal Server Error", err, 500)
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorRes)
			return
		}
		role, err := userRepo.GetRoleByRoleId(user.RoleId)
		if err != nil {
			errorRes := utils.ResponseError("Internal Server Error", err, 500)
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorRes)
			return
		}
		fmt.Println(role.Title)
		if !(role.Title == "checker" || role.Title == "admin") {
			errorRes := utils.ResponseError("You are unathorized", err, 401)
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorRes)
			return
		}
		c.Set("user_id", userId)
	}

}

func JWTAuthSigner(jwtUsecase user_usecase.JwtUseCase, userRepo user_repository.UserRepository) gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		userId, err := jwtUsecase.ValidateTokenAndGetUserId(authHeader)
		if err != nil {
			errorRes := utils.ResponseError("You are unathorized", err, 401)
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorRes)
			return
		}

		userUUID, err := uuid.Parse(userId)

		if err != nil {
			errorRes := utils.ResponseError("Data not found", err, 404)
			c.AbortWithStatusJSON(http.StatusNotFound, errorRes)
			return
		}
		user, err := userRepo.UserDetailById(userUUID)
		if err != nil {
			errorRes := utils.ResponseError("Internal Server Error", err, 500)
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorRes)
			return
		}
		fmt.Println(user.Title)
		if user.Title != "signer" {
			if user.Title != "admin" {
				errorRes := utils.ResponseError("You are unathorized", err, 401)
				c.AbortWithStatusJSON(http.StatusUnauthorized, errorRes)
				return
			}
		}
		c.Set("user_id", userId)
	}
}
