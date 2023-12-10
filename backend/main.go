package main

import (
	"ComedorGo/backend/Api"
	"ComedorGo/backend/db"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	// Configura la base de datos
	db.SetupDatabase()

	// Crea una instancia de Gin
	r := gin.Default()

	// Configurar encabezados CORS
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
	})

	// Incluye las rutas de la API de estudiantes
	Api.EstudianteRoutes(r)

	// Ejemplo para generar un QR y guardarlo a través de la API
	//r.POST("/api/generarqr", Api.GenerarYGuardarQR)

	// Configura el enrutador y comienza a escuchar
	fmt.Println("Escuchando en http://localhost:8000...")
	r.Run("localhost:8000")
}
