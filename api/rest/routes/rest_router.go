// api/rest/router/rest_router.go
package router

import (
	"fiber-tutorial/api/rest/handler"
	"fiber-tutorial/internal/datastore"

	"github.com/gofiber/fiber/v2"
)

// NewRouter, yeni bir HTTP yönlendirici oluşturur.
func InitRoutes(a *fiber.App, db *datastore.Mysqldb) {

	restHandler := handler.NewRESTHandler(db)

	// REST API rotaları
	root := a.Group("/api")
	root.Post("user", restHandler.CreateUser)
	root.Get("user", restHandler.GetUser)

	// Middleware'ler ekle

}
