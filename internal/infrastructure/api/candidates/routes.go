package starshipcommsresolver

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	repository "RECRUITING-API/internal/infrastructure/repository"
	services "RECRUITING-API/internal/pkg/service"
)

func RegisterRoutes(e *gin.Engine, db *sql.DB) {

	Repository := repository.NewBdRepository(db)

	Service := services.NewService(Repository)

	Handler := newHandler(Service, Repository)

	// Registra las rutas candidates
	e.POST("/candidates/", Handler.postCanditates())

	// e.GET("/candidates/:id", Handler.getUsers())
	// e.PUT("/candidates/:id", Handler.putUsers())
	// e.DELETE("/candidates/:id", Handler.deleteUsers())

}
