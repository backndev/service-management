package user_delivery

import (
	"backend-onboarding/model/dto"
	"backend-onboarding/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (user *userDelivery) UserList(c *gin.Context) {
	response := user.userUseCase.UserList()
	if response.Status != "ok" {
		c.JSON(response.StatusCode, response)
		return
	}
	c.JSON(http.StatusOK, response)
}

func (user *userDelivery) UserDetailById(c *gin.Context) {
	id := c.Param("id")
	response := user.userUseCase.UserDetailById(id)
	if response.StatusCode == http.StatusNotFound {
		c.JSON(http.StatusOK, response)
		return
	}

	if response.Status != "ok" {
		c.JSON(response.StatusCode, response)
		return
	}
	c.JSON(http.StatusOK, response)
}

func (user *userDelivery) InsertNewUser(c *gin.Context) {
	request := dto.User{}
	if err := c.ShouldBindJSON(&request); err != nil {
		errorRes := utils.ResponseError("Bad Request", err, 400)
		c.JSON(errorRes.StatusCode, errorRes)
		return
	}
	response := user.userUseCase.InsertNewUser(request)

	if response.Status != "ok" {
		c.JSON(response.StatusCode, response)
		return
	}
	c.JSON(http.StatusCreated, response)
}

func (user *userDelivery) UpdateUserData(c *gin.Context) {
	id := c.Param("id")
	request := dto.User{}
	if err := c.ShouldBindJSON(&request); err != nil {
		errorReq := utils.ResponseError("Bad Request", err, 400)
		c.JSON(errorReq.StatusCode, errorReq)
		return
	}

	response := user.userUseCase.UpdateUserData(request, id)

	if response.Status != "ok" {
		c.JSON(response.StatusCode, response)
		return
	}

	c.JSON(response.StatusCode, response)
}

func (user *userDelivery) DeleteUserById(c *gin.Context) {
	id := c.Param("id")
	response := user.userUseCase.DeleteUserById(id)
	if response.Status != "ok" {
		c.JSON(response.StatusCode, response)
		return
	}

	c.JSON(response.StatusCode, response)
}

func (user *userDelivery) UserLogin(c *gin.Context) {
	request := dto.UserLogin{}
	if err := c.ShouldBindJSON(&request); err != nil {
		errorRes := utils.ResponseError("Bad Request", err, 400)
		c.JSON(errorRes.StatusCode, errorRes)
		return
	}
	response := user.userUseCase.UserLogin(request)

	if response.Status != "ok" {
		c.JSON(response.StatusCode, response)
		return
	}
	c.JSON(response.StatusCode, response)
}
