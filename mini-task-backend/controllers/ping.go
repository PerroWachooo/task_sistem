package controllers

import (
	"mini-task-backend/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func PingController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

// CreateTask crea una nueva tarea
func CreateTaskController(c *gin.Context) {
	var task struct {
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Completed   bool      `json:"completed"`
		CreatedAt   time.Time `json:"createdAt"`
	}

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Asignar la fecha de creación si no fue proporcionada
	if task.CreatedAt.IsZero() {
		task.CreatedAt = time.Now() // Establecer la fecha de creación como la fecha actual
	}

	// Acceder a la colección de tareas
	collection := services.Client.Database("taskdb").Collection("task")

	// Insertar la nueva tarea en la base de datos
	_, err := collection.InsertOne(c, bson.M{
		"title":       task.Title,
		"description": task.Description,
		"completed":   task.Completed,
		"createdAt":   task.CreatedAt,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar la tarea"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tarea creada con éxito"})
}

// GetTasksController obtiene todas las tareas
func GetTasksController(c *gin.Context) {
	// Obtener la colección de tareas
	collection := services.Client.Database("taskdb").Collection("task")

	// Obtener todas las tareas
	cursor, err := collection.Find(c, bson.M{}) //
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las tareas"})
		return
	}
	defer cursor.Close(c)

	var tasks []bson.M
	if err := cursor.All(c, &tasks); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al leer las tareas"})
		return
	}

	// Responder con las tareas obtenidas
	c.JSON(http.StatusOK, tasks)
}

// UpdateTaskController actualiza una tarea por ID
func UpdateTaskController(c *gin.Context) {
	id := c.Param("id") // Obtener el ID de la tarea desde los parámetros de la URL

	// Convertir el ID a ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Definir los cambios que queremos hacer (marcar como completada)
	update := bson.M{
		"$set": bson.M{"completed": true},
	}

	// Obtener la colección de tareas
	collection := services.Client.Database("taskdb").Collection("task")

	// Actualizar la tarea en la base de datos
	result := collection.FindOneAndUpdate(c, bson.M{"_id": objectID}, update)

	if result.Err() == mongo.ErrNoDocuments {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tarea no encontrada"})
		return
	}

	if result.Err() != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar la tarea"})
		return
	}

	// Responder con éxito
	c.JSON(http.StatusOK, gin.H{"message": "Tarea actualizada correctamente"})
}
