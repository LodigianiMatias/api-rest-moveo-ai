package main

import (
	"log"

	"github.com/api-rest-moveo-ai/database"
	"github.com/api-rest-moveo-ai/handlers"

	_ "github.com/api-rest-moveo-ai/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/api-rest-moveo-ai/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

//@title MOVEO.AI CHALLENGE
//@description Este es un proyecto de prueba técnica. El challenge consiste en realizar un API REST con sus respectivos endpoints usando Go.
//@contact.name Matias Leonel Lodigiani
//@contact.email lodigiani.matias97@gmail.com
//@host localhost:8080

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	router := gin.Default()

	// Conectar a la base de datos
	database.ConnectDatabase()

	// Migrar el modelo Task a la base de datos
	database.DB.AutoMigrate(&models.Task{})

	// Rutas CRUD para Task
	router.POST("/tasks", handlers.CreateTask)
	router.GET("/tasks", handlers.GetAllTasks)
	router.GET("/tasks/:id", handlers.GetTask)
	router.PUT("/tasks/:id", handlers.UpdateTask)
	router.DELETE("/tasks/:id", handlers.DeleteTask)

	// Documentación Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Levantar servidor en el puerto 8080
	router.Run(":8080")
}
