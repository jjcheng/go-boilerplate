package repository

import (
	"github.com/jjcheng/go-boilerplate/internal/exception"

	"gorm.io/gorm"
)

type UnitOfWork interface {
	BeginTransaction() UnitOfWork
	Rollback()
	CommitTransaction() *exception.Exception
	DB() *gorm.DB
	// account
	AccountUserRepository() AccountUserRepository
}
