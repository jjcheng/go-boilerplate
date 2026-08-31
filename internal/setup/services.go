package setup

import (
	"github.com/jjcheng/go-boilerplate/internal/repository"
	"github.com/jjcheng/go-boilerplate/internal/service"
)

// initializes and returns all application services
func SetupServices(unitOfWork repository.UnitOfWork, logger *service.Logger) *service.Dependencies {
	fileService := service.NewFileService(logger)
	dependencies := service.NewDependencies(unitOfWork, logger, fileService)
	return dependencies
}
