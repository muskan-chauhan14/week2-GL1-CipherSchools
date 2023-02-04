package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/muskan-chauhan14/week2-GL1-CipherSchools/database"
	"github.com/muskan-chauhan14/week2-GL1-CipherSchools/handler"
)

func Router() *gin.Engine {
	router := gin.Default() //get the defult engine for further customization
	api := handler.Handler{
		DB: database.GetDB(), //set the handler DB
	}
	router.GET("/books", api.GetBooks) //set the function for http://localhost:8080/books : Get request
	//while calling handler function, gin will pass the *gin.Context in the handler function

	router.POST("/books", api.SaveBook)

	return router
}
