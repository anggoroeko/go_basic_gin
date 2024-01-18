package models

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Name       string `json:"name"`
	Address    string `json:"address"`
	Email      string `json:"email"`
	PositionID uint   `json:"position_id"`
	// Position   PositionResponse `json:"position"`
	Position    Position `json:"position"`
	Inventories []EmployeeInventory
}

type EmployeeResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Email   string `json:"email"`
}

type GetEmployeeResponse struct {
	ID        uint              `json:"id"`
	Name      string            `json:"name"`
	Address   string            `json:"address"`
	Email     string            `json:"email"`
	CreatedAt string            `json:"created_at"`
	UpdatedAt string            `json:"updated_at"`
	Position  *PositionResponse `json:"position"`
}
