package store

import (
	"Je-Devlop/sunday-api/sunday"

	"gorm.io/gorm"
)

type GormStore struct {
	db *gorm.DB
}

func NewGormStore(db *gorm.DB) *GormStore {
	return &GormStore{db: db}
}

func (s *GormStore) New(scoop *sunday.Scoop) error {
	return s.db.Create(scoop).Error
}
