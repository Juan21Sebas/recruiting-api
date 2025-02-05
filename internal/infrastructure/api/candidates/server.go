package candidates

import (
	"RECRUITING-API/internal/infrastructure/api/candidates/middleware"
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"

	cors "github.com/itsjamie/gin-cors"
)

func CreateServer(db *sql.DB) *gin.Engine {
	server := gin.Default()

	jwtMiddleware := middleware.NewJWTMiddleware("")

	server.Use(jwtMiddleware.MiddlewareFunc)

	server.Use(cors.Middleware(cors.Config{
		Origins:        "*",
		Methods:        "GET,POST,DELETE,PUT",
		RequestHeaders: "Origin, Authorization, Content-Type, Access-Control-Allow-Origin",
		MaxAge:         50 * time.Second,
	}))

	RegisterRoutes(server, db)

	return server
}

func RunServer(db *sql.DB) {
	server := CreateServer(db)
	server.Run(":8089")
}
