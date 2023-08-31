package handler

import (
	interfaces "api_gateway/pkg/client/interface"
	"api_gateway/pkg/utils/models"
	"api_gateway/pkg/utils/response"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	GRPC_Client interfaces.CartClient
}

func NewCartHandler(client interfaces.CartClient) *CartHandler {
	return &CartHandler{
		GRPC_Client: client,
	}
}

func (i *CartHandler) AddToCart(c *gin.Context) {

	fmt.Println("reaches cart handler")

	var model models.AddToCart
	if err := c.BindJSON(&model); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	if err := i.GRPC_Client.AddToCart(model); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not add the Cart", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully added To cart", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

func (i *CartHandler) GetCart(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "check parameters properly", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	products, err := i.GRPC_Client.GetCart(id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve cart", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully got all products in cart", products, nil)
	c.JSON(http.StatusOK, successRes)

}
