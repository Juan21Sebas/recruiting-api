package starshipcommsresolver

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"RECRUITING-API/internal/pkg/entity/model"
	"RECRUITING-API/internal/pkg/ports"
)

type candidatesHandler struct {
	Service    ports.CommunicationServices
	Repository ports.DBRepository
}

func newHandler(service ports.CommunicationServices, repository ports.DBRepository) *candidatesHandler {
	return &candidatesHandler{
		Service:    service,
		Repository: repository,
	}
}

func (o *candidatesHandler) postCanditates() gin.HandlerFunc {
	return func(c *gin.Context) {
		var Candidate model.CreateCandidate
		if err := c.BindJSON(&Candidate); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Request"})
			return
		}
		entityResponse, err := o.Service.CreateCandidate(c, &Candidate)
		if err != nil {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.Set("entityResponse", *entityResponse)
		c.JSON(http.StatusOK, entityResponse)
	}
}

// func (o *candidatesHandler) getUsers() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var User model.GetUser
// 		User.Id = c.Param("id")
// 		if err := c.ShouldBindQuery(&User); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Request "})
// 			return
// 		}
// 		entityResponse, err := o.Service.SelectUser(c, &User)
// 		if err != nil {
// 			c.JSON(http.StatusNotFound, err.Error())
// 			return
// 		}

// 		c.Set("entityResponse", *entityResponse)
// 		c.JSON(http.StatusOK, entityResponse)
// 	}
// }

// func (o *candidatesHandler) putUsers() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var User model.UpdateUser
// 		User.Id = c.Param("id")
// 		if err := c.ShouldBindQuery(&User); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Request "})
// 			return
// 		}
// 		if err := c.BindJSON(&User); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Request "})
// 			return
// 		}
// 		entityResponse, err := o.Service.UpdateUser(c, &User)
// 		if err != nil {
// 			c.JSON(http.StatusNotFound, err.Error())
// 			return
// 		}

// 		c.Set("entityResponse", *entityResponse)
// 		c.JSON(http.StatusOK, entityResponse)
// 	}
// }

// func (o *candidatesHandler) deleteUsers() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var User model.DeleteUser
// 		User.Id = c.Param("id")
// 		if err := c.ShouldBindQuery(&User); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Request "})
// 			return
// 		}
// 		entityResponse, err := o.Service.DeleteUser(c, &User)
// 		if err != nil {
// 			c.JSON(http.StatusNotFound, err.Error())
// 			return
// 		}

// 		c.Set("entityResponse", *entityResponse)
// 		c.JSON(http.StatusOK, entityResponse)
// 	}
// }
