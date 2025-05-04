package main

import (
    "mini-task-backend/routes"
    "mini-task-backend/services"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

func main() {
    // Inicializar conexión Mongo
    services.InitMongo()

    r := gin.Default()

    // Configurar CORS para permitir solicitudes desde cualquier origen
    r.Use(cors.Default())

    // Registrar rutas
    routes.RegisterRoutes(r)

    r.Run(":8080")
}