package Logins

import (
	"ComedorGo/backend/Estudiantes"
	"ComedorGo/backend/Models"
	"ComedorGo/backend/db"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

var store = sessions.NewCookieStore([]byte("your-secret-key"))

// LoginHandler handles the login logic for students
func LoginHandler(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	fmt.Println("Body:", string(body))

	// Decodificación del cuerpo de la solicitud
	var loginData Models.LoginEstudiante
	err = json.Unmarshal(body, &loginData)
	if err != nil {
		http.Error(w, "Error decoding the request", http.StatusBadRequest)
		return
	}

	loginEstudiante, err := GetLoginEstudiantePorCodigo(db.DB, loginData.FKInformacionEstudiante)
	if err != nil {
		http.Error(w, "Student login not found", http.StatusNotFound)
		return
	}

	fmt.Println("Contraseña ingresada:", loginData.Password)
	fmt.Println("Contraseña almacenada:", loginEstudiante.Password)
	// Get student information by student code
	estudiante, err := Estudiantes.GetEstudiantePorCodigo(db.DB, loginData.FKInformacionEstudiante)
	if err != nil {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	// Logs adicionales
	fmt.Println("Contraseña ingresada:", loginData.Password)
	fmt.Println("Contraseña almacenada:", loginEstudiante.Password)

	// Compare the password using the LoginEstudiante struct
	err = bcrypt.CompareHashAndPassword([]byte(loginEstudiante.Password), []byte(loginData.Password))
	if err != nil {
		fmt.Println("Contraseñas no coinciden:", err)
		http.Error(w, "Incorrect credentials", http.StatusUnauthorized)
		return
	}
	fmt.Println("Contraseñas coinciden")

	// Create session
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, "Error creating session", http.StatusInternalServerError)
		return
	}

	// Store the student code in the session
	session.Values["codigoEstudiante"] = estudiante.CodigoEstudiante
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, "Error saving session", http.StatusInternalServerError)
		return
	}

	// Redirect to the frontend (change "/path/to/index.html" to the correct path)
	http.Redirect(w, r, "/index.html", http.StatusSeeOther)
}

func GetLoginEstudiantePorCodigo(dbInstance *gorm.DB, codigoEstudiante uint) (db.LoginEstudiante, error) {
	var loginEstudiante db.LoginEstudiante
	result := dbInstance.Where("fk_informacion_estudiante = ?", codigoEstudiante).First(&loginEstudiante)
	if result.Error != nil {
		return db.LoginEstudiante{}, result.Error
	}
	return loginEstudiante, nil
}
