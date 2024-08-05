package models

type SubDepartment struct {
	ID           uint32     `gorm:"primaryKey" json:"id"`
	Name         string     `json:"name"`
	DepartmentID uint32     `json:"department_id"`
	Employees    []Employee `gorm:"foreignKey:SubDepartmentID" json:"employees"`
}
