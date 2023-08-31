package handler

import (
	interfaces "api_gateway/pkg/client/interface"
	"api_gateway/pkg/utils/models"
	"api_gateway/pkg/utils/response"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	GRPC_Client interfaces.UserClient
}

func NewUserHandler(UserClient interfaces.UserClient) *UserHandler {
	return &UserHandler{
		GRPC_Client: UserClient,
	}
}

func (cr *UserHandler) LoginHandler(c *gin.Context) {

	log.Print(1)

	var UserDetails models.UserLogin
	if err := c.ShouldBindJSON(&UserDetails); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	log.Print(2)

	admin, err := cr.GRPC_Client.UserLogin(UserDetails)
	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "cannot authenticate user", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	log.Print(3)
	successRes := response.ClientResponse(http.StatusOK, "user authenticated successfully", admin, nil)
	c.JSON(http.StatusOK, successRes)

}

func (cr *UserHandler) UserSignUp(c *gin.Context) {

	var User models.UserSignUp
	if err := c.ShouldBindJSON(&User); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are wrong", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	UserDetails, err := cr.GRPC_Client.UserSignUp(User)
	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "cannot authenticate user", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusCreated, "Successfully signed up the user", UserDetails, nil)
	c.JSON(http.StatusCreated, successRes)

}
