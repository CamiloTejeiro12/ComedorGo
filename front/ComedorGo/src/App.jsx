import { useState } from 'react';

function App() {
  const [count, setCount] = useState(0);

  function desencriptarInfoQR() {
    // Obtener valores del formulario
    var encrypted = document.getElementById('encryptedInput').value;
    var codigoEstudiante = document.getElementById('codigoEstudianteInput').value;

    // Validar que se ingresen datos
    if (!encrypted || !codigoEstudiante) {
      document.getElementById('errorMensaje').innerText =
        'Ingrese la cadena encriptada y el código de estudiante.';
      return;
    }

    // Limpiar mensajes de error
    document.getElementById('errorMensaje').innerText = '';

    // Crear objeto con la información
    var infoQR = {
      encrypted: encrypted,
      codigoEstudiante: codigoEstudiante,
    };

    // Realizar la solicitud al 
    fetch('http://localhost:8000/api/desencriptarqr', {
      method: 'POST', // Agrega esta líneas
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(infoQR),
    })
      .then((response) => {
        // Manejar errores de red
        if (!response.ok) {
          throw new Error('Error de red al realizar la solicitud.');
        }
        return response.json();
      })
      .then((data) => {
        // Mostrar el resultado en el contenedor
        document.getElementById('resultado').innerText = data.decryptedInfo;
      })
      .catch((error) => {
        // Manejar errores generales
        console.error('Error al realizar la solicitud:', error.message);
        document.getElementById('errorMensaje').innerText =
          'Error al realizar la solicitud: ' + error.message;
      });
  }

  return (
    <>
      <h2>Desencriptar InfoQR</h2>

      <label htmlFor="encryptedInput">Cadena encriptada:</label>
      <input type="text" id="encryptedInput" placeholder="Ingrese la cadena encriptada" />

      <label htmlFor="codigoEstudianteInput">Código de estudiante:</label>
      <input
        type="number"
        id="codigoEstudianteInput"
        placeholder="Ingrese el código de estudiante"
      />

      <button onClick={desencriptarInfoQR}>Desencriptar</button>

      <div id="resultContainer">
        <h3>Resultado:</h3>
        <p id="resultado"></p>
        <p id="errorMensaje" style={{ color: 'red' }}></p>
      </div>
    </>
  );
}

export default App;
