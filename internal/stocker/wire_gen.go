// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package stocker

import (
	"github.com/google/wire"
	"h11/backend/internal/stocker/application/usecase"
	"h11/backend/internal/stocker/domain/repository"
	"h11/backend/internal/stocker/domain/service"
	"h11/backend/internal/stocker/infrastructure/database"
	"h11/backend/internal/stocker/infrastructure/repository/implements"
	"h11/backend/internal/stocker/presentation/controller"
)

// Injectors from wire.go:

// コントローラーセット作成
func InitializeController() *ControllersSet {
	jancodeUsecase := usecase.NewJancodeUsecase()
	jancodeController := controller.NewJancodeController(jancodeUsecase)
	db := database.GetMySQLConnection()
	itemRepositoryImpl := implements.NewItemRepositoryImpl(db)
	itemUsecase := usecase.NewItemUsecase(itemRepositoryImpl)
	storeRepositoryImpl := implements.NewStoreRepositoryImpl(db)
	storeAuthorizationUsecase := usecase.NewStoreAuthorizationUsecase(storeRepositoryImpl)
	itemController := controller.NewItemController(itemUsecase, storeAuthorizationUsecase)
	itemStockRepositoryImpl := implements.NewItemStockRepositoryImpl(db)
	itemStockDomainService := service.NewItemStockDomainService(itemRepositoryImpl, itemStockRepositoryImpl)
	itemStockUsecase := usecase.NewItemStockUsecase(itemStockRepositoryImpl, itemRepositoryImpl, itemStockDomainService)
	itemStockController := controller.NewItemStockController(itemStockUsecase, storeAuthorizationUsecase)
	userRepositoryImpl := implements.NewUserRepositoryImpl(db)
	userUsecase := usecase.NewUserUsecase(userRepositoryImpl)
	authorizationUsecase := usecase.NewAuthorizationUsecase(userRepositoryImpl)
	userController := controller.NewUserController(userUsecase, authorizationUsecase)
	storeUsecase := usecase.NewStoreUsecase(userRepositoryImpl, storeRepositoryImpl)
	storeController := controller.NewStoreController(storeUsecase)
	stockInRepositoryImpl := implements.NewStockInRepositoryImpl(db)
	stockInUsecase := usecase.NewStockInUsecase(stockInRepositoryImpl, itemRepositoryImpl)
	stockInController := controller.NewStockInController(stockInUsecase)
	stockOutRepositoryImpl := implements.NewStockOutRepositoryImpl(db)
	stockOutUsecase := usecase.NewStockOutUsecase(itemRepositoryImpl, stockOutRepositoryImpl)
	stockOutController := controller.NewStockOutController(stockOutUsecase)
	controllersSet := &ControllersSet{
		JancodeController:   jancodeController,
		ItemController:      itemController,
		ItemStockController: itemStockController,
		UserController:      userController,
		StoreController:     storeController,
		StockInController:   stockInController,
		StockOutController:  stockOutController,
	}
	return controllersSet
}

// wire.go:

// データベース
var databaseSet = wire.NewSet(database.GetMySQLConnection)

// リポジトリ
var repositorySet = wire.NewSet(implements.NewItemRepositoryImpl, wire.Bind(new(repository.ItemRepository), new(implements.ItemRepositoryImpl)), implements.NewItemStockRepositoryImpl, wire.Bind(new(repository.ItemStockRepository), new(implements.ItemStockRepositoryImpl)), implements.NewUserRepositoryImpl, wire.Bind(new(repository.UserRepository), new(implements.UserRepositoryImpl)), implements.NewStoreRepositoryImpl, wire.Bind(new(repository.StoreRepository), new(implements.StoreRepositoryImpl)), implements.NewStockInRepositoryImpl, wire.Bind(new(repository.StockInRepository), new(implements.StockInRepositoryImpl)), implements.NewStockOutRepositoryImpl, wire.Bind(new(repository.StockOutRepository), new(implements.StockOutRepositoryImpl)))

var serviceSet = wire.NewSet(service.NewItemStockDomainService)

// サービス
var usecaseSet = wire.NewSet(usecase.NewJancodeUsecase, usecase.NewItemUsecase, usecase.NewItemStockUsecase, usecase.NewUserUsecase, usecase.NewAuthorizationUsecase, usecase.NewStoreUsecase, usecase.NewStoreAuthorizationUsecase, usecase.NewStockInUsecase, usecase.NewStockOutUsecase)

// コントローラー
var controllerSet = wire.NewSet(controller.NewJancodeController, controller.NewItemController, controller.NewItemStockController, controller.NewUserController, controller.NewStoreController, controller.NewStockInController, controller.NewStockOutController)

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
