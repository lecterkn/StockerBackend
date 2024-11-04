// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package stocker

import (
	"github.com/google/wire"
	"h11/backend/internal/stocker/application/service"
	"h11/backend/internal/stocker/infrastructure/database"
	"h11/backend/internal/stocker/infrastructure/repository/implements"
	"h11/backend/internal/stocker/presentation/controller"
)

// Injectors from wire.go:

func InitializeController() *ControllersSet {
	db := database.GetMySQLConnection()
	itemRepositoryImpl := implements.NewItemRepositoryImpl(db)
	itemService := service.NewItemService(itemRepositoryImpl)
	itemController := controller.NewItemController(itemService)
	controllersSet := &ControllersSet{
		ItemController: itemController,
	}
	return controllersSet
}

// wire.go:

var databaseSet = wire.NewSet(database.GetMySQLConnection)

var repositorySet = wire.NewSet(implements.NewItemRepositoryImpl)

var serviceSet = wire.NewSet(service.NewItemService)

var controllerSet = wire.NewSet(controller.NewItemController)

type ControllersSet struct {
	ItemController controller.ItemController
}
