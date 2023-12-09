package Estudiantes

import (
	"ComedorGo/backend/Models"
	"fmt"

	//"ComedorGo/backend/db"
	"gorm.io/gorm"
)

// Obtener información de un estudiante por su código
func GetEstudiantePorCodigo(db *gorm.DB, codigoEstudiante uint) (*Models.InformacionEstudiante, error) {
	estudiante := &Models.InformacionEstudiante{}
	result := db.First(estudiante, codigoEstudiante)
	if result.Error != nil {
		return nil, result.Error
	}
	return estudiante, nil
}

// Crear estudiante
func CrearEstudiante(db *gorm.DB, estudiante *Models.InformacionEstudiante) error {
	// Verificar si el estudiante ya existe
	existingEstudiante := &Models.InformacionEstudiante{}
	result := db.Where("codigo_estudiante = ?", estudiante.CodigoEstudiante).First(existingEstudiante)
	if result.RowsAffected > 0 {
		return fmt.Errorf("El estudiante con el código %d ya está registrado", estudiante.CodigoEstudiante)
	}

	// Crear estudiante
	result = db.Create(estudiante)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Actualizar información de un estudiante
func ActualizarEstudiante(db *gorm.DB, codigoEstudiante uint, nombre, apellido string) error {
	estudiante, err := GetEstudiantePorCodigo(db, codigoEstudiante)
	if err != nil {
		return err
	}

	// Actualizar campos
	estudiante.Nombre = nombre
	estudiante.Apellido = apellido

	// Guardar cambios
	result := db.Save(estudiante)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Eliminar estudiante por su código
func EliminarEstudiante(db *gorm.DB, codigoEstudiante uint) error {
	result := db.Delete(&Models.InformacionEstudiante{}, codigoEstudiante)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Listar todos los estudiantes
func ListarEstudiantes(db *gorm.DB) ([]Models.InformacionEstudiante, error) {
	var estudiantes []Models.InformacionEstudiante
	result := db.Find(&estudiantes)
	if result.Error != nil {
		return nil, result.Error
	}
	return estudiantes, nil
}
