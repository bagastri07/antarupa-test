package application

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bagastri07/antarupa-test/controller"
	"github.com/bagastri07/antarupa-test/database"
	"github.com/bagastri07/antarupa-test/model"
	"github.com/bagastri07/antarupa-test/repository"
	"github.com/bagastri07/antarupa-test/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type App struct {
	E  *echo.Echo
	db *gorm.DB
}

func NewApp() *App {
	app := &App{
		E:  echo.New(),
		db: database.InitDatabase(),
	}

	app.initMiddleware()
	app.initValidator()
	app.initRoutes()

	return app
}

func (app *App) initRoutes() {
	// init repositories
	userRepo := repository.NewUserRepo(app.db)
	shopRepo := repository.NewShopRepo(app.db)
	userCurrencyRepo := repository.NewUserCurrencyRepo(app.db)
	userItemRepo := repository.NewUserItemRepo(app.db)

	// init controller
	itemController := controller.NewItemController(userRepo, shopRepo, userCurrencyRepo, userItemRepo)

	// route
	app.E.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, model.MessageResponse{
			Message: "hello world",
		})
	})

	item := app.E.Group("items")
	item.POST("", itemController.PurchaseItem)
}

func (app *App) initValidator() {
	validator.Init(app.E)
}

func (app *App) initMiddleware() {
	app.E.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
	}))
	app.E.Use(middleware.Logger())
}

func (app *App) Start() {
	app.E.HideBanner = true

	// Start server
	go func() {
		if err := app.E.Start(":7000"); err != nil {
			app.E.Logger.Info("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	// Graceful Shutdown see: https://echo.labstack.com/cookbook/graceful-shutdown
	// Make sure no more in-flight request within 10seconds timeout
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	log.Info("Shutting down the server gracefully")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := app.E.Shutdown(ctx); err != nil {
		app.E.Logger.Fatal(err)
	}
}
