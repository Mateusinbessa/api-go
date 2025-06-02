package models

type User struct {
	ID          int    `json:"id"`
	Login       string `json:"login" binding:"required"`
	Name        string `json:"name" binding:"required" gorm:"column:nome"`
	Email       string `json:"email" binding:"required"`
	Password    string `json:"-" gorm:"column:senha" binding:"required,min=6"`
	FullName    string `json:"nome_completo" gorm:"column:nome_completo"`
	Active      bool   `json:"active" gorm:"column:ativo"`
	PhoneNumber string `json:"telefone" gorm:"column:telefone"`
}

type CreateUserRequest struct {
	ID          int    `json:"id"`
	Login       string `json:"login" binding:"required"`
	Name        string `json:"name" binding:"required" gorm:"column:nome"`
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" gorm:"column:senha" binding:"required,min=6"`
	FullName    string `json:"nome_completo" gorm:"column:nome_completo"`
	Active      bool   `json:"active" gorm:"column:ativo"`
	PhoneNumber string `json:"telefone" gorm:"column:telefone"`
}

type UpdateUserRequest struct {
	Login    string `json:"login"`
	Password string `json:"password" gorm:"column:senha"`
}

type UserPagination struct {
	Data       []User
	Total      int64
	Page       int
	Limit      int
	TotalPages int `json:"total_pages"`
}

func (User) TableName() string {
	return "seguranca.usuarios"
}
