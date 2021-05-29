// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package wire

import (
	"go-vue-next-server/pkg/infrastructure/store"
	"go-vue-next-server/pkg/interface/handler"
	"go-vue-next-server/pkg/usecase"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitializeUserAccountHandler(db *gorm.DB) handler.UserAccountHandler {
	userAccountRepository := store.NewUserAccountStore(db)
	userAccountUsecase := usecase.NewUserAccountUsecase(userAccountRepository)
	userAccountHandler := handler.NewUserAccountHandler(userAccountUsecase)
	return userAccountHandler
}
