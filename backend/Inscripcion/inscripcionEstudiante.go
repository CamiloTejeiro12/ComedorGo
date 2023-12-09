package Inscripcion

import (
	"ComedorGo/backend/GeneratedQR"
	"ComedorGo/backend/Models"
	"ComedorGo/backend/db"
	"fmt"
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
