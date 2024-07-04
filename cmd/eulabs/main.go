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
	"github.com/waldrey/eulabs/configs"
	_ "github.com/waldrey/eulabs/pkg/logger"
)

func main() {
	config, err := configs.LoadConfig()
	if err != nil {
		log.Printf("failed load config: %v\n", err)
		return
	}

	configs.ConnectDatabase()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello world")
	})

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
