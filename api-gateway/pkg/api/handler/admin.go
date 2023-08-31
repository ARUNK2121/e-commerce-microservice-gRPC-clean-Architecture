package handler

import (
	interfaces "api_gateway/pkg/client/interface"
	"api_gateway/pkg/utils/models"
	"api_gateway/pkg/utils/response"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	GRPC_Client interfaces.AdminClient
}

func NewAdminHandler(adminClient interfaces.AdminClient) *AdminHandler {
	return &AdminHandler{
		GRPC_Client: adminClient,
	}
}

func (cr *AdminHandler) LoginHandler(c *gin.Context) { // login handler for the admin

	log.Print(1)

	var adminDetails models.AdminLogin
	if err := c.ShouldBindJSON(&adminDetails); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	log.Print(2)

	admin, err := cr.GRPC_Client.AdminLogin(adminDetails)
	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "cannot authenticate user", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	log.Print(3)
	successRes := response.ClientResponse(http.StatusOK, "Admin authenticated successfully", admin, nil)
	c.JSON(http.StatusOK, successRes)

}

func (cr *AdminHandler) AdminSignUp(c *gin.Context) {

	var admin models.AdminSignUp
	if err := c.ShouldBindJSON(&admin); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are wrong", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	adminDetails, err := cr.GRPC_Client.AdminSignUp(admin)
	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "cannot authenticate user", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusCreated, "Successfully signed up the user", adminDetails, nil)
	c.JSON(http.StatusCreated, successRes)

}
