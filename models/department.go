package models

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	Name      string     `json:"name"`
	Code      string     `json:"code"`
	Positions []Position `json:"positions"`
}

type DepartmentResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type GetDepartmentResponse struct {
	ID        uint               `json:"id"`
	Name      string             `json:"name"`
	Code      string             `json:"code"`
	CreatedAt string             `json:"created_at"`
	UpdatedAt string             `json:"updated_at"`
	Positions []PositionResponse `json:"positions"`
}
