package GeneratedQR

import (
	"ComedorGo/backend/Models"
	"ComedorGo/backend/db"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"os"

	"github.com/skip2/go-qrcode"
)

/*
// GuardarQR guarda el código QR en la base de datos
func GuardarQR(estudiante InformacionEstudiante, qrCodeString string) error {
	// Crear una instancia de InfoQR con la información del estudiante y el código QR
	infoQR := InfoQR{
		QR:                  qrCodeString,
		FKInformacionEstudiante: estudiante.CodigoEstudiante,
	}

	// Guardar en la base de datos
	result := DB.Create(&infoQR)
	if result.Error != nil {
		return fmt.Errorf("Error al guardar el código QR en la base de datos: %s", result.Error)
	}

	return nil
}


func GuardarQR(qrCodeString string) error {
	qr := db.InfoQR{QR: qrCodeString}

	result := db.DB.Create(&qr)
	if result.Error != nil {
		return fmt.Errorf("Error al guardar el código QR en la base de datos: %s", result.Error)
	}

	return nil
}
*/

// GuardarQR guarda la representación del código QR y la cadena encriptada en la base de datos
func GuardarQR(qrCodeString string, encryptedInfo string, codigoEstudiante uint) error {
	// Crear una nueva instancia de InfoQR
	infoQR := db.InfoQR{
		FKInformacionEstudiante: codigoEstudiante,
		QRRepresentation:        qrCodeString,
		Encrypted:               encryptedInfo,
	}

	// Crear el registro en la base de datos
	result := db.DB.Create(&infoQR)
	if result.Error != nil {
		return result.Error
	}

	fmt.Println("Código QR y cadena encriptada guardados en la base de datos.")

	return nil
}

// GenerarQR genera un código QR encriptado a partir de la información del estudiante
func GenerarQR(estudiante Models.InformacionEstudiante) (string, error) {
	// Convertir la información del estudiante a cadena
	info := fmt.Sprintf("%d;%s;%s", estudiante.CodigoEstudiante, estudiante.Nombre, estudiante.Apellido)

	// Encriptar la información
	encryptedInfo, err := encrypt(info)
	if err != nil {
		return "", err
	}

	// Generar el código QR
	qrCode, err := qrcode.New(encryptedInfo, qrcode.Medium)
	if err != nil {
		return "", err
	}

	// Obtener la representación en cadena del código QR
	qrCodeString := qrCode.ToSmallString(false)

	// Guardar el código QR en la base de datos
	err = GuardarQR(qrCodeString, encryptedInfo, estudiante.CodigoEstudiante)
	if err != nil {
		return "", fmt.Errorf("Error al guardar el QR en la base de datos: %s", err)
	}

	// Directorio donde quieres guardar la imagen del código QR
	directorio := "C:/Users/tejei/Documents/ProyectoGo/ComedorGo/imagenes_qr"

	// Nombre del archivo de la imagen del código QR
	nombreArchivo := fmt.Sprintf("codigo_qr_%d.png", estudiante.CodigoEstudiante)

	// Ruta completa del archivo
	rutaCompleta := directorio + "/" + nombreArchivo

	// Crear el directorio si no existe
	err = os.MkdirAll(directorio, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("Error al crear el directorio: %s", err)
	}

	// Guardar la imagen del código QR en el archivo
	err = qrCode.WriteFile(256, rutaCompleta)
	if err != nil {
		return "", fmt.Errorf("Error al guardar la imagen del código QR: %s", err)
	}

	fmt.Printf("Imagen del código QR guardada en: %s\n", rutaCompleta)

	return qrCodeString, nil
}

/*
func ObtenerQR() (string, error) {
	var qr db.InfoQR

	result := db.DB.First(&qr)
	if result.Error != nil {
		return "", fmt.Errorf("Error al obtener el código QR de la base de datos: %s", result.Error)
	}

	fmt.Println("Obteniendo QR de la base de datos:", qr.QR)
	return qr.QR, nil
}

*/

// Funciones de encriptación y desencriptación

// AES key para encriptación/desencriptación
var key = []byte("supersecretkey32")

func encrypt(plainText string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], []byte(plainText))

	return base64.URLEncoding.EncodeToString(cipherText), nil
}

func Decrypta(cipherText string) (string, error) {
	fmt.Println("cipherText:", cipherText) // Agrega esta línea para imprimir el valor
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	cipherTextBytes, err := base64.URLEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	iv := cipherTextBytes[:aes.BlockSize]
	cipherTextBytes = cipherTextBytes[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherTextBytes, cipherTextBytes)

	return string(cipherTextBytes), nil
}

// ObtenerEncriptacionDesdeDB recupera la cadena encriptada desde la base de datos para un estudiante específico
func ObtenerEncriptacionDesdeDB(codigoEstudiante uint) (string, error) {
	// Crear una instancia de InfoQR para almacenar el resultado
	var infoQR db.InfoQR

	// Buscar el registro en la base de datos por código de estudiante
	result := db.DB.First(&infoQR, "fk_informacion_estudiante = ?", codigoEstudiante)
	if result.Error != nil {
		return "", result.Error
	}

	// Devolver la cadena encriptada
	return infoQR.Encrypted, nil
}

// ObtenerInformacionDesencriptada recupera y desencripta la información desde la base de datos para un estudiante específico
func ObtenerInformacionDesencriptada(codigoEstudiante uint) (Models.InformacionEstudiante, error) {
	// Obtener la cadena encriptada desde la base de datos
	encriptacion, err := ObtenerEncriptacionDesdeDB(codigoEstudiante)
	if err != nil {
		return Models.InformacionEstudiante{}, fmt.Errorf("Error al obtener la cadena encriptada desde la base de datos: %s", err)
	}

	// Desencriptar la información
	decryptedInfo, err := Decrypta(encriptacion)
	if err != nil {
		return Models.InformacionEstudiante{}, fmt.Errorf("Error al desencriptar la información: %s", err)
	}

	// Parsear la información desencriptada
	var estudiante Models.InformacionEstudiante
	_, err = fmt.Sscanf(decryptedInfo, "%d;%s;%s", &estudiante.CodigoEstudiante, &estudiante.Nombre, &estudiante.Apellido)
	if err != nil {
		return Models.InformacionEstudiante{}, fmt.Errorf("Error al parsear la información desencriptada: %s. DecryptedInfo: %s", err, decryptedInfo)
	}

	return estudiante, nil
}

// ReadQR lee un código QR, desencripta la información y la devuelve como una estructura InformacionEstudiante
func ReadQR(qrCodeString string) (Models.InformacionEstudiante, error) {
	// Leer el código QR desde la base de datos o de donde sea que lo hayas guardado
	//	qrCodeString := ObtenerQR()

	// Desencriptar la información
	decryptedInfo, err := Decrypta(qrCodeString)
	if err != nil {
		return Models.InformacionEstudiante{}, err
	}

	// Parsear la información desencriptada
	var estudiante Models.InformacionEstudiante
	fmt.Sscanf(decryptedInfo, "%d;%s;%s", &estudiante.CodigoEstudiante, &estudiante.Nombre, &estudiante.Apellido)

	fmt.Println("LeerQR")

	return estudiante, nil
}
