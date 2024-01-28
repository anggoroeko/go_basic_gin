package models

import (
	"time"

	"gorm.io/gorm"
)

type EmployeeInventory struct {
	gorm.Model
	EmployeeID uint `json:"employee_id"`
	//:: BISA DENGAN CARA MANUAL MENDEFINISIKAN FK DAN PK NYA
	Employee  Employee  `gorm:"foreignKey:EmployeeID; reference:ID"`
	Inventory Inventory `gorm:"foreignKey:InventoryID; reference:ID"`

	//:: BISA DENGAN CARA OTOMATIS MENDEFINISIKAN FK DAN PK NYA
	// Employee  Employee
	// Inventory Inventory

	InventoryID uint   `json:"inventory_id"`
	Description string `json:"description"`
}

type ResponseGetRental struct {
	ID            uint      `json:"id"`
	Description   string    `json:"description"`
	EmployeeName  string    `json:"employee_name"`
	InventoryName string    `json:"inventory_name"`
	CreatedAt     time.Time `json:"created_at"`
}

type ResponseEmployeeInventory struct {
	EmployeeID  uint      `json:"employee_id"`
	InventoryID uint      `json:"inventory_id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type RequestRental struct {
	EmployeeID  uint   `json:"employee_id"`
	InventoryID uint   `json:"inventory_id"`
	Description string `json:"description"`
}

type ResponseRental struct {
	ID          uint      `json:"id"`
	EmployeeID  uint      `json:"employee_id"`
	InventoryID uint      `json:"inventory_id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
