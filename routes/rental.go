package routes

import (
	"golang_basic_gin_sept_2023/config"
	"golang_basic_gin_sept_2023/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func GetRental(c *gin.Context) {
	EmployeeInventory := []models.EmployeeInventory{}

	config.DB.Preload(clause.Associations).Find(&EmployeeInventory)

	responseGetRental := []models.ResponseGetRental{}

	for _, ei := range EmployeeInventory {
		rgr := models.ResponseGetRental{
			ID:            ei.ID,
			Description:   ei.Description,
			EmployeeName:  ei.Employee.Name,
			InventoryName: ei.Inventory.Name,
			CreatedAt:     ei.CreatedAt,
		}

		responseGetRental = append(responseGetRental, rgr)
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"message": "Get Rental Successfully",
		"data":    responseGetRental,
	})

}

func RentalByEmployeeID(c *gin.Context) {
	var reqRental models.RequestRental

	if err := c.ShouldBindJSON(&reqRental); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"error":   err.Error(),
		})

		c.Abort()
		return
	}

	rental := models.EmployeeInventory{
		EmployeeID:  reqRental.EmployeeID,
		InventoryID: reqRental.InventoryID,
		Description: reqRental.Description,
	}

	insert := config.DB.Create(&rental)

	if insert.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"error":   insert.Error.Error(),
		})

		c.Abort()
		return
	}

	respRental := models.ResponseRental{
		ID:          rental.ID,
		EmployeeID:  rental.EmployeeID,
		InventoryID: rental.InventoryID,
		Description: rental.Description,
		CreatedAt:   rental.CreatedAt,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Insert Rental Successfully",
		"data":    respRental,
	})
}

func GetRentalByInventoryID(c *gin.Context) {
	id := c.Param("id")

	inventories := models.Inventory{}
	emInv := []models.ResponseEmployeeInventory{}

	config.DB.Preload(clause.Associations).First(&inventories, "id = ?", id)

	for _, inv := range inventories.Employee {
		emInvRes := models.ResponseEmployeeInventory{
			EmployeeID:  inv.EmployeeID,
			InventoryID: inv.InventoryID,
			Description: inv.Description,
			CreatedAt:   inv.CreatedAt,
		}
		emInv = append(emInv, emInvRes)
	}

	respInv := models.ResponseInventoryEmployee{
		InventoryName:        inventories.Name,
		InventoryDescription: inventories.Description,
		EmployeeInventory:    emInv,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Rental Data By Inventory ID!",
		"data":    respInv,
	})
}
