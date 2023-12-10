package Logins

import (
	"ComedorGo/backend/Estudiantes"
	"ComedorGo/backend/Models"
	"ComedorGo/backend/db"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

var store = sessions.NewCookieStore([]byte("your-secret-key"))

// LoginHandler handles the login logic for students
func LoginHandler(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error reading request body"})
		return
	}
	fmt.Println("Body:", string(body))

	// Decodificación del cuerpo de la solicitud
	var loginData Models.LoginEstudiante
	err = json.Unmarshal(body, &loginData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error decoding the request"})
		return
	}

	loginEstudiante, err := GetLoginEstudiantePorCodigo(db.DB, loginData.FKInformacionEstudiante)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student login not found"})
		return
	}

	fmt.Println("Contraseña ingresada:", loginData.Password)
	fmt.Println("Contraseña almacenada:", loginEstudiante.Password)
	// Get student information by student code
	estudiante, err := Estudiantes.GetEstudiantePorCodigo(db.DB, loginData.FKInformacionEstudiante)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	// Logs adicionales
	fmt.Println("Contraseña ingresada:", loginData.Password)
	fmt.Println("Contraseña almacenada:", loginEstudiante.Password)

	// Compare the password using the LoginEstudiante struct
	err = bcrypt.CompareHashAndPassword([]byte(loginEstudiante.Password), []byte(loginData.Password))
	if err != nil {
		fmt.Println("Contraseñas no coinciden:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect credentials"})
		return
	}
	fmt.Println("Contraseñas coinciden")

	// Create session
	session, err := store.Get(c.Request, "session-name")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating session"})
		return
	}

	// Store the student code in the session
	session.Values["codigoEstudiante"] = estudiante.CodigoEstudiante
	err = session.Save(c.Request, c.Writer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving session"})
		return
	}

	// Redirect to the frontend (change "/path/to/index.html" to the correct path)
	c.Redirect(http.StatusSeeOther, "/index.html")
}

func GetLoginEstudiantePorCodigo(dbInstance *gorm.DB, codigoEstudiante uint) (db.LoginEstudiante, error) {
	var loginEstudiante db.LoginEstudiante
	result := dbInstance.Where("fk_informacion_estudiante = ?", codigoEstudiante).First(&loginEstudiante)
	if result.Error != nil {
		return db.LoginEstudiante{}, result.Error
	}
	return loginEstudiante, nil
}
