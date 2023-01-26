// Create package models, import "time",
package models

import "time"

// Product struct, ProductResponse struct, ProductUserResponse  struct,
type Product struct {
	ID         int                  `json:"id" gorm:"primary_key:auto_increment"`
	Name       string               `json:"name" gorm:"type: varchar(255)" form:"name"`
	Desc       string               `json:"desc" gorm:"type: text" form:"desc"`
	Price      int                  `json:"price" gorm:"type: int" form:"price"`
	Image      string               `json:"image" gorm:"type: varchar(255)" form:"image"`
	Qty        int                  `json:"qty" gorm:"type: int" form:"qty"`
	UserID     int                  `json:"user_id" gorm:"type: int" form:"user_id"`
	User       UsersProfileResponse `json:"user"`
	Category   []Category           `json:"category" gorm:"many2many:product_categories"`
	CategoryID []int                `json:"category_id" gorm:"-" form:"category_id"`
	CreatedAt  time.Time            `json:"-"`
	UpdatedAt  time.Time            `json:"-"`
}

type ProductResponse struct {
	ID         int                  `json:"id"`
	Name       string               `json:"name"`
	Desc       string               `json:"desc"`
	Price      int                  `json:"price"`
	Image      string               `json:"image"`
	Qty        int                  `json:"qty"`
	UserID     int                  `json:"user_id"`
	User       UsersProfileResponse `json:"user"`
	Category   []Category           `json:"category" gorm:"many2many:product_categories"`
	CategoryID []int                `json:"category_id" gorm:"-" form:"category_id"`
}

type ProductUserResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	Price  int    `json:"price"`
	Image  string `json:"image"`
	Qty    int    `json:"qty"`
	UserID int    `json:"-"`
}

// ProductResponse TableName method and ProductUserResponse TableName method here ...
func (ProductResponse) TableName() string {
	return "products"
}

func (ProductUserResponse) TableName() string {
	return "products"
}
