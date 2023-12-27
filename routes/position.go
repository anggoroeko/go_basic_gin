package routes

import (
	"golang_basic_gin_sept_2023/config"
	"golang_basic_gin_sept_2023/models"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetPositions(c *gin.Context) {
	positions := []models.Position{}

	//:: WITHOUT RELATION
	// config.DB.Find(&positions)

	//:: WITH RELATION
	config.DB.Preload("Department").Find(&positions)

	getPositionsResponse := []models.GetPositionResponse{}

	for _, p := range positions {
		department := models.DepartmentResponse{
			ID:   p.Department.ID,
			Name: p.Department.Name,
			Code: p.Department.Code,
		}

		post := models.GetPositionResponse{
			ID:           p.ID,
			Name:         p.Name,
			Code:         p.Code,
			DepartmentID: p.DepartmentID,
			Department:   &department,
			CreatedAt:    p.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:    p.UpdatedAt.Format("2006-01-02 15:04:05"),
		}

		getPositionsResponse = append(getPositionsResponse, post)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome position",
		"data":    getPositionsResponse,
	})
}

func GetPositionsById(c *gin.Context) {
	id := c.Param("id")

	position := models.Position{}

	//:: WITHOUT RELATIONSHIP
	// data := config.DB.First(&position, "id = ?", id)

	//:: WITH RELATIONSHIP
	data := config.DB.Preload("Department").First(&position, "id = ?", id)

	//:: VALIDATE DATA
	if data.Error != nil {
		log.Printf("Error in GetPositionsById : %s", data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Position Not Found",
		})
		return
	}

	dept := models.DepartmentResponse{
		ID:   position.Department.ID,
		Name: position.Department.Name,
		Code: position.Department.Code,
	}

	GetPositionResponse := models.GetPositionResponse{
		ID:           position.ID,
		Name:         position.Name,
		Code:         position.Code,
		DepartmentID: position.DepartmentID,
		Department:   &dept,
		CreatedAt:    position.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    position.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    GetPositionResponse,
	})
}

func PostPositions(c *gin.Context) {
	reqPos := models.Position{}
	c.BindJSON(&reqPos)

	config.DB.Create(&reqPos)

	getPositionResponse := models.GetPositionResponse{
		ID:           reqPos.ID,
		CreatedAt:    reqPos.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    reqPos.UpdatedAt.Format("2006-01-02 15:04:05"),
		Name:         reqPos.Name,
		Code:         reqPos.Code,
		DepartmentID: reqPos.DepartmentID,
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Insert Successfully",
		"data":    getPositionResponse,
	})
}

func PutPositions(c *gin.Context) {
	id := c.Param("id")
	reqPos := models.Position{}
	c.BindJSON(&reqPos)
	currentTime := time.Now()

	config.DB.Model(&models.Position{}).Where("id = ?", id).Updates(reqPos)

	//:: CONVERT STRING TO UINT
	uint64Value, err := strconv.ParseUint(id, 10, 0)

	if err != nil {
		log.Printf("Error convert string to uint : %s", err)
		return
	}

	//:: IDENTIFICATION UINT AS UINT64
	idVal := uint(uint64Value)

	getPositionResponse := models.GetPositionResponse{
		ID:           idVal,
		Name:         reqPos.Name,
		Code:         reqPos.Code,
		UpdatedAt:    currentTime.Format("2006-01-02 15:04:05"),
		DepartmentID: reqPos.DepartmentID,
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Update Successfully",
		"data":    getPositionResponse,
	})
}

func DeletePositions(c *gin.Context) {
	id := c.Param("id")

	reqPos := models.Position{}

	config.DB.Delete(&reqPos, id)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Delete Successfully",
	})
}
