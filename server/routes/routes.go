package routes

import (
	"fmt"

	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/rachadiannovansyah/go-echo-boilerplate-ipj/services/token"

	s "github.com/rachadiannovansyah/go-echo-boilerplate-ipj/server"
	"github.com/rachadiannovansyah/go-echo-boilerplate-ipj/server/handlers"

	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	sentryTransaction "github.com/rachadiannovansyah/go-echo-boilerplate-ipj/middle"
)

func ConfigureRoutes(server *s.Server) {
	postHandler := handlers.NewPostHandlers(server)
	categoryHandler := handlers.NewCategoryHandlers(server)
	authHandler := handlers.NewAuthHandler(server)
	registerHandler := handlers.NewRegisterHandler(server)

	sentryMiddleware := sentryTransaction.InitMiddleware()
	server.Echo.Use(sentryMiddleware.CORS)
	server.Echo.Use(sentryMiddleware.SENTRY)
	server.Echo.Use(middleware.Logger())
	server.Echo.Use(sentryecho.New(sentryecho.Options{
		Repanic: true,
	}))

	server.Echo.GET("/swagger/*", echoSwagger.WrapHandler)

	server.Echo.POST("/login", authHandler.Login)
	server.Echo.POST("/register", registerHandler.Register)
	server.Echo.POST("/refresh", authHandler.RefreshToken)

	fmt.Println(server.Config.Auth.AccessSecret)

	r := server.Echo.Group("")
	config := middleware.JWTConfig{
		Claims:     &token.JwtCustomClaims{},
		SigningKey: []byte(server.Config.Auth.AccessSecret),
	}
	r.Use(middleware.JWTWithConfig(config))

	r.GET("/posts", postHandler.GetPosts)
	r.POST("/posts", postHandler.CreatePost)
	r.DELETE("/posts/:id", postHandler.DeletePost)
	r.PUT("/posts/:id", postHandler.UpdatePost)

	r.GET("/categories", categoryHandler.GetCategories)
	r.POST("/categories", categoryHandler.CreateCategory)
	r.PUT("/categories/:id", categoryHandler.UpdateCategory)
	r.DELETE("/categories/:id", categoryHandler.DeleteCategory)
	r.GET("/categories/:id", categoryHandler.GetCategory)
}
