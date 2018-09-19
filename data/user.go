package data

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string   `json:"username"`
	Password string   `json:"password"`
	Coupons  []Coupon `json:"-"`
}
