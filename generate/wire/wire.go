// +build wireinject

package wire

import (
	"github.com/birnamwood/go-server-base/pkg/infrastructure/store"
	"github.com/birnamwood/go-server-base/pkg/interface/handler"
	"github.com/birnamwood/go-server-base/pkg/usecase"

	"github.com/google/wire"
	"gorm.io/gorm"
)

//InitializeUserAccountHandler userAccountの依存関係登録
func InitializeUserAccountHandler(db *gorm.DB) handler.UserAccountHandler {
	wire.Build(
		handler.NewUserAccountHandler,
		usecase.NewUserAccountUsecase,
		store.NewUserAccountStore,
	)
	return nil
}
