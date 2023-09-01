package models

type User struct {
	ID       int    `json:"id"`
	Document string `json:"document"`
	Email    string `json:"email"`
	FName    string `json:"f_name"`
	LName    string `json:"l_name"`
	Password string `json:"password"`
	FkRoleId int    `json:"rol"`
	RoleName string `json:"role_name"`
}
