package models

type User struct {
	IDUser       uint32 `gorm: "primaryKey"`
	Phone        string `json: "phone"`
	Password     string `json: "password"`
	ConfirmPass  string `json: "confirmpass"`
	IsEmployee   bool   `json: "isemployee"`
	Role         string `json: "role"`
	AccessToken  string `json: "accessToken`
	RefreshToken string `json: "refreshToken"`
}
