package models

type Department struct {
	ID      uint32          `gorm:"primaryKey"`
	Name    string          `json:"name"`
	SubDeps []SubDepartment `gorm:"foreignKey:DepartmentID" json:"sub"`
}
