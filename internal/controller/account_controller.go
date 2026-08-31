package controller

import (
	"github.com/gin-gonic/gin"
	dto_account "github.com/jjcheng/go-boilerplate/internal/dto/account"
	"github.com/jjcheng/go-boilerplate/internal/feature"
	feature_account_user "github.com/jjcheng/go-boilerplate/internal/feature/account/user"
	"github.com/jjcheng/go-boilerplate/internal/service"
)

func registerAccountController(routerGroup *gin.RouterGroup, dependencies *service.Dependencies, apiGenerator *feature.APIGenerator) {
	registerRoute[*dto_account.User, feature_account_user.Create](routerGroup, dependencies, apiGenerator)
}
