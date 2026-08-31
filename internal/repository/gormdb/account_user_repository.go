package gormdb

import (
	"context"
	"errors"

	dao_account "github.com/jjcheng/go-boilerplate/internal/dao/account"
	"github.com/jjcheng/go-boilerplate/internal/repository"
	"github.com/jjcheng/go-boilerplate/internal/service"

	"gorm.io/gorm"
)

type AccountUserRepository struct {
	db     *gorm.DB
	logger *service.Logger
	repository.Repository[dao_account.User]
}

func NewAccountUserRepository(db *gorm.DB, logger *service.Logger) repository.AccountUserRepository {
	accountUserRepository := AccountUserRepository{
		db:         db,
		logger:     logger,
		Repository: NewRepository[dao_account.User](db, logger),
	}
	return &accountUserRepository
}

func (accountUserRepository *AccountUserRepository) Get(ctx context.Context, id int32) (*dao_account.User, error) {
	var item *dao_account.User
	result := accountUserRepository.db.Model(&dao_account.User{}).Where("id = ?", id).First(&item)
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			accountUserRepository.logger.ErrorFunction(result.Error, id)
		}
		return nil, result.Error
	}
	return item, nil
}

func (accountUserRepository *AccountUserRepository) GetByPhoneNumber(ctx context.Context, phoneNumber string) (*dao_account.User, error) {
	var item *dao_account.User
	result := accountUserRepository.db.WithContext(ctx).Model(&dao_account.User{}).Where("phone_number = ?", phoneNumber).First(&item)
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			accountUserRepository.logger.ErrorFunction(result.Error, phoneNumber)
		}
		return nil, result.Error
	}
	return item, nil
}

func (accountUserRepository *AccountUserRepository) GetByAccessTokenHash(ctx context.Context, accessTokenHash string) (*dao_account.User, error) {
	var item *dao_account.User
	result := accountUserRepository.db.WithContext(ctx).Model(&dao_account.User{}).Where("access_token_hash = ?", accessTokenHash).First(&item)
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			accountUserRepository.logger.ErrorFunction(result.Error, accessTokenHash)
		}
		return nil, result.Error
	}
	return item, nil
}
