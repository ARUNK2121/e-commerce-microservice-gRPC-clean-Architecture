package handler

import (
	interfaces "api_gateway/pkg/client/interface"
	"api_gateway/pkg/utils/models"
	"api_gateway/pkg/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	GRPC_Client interfaces.ProductClient
}

func NewProductHandler(productClient interfaces.ProductClient) *ProductHandler {
	return &ProductHandler{
		GRPC_Client: productClient,
	}
}

func (cr *ProductHandler) AddProduct(c *gin.Context) {

	var inventories models.Inventories
	if err := c.ShouldBindJSON(&inventories); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are wrong", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	status, err := cr.GRPC_Client.AddProduct(inventories)
	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "could not add inventory", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusCreated, "Successfully added inventory", status, nil)
	c.JSON(http.StatusCreated, successRes)

}

func (cr *ProductHandler) ListProducts(c *gin.Context) {

	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page number not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	result, err := cr.GRPC_Client.ListProducts(page)
	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "could not get inventory", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusCreated, "Successfully fetched inventory", result, nil)
	c.JSON(http.StatusCreated, successRes)

}
