package db

import "gorm.io/gorm"

type IceCreamScoop struct {
	gorm.Model
	Name      string `gorm:"type:varchar(50);not null"`
	ImagePath string `gorm:"type:varchar(50);not null"`
}

type IceCreamTopping struct {
	gorm.Model
	Name      string `gorm:"type:varchar(50);not null"`
	ImagePath string `gorm:"type:varchar(50);not null"`
}

func (IceCreamScoop) TableName() string {
	return "ice_cream_scoops"
}

func (IceCreamTopping) TableName() string {
	return "ice_cream_toppings"
}
