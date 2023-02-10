package db

import (
	"time"

	"gorm.io/gorm"
)

type IceCreamScoop struct {
	gorm.Model
	Name      string `gorm:"type:varchar(50);not null"`
	ImagePath string `gorm:"type:varchar(50);not null"`
	Price     int64  `gorm:"not null"`
}

type IceCreamTopping struct {
	gorm.Model
	Name      string `gorm:"type:varchar(50);not null"`
	ImagePath string `gorm:"type:varchar(50);not null"`
	Price     int64  `gorm:"not null"`
}

type OrderDetail struct {
	gorm.Model
	CustomerName string
	OrderDate    time.Time
	TaxID        string
}

type IceCreamScoopDetail struct {
	Name     string `gorm:"not null"`
	Quantity string `gorm:"not null"`
	Price    string `gorm:"not null"`
	OrderId  string `gorm:"not null"`
}

type IceCreamToppingDetail struct {
	Name     string `gorm:"not null"`
	Quantity string `gorm:"not null"`
	Price    string `gorm:"not null"`
	OrderId  string `gorm:"not null"`
}

func (IceCreamScoop) TableName() string {
	return "ice_cream_scoops"
}

func (IceCreamTopping) TableName() string {
	return "ice_cream_toppings"
}

func (OrderDetail) TableName() string {
	return "ice_cream_order"
}

func (IceCreamScoopDetail) TableName() string {
	return "order_scoop_detail"
}

func (IceCreamToppingDetail) TableName() string {
	return "order_topping_detail"
}
