package models

type Author struct {
	Model
	Name    string
	Website *string
	Books   []Book `gorm:"foreignKey:AuthorId"`
}
