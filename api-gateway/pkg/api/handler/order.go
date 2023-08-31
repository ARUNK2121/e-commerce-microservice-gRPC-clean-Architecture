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

type OrderHandler struct {
	GRPC_Client interfaces.OrderClient
}

func NewOrderHandler(client interfaces.OrderClient) *OrderHandler {
	return &OrderHandler{
		GRPC_Client: client,
	}
}

func (i *OrderHandler) MakeOrder(c *gin.Context) {

	fmt.Println("reaches cart handler")

	cart_id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	address_id, err := strconv.Atoi(c.Query("address_id"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	res, err := i.GRPC_Client.MakeOrder(cart_id, address_id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not add the order", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "order Successfull", res, nil)
	c.JSON(http.StatusOK, successRes)

}

func (i *OrderHandler) GetOrders(c *gin.Context) {

	user_id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	res, err := i.GRPC_Client.GetOrders(user_id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not fetch", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully fetched all requests", res, nil)
	c.JSON(http.StatusOK, successRes)

}

func (i *OrderHandler) EditOrderStatus(c *gin.Context) {

	var edit models.EditStatus
	if err := c.ShouldBindJSON(&edit); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are wrong", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	res, err := i.GRPC_Client.EditOrderStatus(edit)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not fetch", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully fetched all requests", res, nil)
	c.JSON(http.StatusOK, successRes)

}
