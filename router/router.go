package router

import (
	"github.com/gin-gonic/gin"

	"example.com/person_crud/controllers"
)

func Router(r *gin.Engine) *gin.Engine {
	r.GET("/persons", controllers.GetAllPersons)
	r.GET("/persons/:id", controllers.GetPerson)
	r.PUT("/persons/:id", controllers.UpdatePerson)
	r.DELETE("/persons/:id", controllers.DeletePerson)
	r.POST("/persons", controllers.PostPerson)

	return r

}
