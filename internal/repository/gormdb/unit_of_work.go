package gormdb

import (
	"github.com/jjcheng/go-boilerplate/internal/exception"
	"github.com/jjcheng/go-boilerplate/internal/repository"
	"github.com/jjcheng/go-boilerplate/internal/service"
	"github.com/jjcheng/go-boilerplate/internal/types"

	"gorm.io/gorm"
)

type UnitOfWork struct {
	db     *gorm.DB
	logger *service.Logger
	// account
	accountUserRepository repository.AccountUserRepository
}

func NewUnitOfWork(db *gorm.DB, logger *service.Logger) repository.UnitOfWork {
	unitOfWork := UnitOfWork{
		db:     db,
		logger: logger,
	}
	// account
	unitOfWork.accountUserRepository = NewAccountUserRepository(db, logger)
	return &unitOfWork
}

func (unitOfWork *UnitOfWork) DB() *gorm.DB {
	return unitOfWork.db
}

// account
func (unitOfWork *UnitOfWork) AccountUserRepository() repository.AccountUserRepository {
	return unitOfWork.accountUserRepository
}

// transaction
func (unitOfWork *UnitOfWork) BeginTransaction() repository.UnitOfWork {
	db := unitOfWork.db.Begin()
	transaction := NewUnitOfWork(db, unitOfWork.logger)
	return transaction
}

func (transaction *UnitOfWork) Rollback() {
	transaction.db.Rollback()
}

func (transaction *UnitOfWork) CommitTransaction() *exception.Exception {
	if err := transaction.db.Commit().Error; err != nil {
		transaction.logger.ErrorFunction(err)
		return exception.NewException(types.ExceptionTypeDatabase)
	}
	return nil
}

//end of transaction
