package Logins

import (
	"ComedorGo/backend/Models"
	"ComedorGo/backend/db"
	"github.com/gorilla/sessions"
	"net/http"
)

var store = sessions.NewCookieStore([]byte("secret-key"))

// Función para manejar el login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")

	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		var estudiante Models.InformacionEstudiante
		var login db.LoginEstudiante

		// Buscar el estudiante por nombre de usuario
		if err := db.DB.Where("nombre = ?", username).First(&estudiante).Error; err != nil {
			http.Error(w, "Usuario no encontrado", http.StatusUnauthorized)
			return
		}

		// Buscar el login del estudiante
		if err := db.DB.Where("fk_informacion_estudiante = ?", estudiante.CodigoEstudiante).First(&login).Error; err != nil {
			http.Error(w, "Error al buscar información de login", http.StatusInternalServerError)
			return
		}

		// Verificar la contraseña
		if password != login.Password {
			http.Error(w, "Contraseña incorrecta", http.StatusUnauthorized)
			return
		}

		// Iniciar sesión
		session.Values["user"] = username
		session.Save(r, w)

		// Redireccionar al dashboard u otra página
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}

}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session-name")
		if _, ok := session.Values["user"]; !ok {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}

/*
// Configurar rutas con Gorilla Mux
func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// Ruta para el login
	router.HandleFunc("/login", LoginHandler).Methods("GET", "POST")

	// Rutas autenticadas
	authRouter := router.PathPrefix("/auth").Subrouter()
	authRouter.Use(AuthMiddleware)
	authRouter.HandleFunc("/dashboard", DashboardHandler).Methods("GET")

	return router
}

*/
