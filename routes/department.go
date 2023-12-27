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

func GetDepartment(c *gin.Context) {
	departments := []models.Department{}

	//:: WITHOUT RELATION
	// config.DB.Find(&departments)

	//:: WITH RELAION
	config.DB.Preload("Positions").Find(&departments)

	getDepartmentResponse := []models.GetDepartmentResponse{}

	for _, d := range departments {
		positions := []models.PositionResponse{}

		for _, p := range d.Positions {
			pos := models.PositionResponse{
				ID:           p.ID,
				Name:         p.Name,
				Code:         p.Code,
				DepartmentID: d.ID,
			}

			positions = append(positions, pos)
		}

		dept := models.GetDepartmentResponse{
			ID:        d.ID,
			Name:      d.Name,
			Code:      d.Code,
			Positions: positions,
		}

		getDepartmentResponse = append(getDepartmentResponse, dept)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome Department",
		"data":    getDepartmentResponse,
	})
}

func GetDepartmentById(c *gin.Context) {
	id := c.Param("id")

	department := models.Department{}
	//:: WITHOUT JOIN
	// data := config.DB.First(&department, "id = ?", id)

	//:: WITH JOIN
	data := config.DB.Preload("Positions").First(&department, "id = ?", id)

	//:: VALIDATE DATA
	if data.Error != nil {
		log.Printf("Error in GetDepartment Id : %s ", data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Department Not Found",
		})
		return
	}

	positions := []models.PositionResponse{}

	for _, p := range department.Positions {
		pos := models.PositionResponse{
			ID:           p.ID,
			Name:         p.Name,
			Code:         p.Code,
			DepartmentID: department.ID,
		}

		positions = append(positions, pos)
	}

	getDepartmentResponse := models.GetDepartmentResponse{
		ID:        department.ID,
		Name:      department.Name,
		Code:      department.Code,
		Positions: positions,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    getDepartmentResponse,
	})
}

func PostDepartment(c *gin.Context) {
	reqDep := models.Department{}
	c.BindJSON(&reqDep)

	config.DB.Create(&reqDep)

	getDepartmenResponse := models.GetDepartmentResponse{
		ID:        reqDep.ID,
		Name:      reqDep.Name,
		Code:      reqDep.Code,
		CreatedAt: reqDep.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: reqDep.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Insert Successfully",
		"data":    getDepartmenResponse,
	})
}

func PutDepartment(c *gin.Context) {
	id := c.Param("id")
	reqDep := models.Department{}
	c.BindJSON(&reqDep)
	currentTime := time.Now()

	//: CONVERT STRING TO UINT
	uint64Value, err := strconv.ParseUint(id, 10, 0)

	if err != nil {
		// Convertion not successful
		log.Printf("Error convert str to int in putDepartment : %s ", err)
		return
	}

	//:: IDENTIFICATION UINT AS UINT64
	idVal := uint(uint64Value)

	config.DB.Model(&models.Department{}).Where("id = ?", id).Updates(reqDep)

	getDepartmentResponse := models.GetDepartmentResponse{
		ID:        idVal,
		Name:      reqDep.Name,
		Code:      reqDep.Code,
		UpdatedAt: currentTime.Format("2006-01-02 15:04:05"),
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Update Successfully",
		"data":    getDepartmentResponse,
	})
}

func DeleteDepartment(c *gin.Context) {
	id := c.Param("id")

	reqDep := models.Department{}

	config.DB.Delete(&reqDep, id)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Delete Successfully",
	})
}
