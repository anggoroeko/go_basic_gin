package main

import (
	"golang_basic_gin_sept_2023/config"
	"golang_basic_gin_sept_2023/middlewares"
	"golang_basic_gin_sept_2023/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()

	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "PONG",
		})
	})
	r.GET("/home", GetHome)

	route := r.Group("/")
	{
		user := route.Group("/user")
		{
			user.POST("/register", routes.RegisterUser)
			user.POST("/login", routes.GenerateToken)
		}

		department := route.Group("departments").Use(middlewares.IsAdmin())
		{
			//:: MASTER DEPARTMENT
			department.GET("/", routes.GetDepartment)
			department.GET("/:id", routes.GetDepartmentById)
			department.POST("/", routes.PostDepartment)
			department.PUT("/:id", routes.PutDepartment)
			department.DELETE("/:id", routes.DeleteDepartment)
		}

		position := route.Group("positions").Use(middlewares.Auth())
		{
			//:: MASTER POSITIONS
			position.GET("/", routes.GetPositions)
			position.GET("/:id", routes.GetPositionsById)
			position.POST("", routes.PostPositions)
			position.PUT("/:id", routes.PutPositions)
			position.DELETE("/:id", routes.DeletePositions)
		}

		inventory := route.Group("inventory")
		{
			//:: MASTER INVENTORY
			inventory.GET("", routes.GetInventory)
			inventory.GET("/:id", routes.GetInventoryById)
			inventory.POST("/", routes.PostInventory)
			inventory.PUT("/:id", routes.PutInventory)
			inventory.DELETE("/:id", routes.DeleteInventory)
		}

		employee := route.Group("employees")
		{
			//:: MASTER EMPLOYEE
			employee.GET("/", routes.GetEmployees)
			employee.GET("/:id", routes.GetEmployeeById)
			employee.POST("/", routes.PostEmployees)
			employee.PUT("/:id", routes.PutEmployees)
			employee.DELETE("/:id", routes.DeleteEmployees)
		}

		rental := route.Group("rental")
		{
			//:: RENTAL
			rental.GET("/", routes.GetRental)
			rental.POST("/employee", routes.RentalByEmployeeID)
			rental.GET("/inventory/:id", routes.GetRentalByInventoryID)
		}
	}

	r.Run(":8010") //listen and serve on 0.0.0.0:8080
}

func GetHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"Message": "Welcome Home!",
	})
}
