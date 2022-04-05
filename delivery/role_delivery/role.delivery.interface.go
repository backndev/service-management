package role_delivery

import (
	"backend-onboarding/usecase/role_usecase"
	"github.com/gin-gonic/gin"
)

type RoleDelivery interface {
	RoleList(*gin.Context)
}

type roleDelivery struct {
	roleUseCase role_usecase.RoleUseCase
}

func GetRoleDelivery(roleUseCase role_usecase.RoleUseCase) RoleDelivery {
	return &roleDelivery{
		roleUseCase: roleUseCase,
	}
}
