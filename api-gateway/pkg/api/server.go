package server

import (
	"api_gateway/pkg/api/handler"
	"log"

	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(adminHandler *handler.AdminHandler, productHandler *handler.ProductHandler, userHandler *handler.UserHandler, cartHandler *handler.CartHandler, orderhandler *handler.OrderHandler) *ServerHTTP {
	router := gin.New()

	router.Use(gin.Logger())

	//admin side
	router.POST("/admin/login", adminHandler.LoginHandler)
	router.POST("/admin/signup", adminHandler.AdminSignUp)

	//user side
	router.POST("/user/signup", userHandler.UserSignUp)
	router.POST("/user/login", userHandler.LoginHandler)

	//product side
	router.GET("/user/home", productHandler.ListProducts)
	router.POST("/admin/inventories/add", productHandler.AddProduct)

	//cart side
	router.POST("/user/cart", cartHandler.AddToCart)
	router.GET("/user/cart", cartHandler.GetCart)

	//order
	router.POST("user/cart/order", orderhandler.MakeOrder)
	router.GET("user/orders", orderhandler.GetOrders)

	router.PATCH("admin/orders/status", orderhandler.EditOrderStatus)

	return &ServerHTTP{engine: router}
}

func (s *ServerHTTP) Start() {
	log.Printf("starting server on :3000")
	err := s.engine.Run(":3000")
	if err != nil {
		log.Printf("error while starting the server")
	}
}
