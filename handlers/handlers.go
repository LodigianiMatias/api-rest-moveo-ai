package handlers

import (
	"net/http"

	"github.com/api-rest-moveo-ai/models"

	"github.com/api-rest-moveo-ai/database"
	"github.com/gin-gonic/gin"
)

// CreateTask godoc
// @Summary Crear
// @Description Añade una nueva tarea a la lista
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param task body models.Task true "Task a crear"
// @Router /tasks [post]
func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

// GetAllTasks godoc
// @Summary Obtener
// @Description Obtiene la lista completa de tareas
// @Tags tasks
// @Router /tasks [get]
func GetAllTasks(c *gin.Context) {
	var tasks []models.Task

	if err := database.DB.Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// GetTask godoc
// @Summary Obtener
// @Description Obtiene la tarea por id
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param id path int true "Task ID"
// @Router /tasks/{id} [get]
func GetTask(c *gin.Context) {
	var task models.Task
	id := c.Param("id")

	if err := database.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

// UpdateTask godoc
// @Summary Actualizar
// @Description Modifica la tarea por id
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param id path int true "Task ID"
// @Param task body models.Task true "Task a actualizar"
// @Success 200 {array} models.Task
// @Router /tasks/{id} [put]
func UpdateTask(c *gin.Context) {
	var task models.Task
	id := c.Param("id")

	if err := database.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

// DeleteTask godoc
// @Summary Eliminar
// @Description Elimina la tarea por id
// @Tags tasks
// @Param id path int true "Task ID"
// @Router /tasks/{id} [delete]
func DeleteTask(c *gin.Context) {
	var task models.Task
	id := c.Param("id")

	if err := database.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	database.DB.Delete(&task)
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
