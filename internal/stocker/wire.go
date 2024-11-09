//go:build wireinject
// +build wireinject

package stocker

import (
	"github.com/google/wire"
	"h11/backend/internal/stocker/application/service"
	"h11/backend/internal/stocker/domain/repository"
	"h11/backend/internal/stocker/infrastructure/database"
	"h11/backend/internal/stocker/infrastructure/repository/implements"
	"h11/backend/internal/stocker/presentation/controller"
)

// データベース
var databaseSet = wire.NewSet(
	database.GetMySQLConnection,
)

// リポジトリ
var repositorySet = wire.NewSet(
	implements.NewItemRepositoryImpl,
	wire.Bind(new(repository.ItemRepository), new(implements.ItemRepositoryImpl)),
)

// サービス
var serviceSet = wire.NewSet(
	service.NewItemService,
)

// コントローラー
var controllerSet = wire.NewSet(
	controller.NewItemController,
)

// コントローラーセット
type ControllersSet struct {
	ItemController controller.ItemController
}

// コントローラーセット作成
func InitializeController() *ControllersSet {
	wire.Build(
		databaseSet,
		repositorySet,
		serviceSet,
		controllerSet,
		wire.Struct(new(ControllersSet), "*"),
	)
	return nil
}
