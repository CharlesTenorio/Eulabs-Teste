package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/eulabs/back-end/api-product/internal/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewHTTPServer(e *echo.Echo, conf *config.Config) *http.Server {
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	srv := &http.Server{
		Addr:         ":" + conf.PORT,
		Handler:      e,
		ReadTimeout:  10 * time.Second, // Wait for 10 seconds for a request to be fully read
		WriteTimeout: 10 * time.Second, // Respond within 10 seconds
		ErrorLog:     log.New(os.Stderr, "logger: ", log.Lshortfile),
	}

	return srv
}
