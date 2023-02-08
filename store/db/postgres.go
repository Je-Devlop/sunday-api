package db

import (
	"Je-Devlop/sunday-api/sunday"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	db *gorm.DB
}

func NewPostgres(dns string) (*Postgres, error) {

	db, err := gorm.Open(postgres.Open(dns))
	if err != nil {
		panic(err.Error())
	}

	if err := db.AutoMigrate(&sunday.Scoop{}); err != nil {
		return nil, err
	}
	return &Postgres{db}, nil
}

func (s *Postgres) New(scoop *sunday.Scoop) error {
	return s.db.Create(scoop).Error
}
