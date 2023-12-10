package Api

import (
	"ComedorGo/backend/GeneratedQR"
	"ComedorGo/backend/Models"
	"encoding/json"
	"fmt"
	"net/http"
)

// GenerarYGuardarQR es un ejemplo de cómo podrías integrar la generación de QR a través de la API
func GenerarYGuardarQR(w http.ResponseWriter, r *http.Request) {
	var nuevoEstudiante Models.InformacionEstudiante
	err := json.NewDecoder(r.Body).Decode(&nuevoEstudiante)
	if err != nil {
		http.Error(w, "Error al decodificar la solicitud", http.StatusBadRequest)
		return
	}

	qrCodeString, err := GeneratedQR.GenerarQR(nuevoEstudiante)
	if err != nil {
		http.Error(w, "Error al generar el QR", http.StatusInternalServerError)
		return
	}
	print(qrCodeString)

	// Puedes guardar el código QR en la base de datos u realizar otras acciones según tus necesidades

	// Configurar la respuesta HTTP con el encabezado y el cuerpo JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Código QR generado y guardado con éxito"})
}
func DesencriptarQR(w http.ResponseWriter, r *http.Request) {
	// Leer la información de InfoQR desde el cuerpo de la solicitud
	var infoQRRequest struct {
		Encrypted               string `json:"encrypted"`
		FKInformacionEstudiante uint   `json:"codigoEstudiante"`
	}

	err := json.NewDecoder(r.Body).Decode(&infoQRRequest)
	if err != nil {
		http.Error(w, `{"error": "Error al decodificar la solicitud"}`, http.StatusBadRequest)
		return
	}

	// Imprimir detalles de la solicitud en el servidor
	fmt.Printf("Solicitud recibida: %+v\n", infoQRRequest)

	// Desencriptar la información utilizando los métodos en qr.go
	desencriptado, err := GeneratedQR.Decrypta(infoQRRequest.Encrypted)
	if err != nil {
		fmt.Printf("Error al desencriptar la información: %s\n", err)
		http.Error(w, `{"error": "Error al desencriptar la información"}`, http.StatusInternalServerError)
		return
	}

	// Imprimir información desencriptada en el servidor
	fmt.Printf("Información desencriptada: %s\n", desencriptado)

	// Configurar la respuesta HTTP con el encabezado y el cuerpo JSON
	w.Header().Set("Content-Type", "application/json")

	// Verificar si la desencriptación fue exitosa
	if desencriptado == "" {
		http.Error(w, `{"error": "La desencriptación no produjo resultados válidos"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"decryptedInfo": desencriptado})
}
