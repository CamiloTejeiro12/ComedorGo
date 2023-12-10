package Api

import (
	"ComedorGo/backend/GeneratedQR"
	"ComedorGo/backend/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GenerarYGuardarQR es un ejemplo de cómo podrías integrar la generación de QR a través de la API
func GenerarYGuardarQR(c *gin.Context) {
	var nuevoEstudiante Models.InformacionEstudiante
	if err := c.ShouldBindJSON(&nuevoEstudiante); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar la solicitud"})
		return
	}

	qrCodeString, err := GeneratedQR.GenerarQR(nuevoEstudiante)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar el QR"})
		return
	}
	fmt.Println(qrCodeString)

	// Puedes guardar el código QR en la base de datos u realizar otras acciones según tus necesidades

	// Configurar la respuesta HTTP con el encabezado y el cuerpo JSON
	c.JSON(http.StatusOK, gin.H{"message": "Código QR generado y guardado con éxito"})
}

func DesencriptarQR(c *gin.Context) {
	// Leer la información de InfoQR desde el cuerpo de la solicitud
	var infoQRRequest struct {
		Encrypted               string `json:"encrypted"`
		FKInformacionEstudiante uint   `json:"codigoEstudiante"`
	}

	if err := c.ShouldBindJSON(&infoQRRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar la solicitud"})
		return
	}

	// Imprimir detalles de la solicitud en el servidor
	fmt.Printf("Solicitud recibida: %+v\n", infoQRRequest)

	// Desencriptar la información utilizando los métodos en qr.go
	desencriptado, err := GeneratedQR.Decrypta(infoQRRequest.Encrypted)
	if err != nil {
		fmt.Printf("Error al desencriptar la información: %s\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al desencriptar la información"})
		return
	}

	// Imprimir información desencriptada en el servidor
	fmt.Printf("Información desencriptada: %s\n", desencriptado)

	// Configurar la respuesta HTTP con el encabezado y el cuerpo JSON
	c.Header("Content-Type", "application/json")

	// Verificar si la desencriptación fue exitosa
	if desencriptado == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "La desencriptación no produjo resultados válidos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"decryptedInfo": desencriptado})
}
