package model

type User struct {
	ID         uint       `gorm:"primaryKey"`
	CpfCnpj    int64      `gorm:"unique" json:"cpf_cnpj"`
	FullName   string     `gorm:"size:30" json:"full_name"`
	Email      string     `gorm:"size:30; unique " json:"email"`
	CategoryID int        `json:"id_category"`
	Categories Categories `gorm:"foreignKey:CategoryID"`
	Password   string     `gorm:"size:15" json:"password"`
}
