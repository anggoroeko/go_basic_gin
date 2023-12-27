package main

import (
	"golang_basic_gin_sept_2023/config"
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

	r.GET("/departments", routes.GetDepartment)
	r.GET("/departments/:id", routes.GetDepartmentById)
	r.POST("/departments", routes.PostDepartment)
	r.PUT("/departments/:id", routes.PutDepartment)
	r.DELETE("/departments/:id", routes.DeleteDepartment)

	r.GET("/positions", routes.GetPositions)
	r.GET("/positions/:id", routes.GetPositionsById)
	r.POST("/positions", routes.PostPositions)
	r.PUT("/positions/:id", routes.PutPositions)
	r.DELETE("/positions/:id", routes.DeletePositions)

	r.GET("/inventory", routes.GetInventory)
	r.GET("/inventory/:id", routes.GetInventoryById)
	r.POST("/inventory", routes.PostInventory)
	r.PUT("/inventory/:id", routes.PutInventory)
	r.DELETE("/inventory/:id", routes.DeleteInventory)

	r.GET("/employees", routes.GetEmployees)
	r.GET("/employees/:id", routes.GetEmployeeById)
	r.POST("/employees", routes.PostEmployees)
	r.PUT("/employees/:id", routes.PutEmployees)
	r.DELETE("/employees/:id", routes.DeleteEmployees)

	r.Run(":8010") //listen and serve on 0.0.0.0:8080
}

func GetHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"Message": "Welcome Home!",
	})
}
