//go:build wireinject
// +build wireinject

package stocker

import (
	"h11/backend/internal/stocker/application/usecase"
	"h11/backend/internal/stocker/domain/repository"
	"h11/backend/internal/stocker/domain/service"
	"h11/backend/internal/stocker/infrastructure/database"
	"h11/backend/internal/stocker/infrastructure/repository/implements"
	"h11/backend/internal/stocker/presentation/controller"

	"github.com/google/wire"
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
	implements.NewStockInRepositoryImpl,
	wire.Bind(new(repository.StockInRepository), new(implements.StockInRepositoryImpl)),
	implements.NewStockOutRepositoryImpl,
	wire.Bind(new(repository.StockOutRepository), new(implements.StockOutRepositoryImpl)),
)

var serviceSet = wire.NewSet(
	service.NewItemStockDomainService,
	service.NewStockOutDomainService,
	service.NewStockInDomainService,
)

// サービス
var usecaseSet = wire.NewSet(
	usecase.NewJancodeUsecase,
	usecase.NewItemUsecase,
	usecase.NewItemStockUsecase,
	usecase.NewUserUsecase,
	usecase.NewAuthorizationUsecase,
	usecase.NewStoreUsecase,
	usecase.NewStoreAuthorizationUsecase,
	usecase.NewStockInUsecase,
	usecase.NewStockOutUsecase,
)

// コントローラー
var controllerSet = wire.NewSet(
	controller.NewJancodeController,
	controller.NewItemController,
	controller.NewItemStockController,
	controller.NewUserController,
	controller.NewStoreController,
	controller.NewStockInController,
	controller.NewStockOutController,
)

// コントローラーセット
type ControllersSet struct {
	JancodeController   controller.JancodeController
	ItemController      controller.ItemController
	ItemStockController controller.ItemStockController
	UserController      controller.UserController
	StoreController     controller.StoreController
	StockInController   controller.StockInController
	StockOutController  controller.StockOutController
}

// コントローラーセット作成
func InitializeController() *ControllersSet {
	wire.Build(
		databaseSet,
		repositorySet,
		serviceSet,
		usecaseSet,
		controllerSet,
		wire.Struct(new(ControllersSet), "*"),
	)
	return nil
}
