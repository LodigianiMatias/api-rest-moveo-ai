package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

//Estructura que va a tener la tarea

type Task struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	DueDate     time.Time `json:"due_date"`
}

// Validación del estado antes de guardar

func (t *Task) BeforeSave(tx *gorm.DB) (err error) {
	validStatuses := []string{"pending", "completed", "in-progress"}
	isValid := false

	//Recorre el array de validstatuses y busca coincidencia
	for _, s := range validStatuses {
		if t.Status == s {
			isValid = true
			break
		}
	}

	if !isValid {
		return fmt.Errorf("invalid status: %s, it must be completed, in-progress or pending", t.Status)
	}

	// Si due_date no está establecido, asigna una fecha por defecto
	if t.DueDate.IsZero() {
		t.DueDate = time.Now()
	}

	return nil
}
