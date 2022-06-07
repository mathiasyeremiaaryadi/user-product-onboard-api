package deliveries

import (
	"net/http"
	"user-product-service/dto"
	"user-product-service/usecases"

	"github.com/gin-gonic/gin"
)

type RoleDelivery struct {
	roleUseCase usecases.RoleUseCase
}

func NewRoleDelivery(roleUseCase usecases.RoleUseCase) RoleDelivery {
	return RoleDelivery{
		roleUseCase: roleUseCase,
	}
}

func (roleDelivery RoleDelivery) GetRolesDelivery(c *gin.Context) {
	roleRespsonseBody := dto.ResponseBody{}
	roles, err := roleDelivery.roleUseCase.GetRolesUseCase()

	if err != nil {
		roleRespsonseBody.
			SetStatus("failed").
			SetStatusCode(500).
			SetError("Failed to retrieve roles, server disconnected or error").
			SetData(nil)

		c.AbortWithStatusJSON(http.StatusInternalServerError, roleRespsonseBody.GetResponseBody())

		return
	}

	roleRespsonseBody.
		SetStatus("ok").
		SetStatusCode(200).
		SetError(nil).
		SetData(roles)
	c.JSON(http.StatusOK, roleRespsonseBody.GetResponseBody())
}
