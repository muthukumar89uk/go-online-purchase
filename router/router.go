package router

import (
	//user defined packages
	"online/handler"
	"online/middleware"

	//Third party packages
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

// Signup and Login Handlers
func LoginHandlers(Db *gorm.DB, app *echo.Echo) {
	handler := handler.Database{Connection: Db}
	app.POST("/signup", handler.Signup)
	app.POST("/login", handler.Login)
}

// These handlers are accessible only by admin
func AdminHandlers(Db *gorm.DB, app *echo.Echo) {
	handler := handler.Database{Connection: Db}
	middleware := middleware.Database{Connection: Db}
	admin := app.Group("/admin", middleware.AuthMiddleware)
	admin.POST("/post-product", handler.PostProduct)
	admin.PUT("/update-product/:product_id", handler.UpdateProductById)
	admin.DELETE("/delete-product/:product_id", handler.DeleteProductById)
	admin.PUT("/update-status/:order_id", handler.UpdateOrderStatusById)
	admin.GET("/get-order-statuses", handler.GetAllOrderStatus)
}

// These handlers are accessible only by user
func UserHandlers(Db *gorm.DB, app *echo.Echo) {
	handler := handler.Database{Connection: Db}
	middleware := middleware.Database{Connection: Db}
	user := app.Group("/user", middleware.AuthMiddleware)
	user.POST("/post-order", handler.AddOrder)
	user.DELETE("/cancel-order/:order_id", handler.CancelOrderById)
	user.POST("/payment/:order_id", handler.Payment)
}

// These handlers are accessible by both admin and user
func CommonHandlers(Db *gorm.DB, app *echo.Echo) {
	handler := handler.Database{Connection: Db}
	middleware := middleware.Database{Connection: Db}
	common := app.Group("/common", middleware.AuthMiddleware)
	common.GET("/get-all-products", handler.GetAllProducts)
	common.GET("/get-orders", handler.GetOrders)
	common.GET("/get-order-status/:order_id", handler.GetOrderStatusById)
}
