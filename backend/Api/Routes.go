package Api

import (
	"ComedorGo/backend/Estudiantes"
	"ComedorGo/backend/db"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	//"strconv"
)

// ApiRoutes define las rutas de la API
func EstudianteRoutes(r *mux.Router) {
	r.HandleFunc("/estudiantes", ListarEstudiantes).Methods("GET")
	r.HandleFunc("/estudiantes/{codigo}", ObtenerEstudiante).Methods("GET")
	r.HandleFunc("/estudiantes", CrearEstudiante).Methods("POST")
}

// ListarEstudiantes devuelve la lista de estudiantes
func ListarEstudiantes(w http.ResponseWriter, r *http.Request) {
	// Obtener la lista de estudiantes desde el paquete Estudiantes
	estudiantes, err := Estudiantes.ListarEstudiantes(db.DB)
	if err != nil {
		http.Error(w, "Error al obtener la lista de estudiantes", http.StatusInternalServerError)
		return
	}

	// Configurar la respuesta HTTP con el encabezado y el cuerpo JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(estudiantes)
}
func ObtenerEstudiante(w http.ResponseWriter, r *http.Request) {

}

func CrearEstudiante(w http.ResponseWriter, r *http.Request) {

}

/*
func GenerarQRYGuardar(w http.ResponseWriter, r *http.Request) {
	// Obtener el código del estudiante desde la URL
	params := mux.Vars(r)
	codigoEstudianteStr := params["codigo"]

	codigoEstudiante, err := strconv.ParseUint(codigoEstudianteStr, 10, 64)
	if err != nil {
		http.Error(w, "Error al convertir el código de estudiante", http.StatusBadRequest)
		return
	}


		// Obtener la información del estudiante
		estudiante, err := Estudiantes.GetEstudiantePorCodigo(db.DB, uint(codigoEstudiante))
		if err != nil {
			http.Error(w, "Estudiante no encontrado", http.StatusNotFound)
			return
		}

		// Generar código QR y obtener la cadena encriptada

		qrCodeString, err := GeneratedQR.GenerarQR(estudiante)
		if err != nil {
			http.Error(w, "Error al generar el QR", http.StatusInternalServerError)
			return
		}
		fmt.Println("QR generado:", qrCodeString)


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Código QR generado y guardado con éxito"})
}
*/
