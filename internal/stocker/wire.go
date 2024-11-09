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

var databaseSet = wire.NewSet(
	database.GetMySQLConnection,
)

var repositorySet = wire.NewSet(
	implements.NewItemRepositoryImpl,
	wire.Bind(new(repository.ItemRepository), new(implements.ItemRepositoryImpl)),
)

var serviceSet = wire.NewSet(
	service.NewItemService,
)

var controllerSet = wire.NewSet(
	controller.NewItemController,
)

type ControllersSet struct {
	ItemController controller.ItemController
}

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
