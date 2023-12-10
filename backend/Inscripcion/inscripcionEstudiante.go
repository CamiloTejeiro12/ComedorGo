package Inscripcion

import (
	"ComedorGo/backend/GeneratedQR"
	"ComedorGo/backend/Models"
	"ComedorGo/backend/db"
	"fmt"
	"gorm.io/gorm"
)

func ValidarCodigoQR(codigoEstudiante uint, qrCodeString string) (Models.InformacionEstudiante, error) {
	// Obtener la cadena encriptada almacenada en la base de datos para el estudiante
	encriptacionDB, err := GeneratedQR.ObtenerEncriptacionDesdeDB(codigoEstudiante)
	if err != nil {
		return Models.InformacionEstudiante{}, err
	}

	// Desencriptar la cadena encriptada del código QR leído
	infoDesencriptada, err := GeneratedQR.Decrypta(qrCodeString)
	if err != nil {
		return Models.InformacionEstudiante{}, err
	}

	// Verificar si las cadenas desencriptadas coinciden
	if infoDesencriptada != encriptacionDB {
		return Models.InformacionEstudiante{}, fmt.Errorf("Código QR no válido")
	}

	// Si la validación es exitosa, devolver la información del estudiante
	estudiante, err := GeneratedQR.ObtenerInformacionDesencriptada(codigoEstudiante)
	if err != nil {
		return Models.InformacionEstudiante{}, err
	}

	return estudiante, nil
}

// Modificar la firma de la función InscribirAlComedor para devolver un error
func InscribirAlComedor(codigoEstudiante uint, qrCodeString string) error {
	// Validar el código QR y obtener la información del estudiante
	estudiante, err := ValidarCodigoQR(codigoEstudiante, qrCodeString)
	if err != nil {
		return err
	}

	fmt.Printf("Estudiante validado: %+v\n", estudiante)

	// Realizar la inscripción al comedor
	inscripcion := db.InscripcionComedor{
		FKInformacionEstudiante: codigoEstudiante,
	}

	// Obtener el contador actual
	contadorActual := ObtenerContadorComedor()

	// Verificar si se puede inscribir
	if contadorActual >= 800 {
		return fmt.Errorf("No se pueden realizar más inscripciones, límite alcanzado")
	}

	// Incrementar el contador
	inscripcion.Contador = contadorActual + 1

	// Guardar la inscripción en la base de datos
	result := db.DB.Create(&inscripcion)
	if result.Error != nil {
		return fmt.Errorf("Error al realizar la inscripción: %s", result.Error)
	}

	fmt.Printf("Inscripción al comedor realizada para el estudiante %d\n", codigoEstudiante)

	return nil
}

func ObtenerContadorComedor() uint {
	var contador uint

	// Consulta a la base de datos para obtener el contador actual
	result := db.DB.Model(&db.InscripcionComedor{}).Select("MAX(contador)").Scan(&contador)
	if result.Error != nil {
		fmt.Printf("Error al obtener el contador del comedor: %s\n", result.Error)
		return 0
	}

	return contador
}

// InscripcionPorCadena inscribe a un estudiante basándose en una cadena comparada con la base de datos InfoQR
func InscripcionPorCadena(codigoEstudiante uint, cadenaComparar string) error {
	// Obtener la cadena encriptada almacenada en la base de datos para el estudiante
	encriptacionDB, err := GeneratedQR.ObtenerEncriptacionDesdeDB(codigoEstudiante)
	if err != nil {
		return err
	}

	// Comparar la cadena proporcionada con la cadena encriptada de la base de datos
	if cadenaComparar != encriptacionDB {
		return fmt.Errorf("La cadena proporcionada no coincide con la base de datos")
	}

	// Realizar la inscripción al comedor
	inscripcion := db.InscripcionComedor{
		FKInformacionEstudiante: codigoEstudiante,
	}

	// Obtener el contador actual
	contadorActual := ObtenerContadorComedor()

	// Verificar si se puede inscribir
	if contadorActual >= 800 {
		return fmt.Errorf("No se pueden realizar más inscripciones, límite alcanzado")
	}

	// Incrementar el contador
	inscripcion.Contador = contadorActual + 1

	// Guardar la inscripción en la base de datos
	result := db.DB.Create(&inscripcion)
	if result.Error != nil {
		return fmt.Errorf("Error al realizar la inscripción: %s", result.Error)
	}

	fmt.Printf("Inscripción al comedor realizada para el estudiante %d\n", codigoEstudiante)

	return nil
}

func ObtenerEstudiantesInscritos(db *gorm.DB) ([]Models.InformacionEstudiante, error) {
	var estudiantesInscritos []Models.InformacionEstudiante

	// Realiza la consulta para obtener los estudiantes inscritos
	result := db.Table("inscripcion_comedors").
		Select("informacion_estudiantes.*").
		Joins("JOIN informacion_estudiantes ON inscripcion_comedors.fk_informacion_estudiante = informacion_estudiantes.codigo_estudiante").
		Scan(&estudiantesInscritos)

	if result.Error != nil {
		return nil, result.Error
	}

	return estudiantesInscritos, nil
}
