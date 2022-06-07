package deliveries

import (
	"net/http"
	"user-product-service/dto"
	"user-product-service/usecases"

	"github.com/gin-gonic/gin"
)

type UserDelivery struct {
	userUseCase usecases.UserUseCase
}

func NewUserDelivery(userUseCase usecases.UserUseCase) UserDelivery {
	return UserDelivery{
		userUseCase: userUseCase,
	}
}

func (userDelivery UserDelivery) GetUsersDelivery(c *gin.Context) {
	userResponseBody := dto.ResponseBody{}
	users, err := userDelivery.userUseCase.GetUsersUseCase()

	if err != nil {
		userResponseBody.
			SetStatus("failed").
			SetStatusCode(500).
			SetError("Failed to retrieve users, server disconnected or error").
			SetData(nil)

		c.AbortWithStatusJSON(http.StatusInternalServerError, userResponseBody.GetResponseBody())
	}

	userResponseBody.
		SetStatus("ok").
		SetStatusCode(200).
		SetError(nil).
		SetData(users)

	c.JSON(http.StatusOK, userResponseBody.GetResponseBody())
}

func (userDelivery UserDelivery) GetUserDelivery(c *gin.Context) {
	userResponseBody := dto.ResponseBody{}
	userID := c.Param("id")

	user, err := userDelivery.userUseCase.GetUserUseCase(userID)
	if err != nil {
		userResponseBody.
			SetStatus("failed").
			SetStatusCode(404).
			SetError("Failed to retrieve user, user not found").
			SetData(nil)

		c.AbortWithStatusJSON(http.StatusNotFound, userResponseBody.GetResponseBody())

		return
	}

	userResponseBody.
		SetStatus("ok").
		SetStatusCode(200).
		SetError(nil).
		SetData(user)

	c.JSON(http.StatusOK, userResponseBody.GetResponseBody())
}

func (userDelivery UserDelivery) CreateUserDelivery(c *gin.Context) {
	userPostRequestBody := dto.UserPostRequestBody{}
	userResponseBody := dto.ResponseBody{}

	if err := c.ShouldBindJSON(&userPostRequestBody); err != nil {
		userResponseBody.
			SetStatus("failed").
			SetStatusCode(400).
			SetError("Failed to bind data, data is invalid").
			SetData(nil)

		c.AbortWithStatusJSON(http.StatusBadRequest, userResponseBody.GetResponseBody())

		return
	}

	createdUser, err := userDelivery.userUseCase.CreateUserUseCase(userPostRequestBody)
	if err != nil {
		userResponseBody.
			SetStatus("failed").
			SetStatusCode(500).
			SetError("Failed to create user, server disconnected or error").
			SetData(nil)

		c.AbortWithStatusJSON(http.StatusInternalServerError, userResponseBody.GetResponseBody())

		return
	}

	userResponseBody.
		SetStatus("ok").
		SetStatusCode(201).
		SetError(nil).
		SetData(map[string]interface{}{
			"id": createdUser.ID,
		})

	c.JSON(http.StatusCreated, userResponseBody.GetResponseBody())
}

func (userDelivery UserDelivery) UpdateUserDelivery(c *gin.Context) {
	userPutRequestBody := dto.UserPutRequestBody{}
	userResponseBody := dto.ResponseBody{}
	userID := c.Param("id")

	if err := c.ShouldBindJSON(&userPutRequestBody); err != nil {
		userResponseBody.
			SetStatus("failed").
			SetStatusCode(400).
			SetError("Failed to bind data, data is invalid").
			SetData(nil)

		c.AbortWithStatusJSON(http.StatusBadRequest, userResponseBody.GetResponseBody())

		return
	}

	updatedUserID, err := userDelivery.userUseCase.UpdateUserUseCase(userPutRequestBody, userID)
	if err != nil {
		userResponseBody.
			SetStatus("failed").
			SetStatusCode(500).
			SetError("Failed to update user, server disconnected or error").
			SetData(nil)

		c.AbortWithStatusJSON(http.StatusInternalServerError, userResponseBody.GetResponseBody())

		return
	}

	userResponseBody.
		SetStatus("ok").
		SetStatusCode(200).
		SetError(nil).
		SetData(map[string]interface{}{
			"id": updatedUserID,
		})

	c.JSON(http.StatusOK, userResponseBody.GetResponseBody())
}

func (userDelivery UserDelivery) DeleteUserDelivery(c *gin.Context) {
	userResponseBody := dto.ResponseBody{}
	userID := c.Param("id")

	err := userDelivery.userUseCase.DeleteUserUseCase(userID)
	if err != nil {
		userResponseBody.
			SetStatus("failed").
			SetStatusCode(404).
			SetError("Failed to delete user, user not found").
			SetData(nil)

		c.AbortWithStatusJSON(http.StatusNotFound, userResponseBody.GetResponseBody())

		return
	}

	userResponseBody.
		SetStatus("ok").
		SetStatusCode(200).
		SetError(nil).
		SetData(nil)

	c.JSON(http.StatusOK, userResponseBody.GetResponseBody())
}
