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
	implements.NewItemStockRepositoryImpl,
	wire.Bind(new(repository.ItemStockRepository), new(implements.ItemStockRepositoryImpl)),
	implements.NewUserRepositoryImpl,
	wire.Bind(new(repository.UserRepository), new(implements.UserRepositoryImpl)),
	implements.NewStoreRepositoryImpl,
	wire.Bind(new(repository.StoreRepository), new(implements.StoreRepositoryImpl)),
)

// サービス
var serviceSet = wire.NewSet(
	service.NewItemService,
	service.NewItemStockService,
	service.NewUserService,
	service.NewAuthorizationService,
	service.NewStoreService,
	service.NewStoreAuthorizationService,
)

// コントローラー
var controllerSet = wire.NewSet(
	controller.NewItemController,
	controller.NewItemStockController,
	controller.NewUserController,
	controller.NewStoreController,
)

// コントローラーセット
type ControllersSet struct {
	ItemController      controller.ItemController
	ItemStockController controller.ItemStockController
	UserController controller.UserController
	StoreController controller.StoreController
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
