package models

type User struct {
	ID          int    `json:"id"`
	Login       string `json:"login"`
	Name        string `json:"name" gorm:"column:nome"`
	Email       string `json:"email"`
	Password    string `json:"-" gorm:"column:senha"`
	FullName    string `json:"nome_completo" gorm:"column:nome_completo"`
	Active      bool   `json:"active" gorm:"column:ativo"`
	PhoneNumber string `json:"telefone" gorm:"column:telefone"`
}

// TableName overrides the table name used by User to `seguranca.usuarios`.
// GO automatically calls this method when querying the User model.
func (User) TableName() string {
	return "seguranca.usuarios"
}
