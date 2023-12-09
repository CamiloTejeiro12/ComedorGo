package main

import (
	"ComedorGo/backend/Estudiantes"
	"ComedorGo/backend/GeneratedQR"
	"ComedorGo/backend/Models"
	"ComedorGo/backend/db"
	"fmt"
	//"html/template"
	//"log"
	//"net/http"
)

/*

func Index(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Fprintf(w, "<h1>Página no encontrada</h1>")
		panic(err)
	} else {
		template.Execute(w, nil)
	}
}

*/

//var DB *gorm.DB

func main() {
	db.SetupDatabase()

	/*
		// Ejemplos de uso
		nuevoEstudiante := &Models.InformacionEstudiante{
			CodigoEstudiante: 16000412371,
			Nombre:           "Juan",
			Apellido:         "Pérez",
		}

		// Crear estudiante
		err := Estudiantes.CrearEstudiante(db.DB, nuevoEstudiante)
		if err != nil {
			panic("Error al crear estudiante: " + err.Error())
		}
	*/
	// Listar todos los estudiantes
	estudiantes, err := Estudiantes.ListarEstudiantes(db.DB)
	if err != nil {
		panic("Error al obtener la lista de estudiantes: " + err.Error())
	}

	// Iterar y hacer algo con cada estudiante
	for _, estudiante := range estudiantes {
		// Hacer algo con estudiante
		fmt.Println(estudiante.CodigoEstudiante, estudiante.Nombre, estudiante.Apellido)
	}

	estudiante := Models.InformacionEstudiante{
		CodigoEstudiante: 160004127,
		Nombre:           "Esteban",
		Apellido:         "Garcia",
	}

	qrCodeString, err := GeneratedQR.GenerarQR(estudiante)
	if err != nil {
		fmt.Println("Error al generar el QR:", err)
		return
	}
	fmt.Println("QR generado:", qrCodeString)

	/*
		// Leer el código QR para obtener la información del estudiante
		estudianteLeido, err := GeneratedQR.ReadQR(qrCodeString)
		if err != nil {
			fmt.Println("Error al leer el código QR:", err)
			return
		}

		fmt.Printf("Información del estudiante:\nCódigo: %d\nNombre: %s\nApellido: %s\n",
			estudianteLeido.CodigoEstudiante, estudianteLeido.Nombre, estudianteLeido.Apellido)

	*/
	// Código de estudiante que te interesa
	//codigoEstudiante := uint(6) // Reemplaza con el código de estudiante que necesites

	// Obtener la cadena encriptada desde la base de datos
	encriptacion, err := GeneratedQR.ObtenerEncriptacionDesdeDB(estudiante.CodigoEstudiante)
	if err != nil {
		fmt.Printf("Error al obtener la cadena encriptada desde la base de datos: %s\n", err)
		return
	}

	fmt.Println("Cadena encriptada obtenida desde la base de datos:", encriptacion)

	// Obtener y mostrar la información desencriptada desde la base de datos
	informacionEstudiante, err := GeneratedQR.ObtenerInformacionDesencriptada(estudiante.CodigoEstudiante)
	if err != nil {
		fmt.Printf("Error al obtener la información desencriptada desde la base de datos: %s\n", err)
		return
	}

	fmt.Println("Información desencriptada del estudiante:")
	fmt.Println("Código:", informacionEstudiante.CodigoEstudiante)
	fmt.Println("Nombre:", informacionEstudiante.Nombre)
	fmt.Println("Apellido:", informacionEstudiante.Apellido)

}

/*

		// Ejemplo de uso
		estudiante, err := Estudiantes.GetEstudiantePorCodigo(db.DB, 1)
		if err != nil {
			panic("Error al obtener estudiante: " + err.Error())
		}

		// Actualizar información del estudiante
		err = Estudiantes.ActualizarEstudiante(160004329, "NuevoNombre", "NuevoApellido")
		if err != nil {
			panic("Error al actualizar estudiante: " + err.Error())
		}

		// Eliminar estudiante por código
		err = Estudiantes.EliminarEstudiante(160004329)
		if err != nil {
			panic("Error al eliminar estudiante: " + err.Error())
}
*/

/*

		http.HandleFunc("/", Index)
		http.HandleFunc("/Login", func(w http.ResponseWriter, r *http.Request) {
			MejorPromedio(w, r)
		})


		fmt.Println("Escuchando en http://localhost:8000...")
		http.ListenAndServe("localhost:8000", nil)
}
*/
