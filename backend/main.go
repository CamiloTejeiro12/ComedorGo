package main

import (
	"ComedorGo/backend/Api"
	"ComedorGo/backend/db"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	db.SetupDatabase()

	r := mux.NewRouter()

	// Configurar encabezados CORS
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			next.ServeHTTP(w, r)
		})
	})

	// Incluye las rutas de la API de estudiantes

	Api.EstudianteRoutes(r)
	// Ejemplo para generar un QR y guardarlo a través de la API
	r.HandleFunc("/api/generarqr", Api.GenerarYGuardarQR).Methods("POST")

	// Configura el enrutador y comienza a escuchar
	http.Handle("/", r)
	fmt.Println("Escuchando en http://localhost:8000...")
	http.ListenAndServe("localhost:8000", nil)
}
