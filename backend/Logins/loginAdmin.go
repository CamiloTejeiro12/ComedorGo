package Logins

import (
	"ComedorGo/backend/db"
	"github.com/gorilla/mux"
	"net/http"
)

// Función para manejar el login del admin
func AdminLoginHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")

	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		var admin db.InformacionAdmin
		var loginAdmin db.LoginAdmin

		// Buscar el admin por nombre de usuario
		if err := db.DB.Where("nombre = ?", username).First(&admin).Error; err != nil {
			http.Error(w, "Admin no encontrado", http.StatusUnauthorized)
			return
		}

		// Buscar el login del admin
		if err := db.DB.Where("fk_informacion_admin = ?", admin.CodigoAdmin).First(&loginAdmin).Error; err != nil {
			http.Error(w, "Error al buscar información de login", http.StatusInternalServerError)
			return
		}

		// Verificar la contraseña
		if password != loginAdmin.Password {
			http.Error(w, "Contraseña incorrecta", http.StatusUnauthorized)
			return
		}

		// Iniciar sesión
		session.Values["admin"] = username
		session.Save(r, w)

		// Redireccionar al panel de administrador u otra página
		http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
		return
	}

	// Renderizar el formulario de login (puedes usar un template HTML para esto)
}

// Configurar rutas para el login del admin con Gorilla Mux
func SetupAdminLoginRoutes() *mux.Router {
	router := mux.NewRouter()

	// Ruta para el login del admin
	router.HandleFunc("/admin/login", AdminLoginHandler).Methods("GET", "POST")

	return router
}
