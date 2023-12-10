package Api

import (
	"ComedorGo/backend/Estudiantes"
	"ComedorGo/backend/Inscripcion"
	"ComedorGo/backend/Logins"
	"ComedorGo/backend/Models"
	"ComedorGo/backend/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// EstudianteRoutes define las rutas de la API
func EstudianteRoutes(r *gin.Engine) {
	// Configurar encabezados CORS
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
	})

	// Grupo de rutas para estudiantes
	estudiantesGroup := r.Group("/estudiantes")
	{
		estudiantesGroup.GET("", ListarEstudiantes)
		estudiantesGroup.GET("/:codigo", ObtenerEstudiante)
		estudiantesGroup.POST("", CrearEstudiante)
	}

	// Ruta para obtener la lista de estudiantes inscritos
	r.GET("/estudiantesinscritos", ObtenerEstudiantesInscritosHandler)

	// Rutas adicionales
	r.POST("/loginestudiante", Logins.LoginHandler)
	r.POST("/api/desencriptarqr", DesencriptarQR)
	r.POST("/inscribircomedor", InscribirComedorHandler)
	r.POST("/inscripcionporcadena", InscripcionPorCadenaHandler)

	// Ruta para generar y guardar QR
	r.POST("/api/generarqr", GenerarYGuardarQR)

	// Ruta para servir archivos estáticos
	r.Static("/static", "./static")
}

// ListarEstudiantes devuelve la lista de estudiantes
func ListarEstudiantes(c *gin.Context) {
	// Obtener la lista de estudiantes desde el paquete Estudiantes
	estudiantes, err := Estudiantes.ListarEstudiantes(db.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener la lista de estudiantes"})
		return
	}

	// Configurar la respuesta HTTP con el cuerpo JSON
	c.JSON(http.StatusOK, estudiantes)
}

// ObtenerEstudiante devuelve un estudiante por su código
func ObtenerEstudiante(c *gin.Context) {
	// Obtener el código del estudiante desde la URL
	codigoEstudianteStr := c.Param("codigo")

	codigoEstudiante, err := strconv.ParseUint(codigoEstudianteStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al convertir el código de estudiante"})
		return
	}

	// Obtener la información del estudiante
	estudiante, err := Estudiantes.GetEstudiantePorCodigo(db.DB, uint(codigoEstudiante))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Estudiante no encontrado"})
		return
	}

	// Configurar la respuesta HTTP con el cuerpo JSON
	c.JSON(http.StatusOK, estudiante)
}

// CrearEstudiante crea un nuevo estudiante
func CrearEstudiante(c *gin.Context) {
	// Leer y decodificar el estudiante desde el cuerpo de la solicitud
	var nuevoEstudiante Models.InformacionEstudiante
	if err := c.ShouldBindJSON(&nuevoEstudiante); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar la solicitud"})
		return
	}

	// Crear estudiante
	err := Estudiantes.CrearEstudiante(db.DB, &nuevoEstudiante)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear estudiante"})
		return
	}

	// Configurar la respuesta HTTP con el cuerpo JSON
	c.JSON(http.StatusOK, gin.H{"message": "Estudiante creado con éxito"})
}

// ObtenerEstudiantesInscritos devuelve la lista de estudiantes inscritos
func ObtenerEstudiantesInscritosHandler(c *gin.Context) {
	// Implementa la lógica para obtener la lista de estudiantes inscritos
	estudiantesInscritos, err := Inscripcion.ObtenerEstudiantesInscritos(db.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener la lista de estudiantes inscritos"})
		return
	}

	// Configura la respuesta HTTP con el cuerpo JSON
	c.JSON(http.StatusOK, estudiantesInscritos)
}

var qrCodeString string

// InscribirComedorHandler maneja la inscripción al comedor
func InscribirComedorHandler(c *gin.Context) {
	// Leer y decodificar la información del estudiante desde el cuerpo de la solicitud
	var nuevoEstudiante Models.InformacionEstudiante
	if err := c.ShouldBindJSON(&nuevoEstudiante); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar la solicitud"})
		return
	}

	// Llamar a la función InscribirAlComedor y manejar el error
	err := Inscripcion.InscribirAlComedor(nuevoEstudiante.CodigoEstudiante, qrCodeString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al inscribir al comedor", "message": err.Error()})
		return
	}

	// Configurar la respuesta HTTP con el cuerpo JSON
	c.JSON(http.StatusOK, gin.H{"message": "Estudiante inscrito en el comedor con éxito"})
}

// InscripcionPorCadenaHandler maneja la inscripción por cadena
func InscripcionPorCadenaHandler(c *gin.Context) {
	// Leer y decodificar la información del estudiante y la cadena desde el cuerpo de la solicitud
	var solicitud struct {
		CodigoEstudiante uint   `json:"codigoEstudiante"`
		CadenaComparar   string `json:"cadenaComparar"`
	}

	if err := c.ShouldBindJSON(&solicitud); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar la solicitud"})
		return
	}

	// Llamar a la función InscripcionPorCadena y manejar el error
	err := Inscripcion.InscripcionPorCadena(solicitud.CodigoEstudiante, solicitud.CadenaComparar)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al realizar la inscripción por cadena", "message": err.Error()})
		return
	}

	// Configurar la respuesta HTTP con el cuerpo JSON
	c.JSON(http.StatusOK, gin.H{"message": "Estudiante inscrito en el comedor con éxito por cadena"})
}

// Resto del código...
