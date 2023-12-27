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

func GetEmployees(c *gin.Context) {
	employees := []models.Employee{}

	//:: WITHOUT RELATION
	// config.DB.Find(&departments)

	//:: WITH RELAION
	config.DB.Preload("Position").Find(&employees)

	getEmployeeResponse := []models.GetEmployeeResponse{}

	for _, d := range employees {
		pos := models.PositionResponse{
			ID:           d.Position.ID,
			Name:         d.Position.Name,
			Code:         d.Position.Code,
			DepartmentID: d.ID,
		}

		empl := models.GetEmployeeResponse{
			ID:        d.ID,
			Name:      d.Name,
			Address:   d.Address,
			Email:     d.Email,
			Position:  &pos,
			CreatedAt: d.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: d.UpdatedAt.Format("2006-01-02 15:04:05"),
		}

		getEmployeeResponse = append(getEmployeeResponse, empl)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome Employee",
		"data":    getEmployeeResponse,
	})
}

func GetEmployeeById(c *gin.Context) {
	id := c.Param("id")

	employee := models.Employee{}
	//:: WITHOUT JOIN
	// data := config.DB.First(&department, "id = ?", id)

	//:: WITH JOIN
	data := config.DB.Preload("Position").First(&employee, "id = ?", id)

	//:: VALIDATE DATA
	if data.Error != nil {
		log.Printf("Error in GetEmployee Id : %s ", data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Employee Not Found",
		})
		return
	}

	positions := models.PositionResponse{
		ID:           employee.Position.ID,
		Name:         employee.Position.Name,
		Code:         employee.Position.Code,
		DepartmentID: employee.Position.DepartmentID,
	}

	getEmployeeResponse := models.GetEmployeeResponse{
		ID:       employee.ID,
		Name:     employee.Name,
		Address:  employee.Address,
		Email:    employee.Email,
		Position: &positions,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    getEmployeeResponse,
	})
}

func PostEmployees(c *gin.Context) {
	reqEmpl := models.Employee{}
	c.BindJSON(&reqEmpl)

	config.DB.Create(&reqEmpl)

	getEmployeeResponse := models.GetEmployeeResponse{
		ID:        reqEmpl.ID,
		Name:      reqEmpl.Name,
		Address:   reqEmpl.Address,
		Email:     reqEmpl.Email,
		CreatedAt: reqEmpl.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: reqEmpl.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Insert Successfully",
		"data":    getEmployeeResponse,
	})
}

func PutEmployees(c *gin.Context) {
	id := c.Param("id")
	reqEmpl := models.Employee{}
	c.BindJSON(&reqEmpl)
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

	config.DB.Model(&models.Employee{}).Where("id = ?", id).Updates(reqEmpl)

	getEmployeeResponse := models.GetEmployeeResponse{
		ID:        idVal,
		Name:      reqEmpl.Name,
		Address:   reqEmpl.Address,
		Email:     reqEmpl.Email,
		UpdatedAt: currentTime.Format("2006-01-02 15:04:05"),
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Update Successfully",
		"data":    getEmployeeResponse,
	})
}

func DeleteEmployees(c *gin.Context) {
	id := c.Param("id")

	reqEmpl := models.Employee{}

	config.DB.Delete(&reqEmpl, id)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Delete Successfully",
	})
}
