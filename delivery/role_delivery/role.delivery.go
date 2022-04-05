package role_delivery

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (role *roleDelivery) RoleList(c *gin.Context) {
	result := role.roleUseCase.RoleList()
	if result.Status != "ok" {
		c.JSON(result.StatusCode, result)
		return
	}
	c.JSON(http.StatusOK, result)
}
