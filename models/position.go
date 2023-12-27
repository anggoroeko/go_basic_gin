package models

import "gorm.io/gorm"

type Position struct {
	gorm.Model
	Name         string     `json:"name"`
	Code         string     `json:"code"`
	DepartmentID uint       `json:"department_id"`
	Department   Department `json:"department"`
}

type PositionResponse struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Code         string `json:"code"`
	DepartmentID uint   `json:"department_id"`
}

type GetPositionResponse struct {
	ID           uint                `json:"id"`
	Name         string              `json:"name"`
	Code         string              `json:"code"`
	DepartmentID uint                `json:"department_id"`
	CreatedAt    string              `json:"created_at"`
	UpdatedAt    string              `json:"updated_at"`
	Department   *DepartmentResponse `json:"department"`
}
