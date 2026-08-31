package service

import (
	"github.com/jjcheng/go-boilerplate/internal/repository"
)

type Dependencies struct {
	UnitOfWork repository.UnitOfWork
	Logger     *Logger
	File       *File
}

func NewDependencies(unitOfWork repository.UnitOfWork, logger *Logger, file *File) *Dependencies {
	return &Dependencies{
		UnitOfWork: unitOfWork,
		Logger:     logger,
		File:       file,
	}
}
