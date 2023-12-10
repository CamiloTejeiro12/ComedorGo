package Api

import (
	"ComedorGo/backend/Estudiantes"
	"ComedorGo/backend/Inscripcion"
	"ComedorGo/backend/Logins"
	"ComedorGo/backend/Models"
	"ComedorGo/backend/db"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	//"strconv"
)

// EstudianteRoutes define las rutas de la API
func EstudianteRoutes(r *mux.Router) {
	r.HandleFunc("/estudiantes", ListarEstudiantes).Methods("GET")
	r.HandleFunc("/estudiantesinscritos", ObtenerEstudiantesInscritosHandler).Methods("GET")
	r.HandleFunc("/estudiantes/{codigo}", ObtenerEstudiante).Methods("GET")
	r.HandleFunc("/estudiantes", CrearEstudiante).Methods("POST")
	r.HandleFunc("/loginestudiante", Logins.LoginHandler).Methods("POST")
	r.HandleFunc("/api/desencriptarqr", DesencriptarQR).Methods("POST")
	r.HandleFunc("/inscribircomedor", InscribirComedorHandler).Methods("POST")
	r.HandleFunc("/inscripcionporcadena", InscripcionPorCadenaHandler).Methods("POST") // Nuevo endpoint

	r.HandleFunc("/api/generarqr)", GenerarYGuardarQR).Methods("POST")
	// Agrega estas líneas al final de EstudianteRoutes en Routes.go
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	//r.HandleFunc("/login", MostrarPaginaLogin).Methods("GET")

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

// ObtenerEstudiante devuelve un estudiante por su código
func ObtenerEstudiante(w http.ResponseWriter, r *http.Request) {
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

	// Configurar la respuesta HTTP con el encabezado y el cuerpo JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(estudiante)
}

// CrearEstudiante crea un nuevo estudiante
func CrearEstudiante(w http.ResponseWriter, r *http.Request) {
	// Leer y decodificar el estudiante desde el cuerpo de la solicitud
	var nuevoEstudiante Models.InformacionEstudiante
	err := json.NewDecoder(r.Body).Decode(&nuevoEstudiante)
	if err != nil {
		http.Error(w, "Error al decodificar la solicitud", http.StatusBadRequest)
		return
	}

	// Crear estudiante
	err = Estudiantes.CrearEstudiante(db.DB, &nuevoEstudiante)
	if err != nil {
		http.Error(w, "Error al crear estudiante", http.StatusInternalServerError)
		return
	}

	// Configurar la respuesta HTTP con el encabezado y el cuerpo JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Estudiante creado con éxito"})
}

// ObtenerEstudiantesInscritos devuelve la lista de estudiantes inscritos
func ObtenerEstudiantesInscritosHandler(w http.ResponseWriter, r *http.Request) {
	// Implementa la lógica para obtener la lista de estudiantes inscritos
	estudiantesInscritos, err := Inscripcion.ObtenerEstudiantesInscritos(db.DB)
	if err != nil {
		http.Error(w, "Error al obtener la lista de estudiantes inscritos", http.StatusInternalServerError)
		return
	}

	// Configura la respuesta HTTP con el encabezado y el cuerpo JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(estudiantesInscritos)
}

var qrCodeString string

// Implementar la función InscribirAlComedor
// Ejemplo de un manejador HTTP
func InscribirComedorHandler(w http.ResponseWriter, r *http.Request) {
	// ... Leer y decodificar la información del estudiante desde el cuerpo de la solicitud
	var nuevoEstudiante Models.InformacionEstudiante
	err := json.NewDecoder(r.Body).Decode(&nuevoEstudiante)
	if err != nil {
		http.Error(w, "Error al decodificar la solicitud", http.StatusBadRequest)
		return
	}

	// Llamar a la función InscribirAlComedor y manejar el error
	err = Inscripcion.InscribirAlComedor(nuevoEstudiante.CodigoEstudiante, qrCodeString)
	if err != nil {
		http.Error(w, "Error al inscribir al comedor: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Configurar la respuesta HTTP con el encabezado y el cuerpo JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Estudiante inscrito en el comedor con éxito"})
}

func InscripcionPorCadenaHandler(w http.ResponseWriter, r *http.Request) {
	// Leer y decodificar la información del estudiante y la cadena desde el cuerpo de la solicitud
	var solicitud struct {
		CodigoEstudiante uint   `json:"codigoEstudiante"`
		CadenaComparar   string `json:"cadenaComparar"`
	}

	err := json.NewDecoder(r.Body).Decode(&solicitud)
	if err != nil {
		http.Error(w, "Error al decodificar la solicitud", http.StatusBadRequest)
		return
	}

	// Llamar a la función InscripcionPorCadena y manejar el error
	err = Inscripcion.InscripcionPorCadena(solicitud.CodigoEstudiante, solicitud.CadenaComparar)
	if err != nil {
		http.Error(w, "Error al realizar la inscripción por cadena: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Configurar la respuesta HTTP con el encabezado y el cuerpo JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Estudiante inscrito en el comedor con éxito por cadena"})
}
