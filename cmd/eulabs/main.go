package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/waldrey/eulabs/configs"
	_ "github.com/waldrey/eulabs/docs"
	"github.com/waldrey/eulabs/internal/handlers"
	"github.com/waldrey/eulabs/internal/infra/database"
	"github.com/waldrey/eulabs/internal/infra/service"
	_ "github.com/waldrey/eulabs/pkg/logger"
)

// @title           Eulabs Products API
// @version         1.0
// @description     API de Produtos

// @host      localhost:8080
// @BasePath  /api/v1
func main() {
	config, err := configs.LoadConfig()
	if err != nil {
		log.Printf("failed load config: %v\n", err)
		return
	}
	db := configs.ConnectDatabase()

	e := echo.New()
	e.Use(middleware.Recover())

	e.GET("/docs/*", echoSwagger.WrapHandler)
	api := e.Group("api/v1/")

	// Handler Product
	productRepository := database.ProductRepository(db)
	productService := service.ProductService(productRepository)
	productHandler := handlers.NewProductHandler(productService)

	productRoutes := api.Group("products")
	productRoutes.POST("/", productHandler.Create)
	productRoutes.GET("", productHandler.List)
	productRoutes.GET("/:id", productHandler.FindOne)
	productRoutes.DELETE("/:id", productHandler.Delete)
	productRoutes.PUT("/:id", productHandler.UpdatePut)
	productRoutes.PATCH("/:id", productHandler.UpdatePatch)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		if err := e.Start(fmt.Sprintf(":%s", config.WebServerPort)); err != nil && err != http.ErrServerClosed {
			log.Fatalf("shutting down server: %v", err)
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		log.Fatalf("could not gracefully shutdown server: %v", err)
	}

	log.Print("server stopped")
}
