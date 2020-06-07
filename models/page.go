package models

// Page
type PageCategory struct {
	Model
	Name string `json:"name" gorm:"not null"`
	Slug string `json:"slug" gorm:"unique_index;not null"`
}

// Page
type Page struct {
	Model
	Title      string       `json:"title" gorm:"not null"`
	Body       string       `json:"body" gorm:"not null"`
	Short      string       `json:"short" gorm:"not null"`
	Slug       string       `json:"slug" gorm:"unique_index;not null"`
	UserID     uint         `json:"-" gorm:"not null"`
	User       User         `json:"created_by" gorm:"foreignkey:UserID"`
	CategoryID uint         `json:"-" gorm:"not null"`
	Category   PageCategory `json:"page_category" gorm:"foreignkey:CategoryID"`
}