package main

import (
	"auth-service/internal/config"
	"auth-service/internal/handlers"
	"auth-service/internal/middleware"

	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.StaticFile("/", "./frontend/index.html")
	router.StaticFile("/login.html", "./frontend/login.html")
	router.StaticFile("/register.html", "./frontend/register.html")
	router.StaticFile("/403.html", "./frontend/403.html")
	router.Static("/css", "./frontend/css")

	err := router.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		panic(err)
	}

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:5500"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           1 * time.Hour,
	}))

	config.LoadDB()
	config.InitSecret()

	api := router.Group("/api")
	handlers.AuthRoutes(api)
	api.Use(middleware.RequireNoAuth())


	serverConfig := config.LoadServerConfig()


	srv := &http.Server{
		Addr:         ":" + serverConfig.Port,
		Handler:      router,
		ReadTimeout:  serverConfig.ReadTimeout,
		WriteTimeout: serverConfig.WriteTimeout,
		IdleTimeout:  serverConfig.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
