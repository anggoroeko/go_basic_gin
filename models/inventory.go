package models

import "gorm.io/gorm"

type Inventory struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Archive     Archive
	Employee    []EmployeeInventory
}

type GetInventoryResponse struct {
	InventoryID          uint   `json:"id"`
	InventoryName        string `json:"inventory_name"`
	InventoryDescription string `json:"inventory_description"`
	ArchiveName          string `json:"archive_name"`
	ArchiveDescription   string `json:"archive_description"`
}

type RequestInventory struct {
	InventoryName        string `json:"inventory_name"`
	InventoryDescription string `json:"inventory_description"`
	ArchiveName          string `json:"archive_name"`
	ArchiveDescription   string `json:"archive_description"`
}

type ResponseInventoryEmployee struct {
	InventoryName        string                      `json:"inventory_name"`
	InventoryDescription string                      `json:"inventory_description"`
	EmployeeInventory    []ResponseEmployeeInventory `json:"employee_inventory"`
}
