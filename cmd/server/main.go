package main

import (
	"github.com/eron97/bff-golang.git/cmd/config"
	"github.com/eron97/bff-golang.git/src/controller"
	"github.com/eron97/bff-golang.git/src/controller/routes"
	"github.com/eron97/bff-golang.git/src/model/persistence"
	"github.com/eron97/bff-golang.git/src/model/service"
	"github.com/eron97/bff-golang.git/src/model/service/crypto"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func main() {

	_ = config.NewConfig()

	db, err := config.NewDatabaseConnection()
	if err != nil {
		zap.L().Fatal("Failed to connect to database", zap.Error(err))
	}

	userController := initDependencies(db)

	app := fiber.New()

	routes.SetupRoutes(app, userController)

	if err := app.Listen(":8080"); err != nil {
		zap.L().Fatal("Failed to start server", zap.Error(err))
	}

}

func initDependencies(database *gorm.DB) controller.ControllerInterface {
	cryptoService := &crypto.Crypto{}
	persistence := persistence.NewDBConnection(database)
	service := service.NewServiceInstance(cryptoService, persistence)
	return controller.NewControllerInstance(service)
}
