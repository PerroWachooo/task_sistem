package routes

import (
    "github.com/gin-gonic/gin"
    "mini-task-backend/controllers"
)

func RegisterRoutes(router *gin.Engine) {
    router.GET("/ping", controllers.PingController)
    router.POST("/tasks", controllers.CreateTaskController) // Create a new task
    router.GET("/tasks", controllers.GetTasksController) // Get all tasks
    router.PUT("/tasks/:id", controllers.UpdateTaskController) // Update a task by ID
    router.DELETE("/tasks/:id", controllers.DeleteTaskController) // Delete a task by ID
    
}