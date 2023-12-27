package routes

import (
	"golang_basic_gin_sept_2023/config"
	"golang_basic_gin_sept_2023/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetInventory(c *gin.Context) {
	inventory := []models.Inventory{}

	//:: WITHOUT RELATION
	// config.DB.Find(&departments)

	//:: WITH RELATION
	config.DB.Preload("Archive").Find(&inventory)

	getInventoryResponse := []models.GetInventoryResponse{}

	for _, d := range inventory {
		inv := models.GetInventoryResponse{
			InventoryID:          d.ID,
			InventoryName:        d.Name,
			InventoryDescription: d.Description,
			ArchiveName:          d.Archive.Name,
			ArchiveDescription:   d.Archive.Description,
		}

		getInventoryResponse = append(getInventoryResponse, inv)
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Welcome Inventory",
		"data":    getInventoryResponse,
	})
}

func PostInventory(c *gin.Context) {
	reqInv := models.RequestInventory{}
	c.BindJSON(&reqInv)

	inventory := models.Inventory{
		Name:        reqInv.InventoryName,
		Description: reqInv.InventoryDescription,
		Archive: models.Archive{
			Name:        reqInv.ArchiveName,
			Description: reqInv.InventoryDescription,
		},
	}

	config.DB.Create(&inventory)

	c.JSON(http.StatusCreated, gin.H{
		"Message": "Insert Successfully",
		"data":    inventory,
	})
}

func GetInventoryById(c *gin.Context) {
	id := c.Param("id")

	inventory := models.Inventory{}

	//:: WITH RELATIONSHIP
	data := config.DB.Preload("Archive").First(&inventory, "id = ?", id)

	//:: VALIDATE DATA
	if data.Error != nil {
		log.Printf("Error GetInventoryById : %s ", data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not Found",
		})
		return
	}

	inv := models.GetInventoryResponse{
		InventoryID:          inventory.ID,
		InventoryName:        inventory.Name,
		InventoryDescription: inventory.Description,
		ArchiveName:          inventory.Archive.Name,
		ArchiveDescription:   inventory.Description,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    inv,
	})
}

func PutInventory(c *gin.Context) {
	id := c.Param("id")

	data := config.DB.First(&models.Inventory{}, "id = ?", id)

	if data.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Data Not Found",
			"message": "Data Not Found",
		})
		return
	}

	reqInv := models.RequestInventory{}

	c.BindJSON(&reqInv)

	inv := models.Inventory{
		Name:        reqInv.InventoryName,
		Description: reqInv.InventoryDescription,
	}

	config.DB.Model(&models.Inventory{}).Where("id = ?", id).Updates(&inv)

	archive := models.Archive{
		Name:        reqInv.ArchiveName,
		Description: reqInv.InventoryDescription,
	}

	config.DB.Model(&models.Archive{}).Where("inventory_id = ?", id).Updates(&archive)

	c.JSON(http.StatusCreated, gin.H{
		"Message": "Update Successfully",
		"data":    inv,
	})
}

func DeleteInventory(c *gin.Context) {
	id := c.Param("id")

	Inventory := models.Inventory{}

	data := config.DB.First(&Inventory, "id = ?", id)

	if data.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Data Not Found",
			"message": "Data Not Found",
		})
		return
	}

	config.DB.Delete(&Inventory, "id = ?", id)

	c.JSON(http.StatusCreated, gin.H{
		"Message": "Delete successfully",
	})
}
