package database

import (
	"go-clean/internal/domain"

	"gorm.io/gorm"
)

type gormTransaction struct {
	tx *gorm.DB
}

func (t *gormTransaction) Commit() error {
	return t.tx.Commit().Error
}

func (t *gormTransaction) Rollback() error {
	return t.tx.Rollback().Error
}

func newGormTransaction(tx *gorm.DB) domain.Transaction {
	return &gormTransaction{tx: tx}
}