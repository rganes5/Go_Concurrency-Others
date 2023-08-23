package models

import "gorm.io/gorm"

type Post struct {
	Body   string `json:"body"`
	UserID int    `json:"userId"`
}

type Req struct {
	Amount float64 `json:"amount"`
}

type Res struct {
	Discount float64 `json:"discount"`
	Suc      string  `json:"suc"`
}

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Password string `json:"password"`
}
