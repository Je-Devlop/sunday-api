package db

import (
	"Je-Devlop/sunday-api/sunday"
	"errors"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	*gorm.DB
}

func NewPostgres(dns string) (*Postgres, error) {

	db, err := gorm.Open(postgres.Open(dns))
	if err != nil {
		panic(err.Error())
	}

	if err := db.AutoMigrate(&IceCreamScoop{}, &IceCreamTopping{}); err != nil {
		return nil, err
	}
	return &Postgres{db}, nil
}

func (db *Postgres) CreateICreamScoop(scoop sunday.Scoop) error {
	timeNow := time.Now()
	t := db.DB.Exec("INSERT INTO ice_cream_scoops(name, image_path, created_at, updated_at, deleted_at) VALUES(?,?,?,?,?)", scoop.Name, scoop.ImagePath, timeNow, timeNow, timeNow)
	if err := t.Error; err != nil {
		return err
	}

	return nil
}

func (db *Postgres) GetAllIceCreamScoops() ([]sunday.Scoop, error) {
	var scoop []sunday.Scoop
	r := db.Table("ice_cream_scoops").Find(&scoop)

	if err := r.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []sunday.Scoop{}, nil
		} else {
			return []sunday.Scoop{}, err
		}
	}

	return scoop, nil
}

func (db *Postgres) CreateICreamTopping(topping sunday.Topping) error {
	timeNow := time.Now()
	t := db.DB.Exec("INSERT INTO ice_cream_toppings(name, image_path, created_at, updated_at, deleted_at) VALUES(?,?,?,?,?)", topping.Name, topping.ImagePath, timeNow, timeNow, timeNow)
	if err := t.Error; err != nil {
		return err
	}

	return nil
}

func (db *Postgres) GetAllIceCreamToppings() ([]sunday.Topping, error) {
	var toppings []sunday.Topping
	r := db.Table("ice_cream_toppings").Find(&toppings)

	if err := r.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []sunday.Topping{}, nil
		} else {
			return []sunday.Topping{}, err
		}
	}

	return toppings, nil
}
