package entity

type Employee struct {
	IDEm            uint32  `gorm:"primaryKey" json:"idem"`
	Name            string  `json:"name"`
	Phone           string  `json:"phone"`
	Password        string  `json:"password"`
	Salary          float64 `json:"salary"`
	SubDepartmentID uint32  `json:"sub_department_id"`
	Role            string  `json:"role"`
	// AccessToken     string `json:"accessToken"`
	// RefreshToken    string `json:"refreshToken"`
}
