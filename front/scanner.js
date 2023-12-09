document.addEventListener('DOMContentLoaded', function () {
  let video = document.getElementById('preview');
  let fileInput = document.getElementById('fileInput');

  if (navigator.mediaDevices && navigator.mediaDevices.getUserMedia) {
    let scanner = new Instascan.Scanner({ video: video });

    scanner.addListener('scan', function (content) {
      alert('Código QR escaneado: ' + content);
      verificarContenidoQR(content);
    });

    Instascan.Camera.getCameras().then(function (cameras) {
      if (cameras.length > 0) {
        scanner.start(cameras[0]);
      } else {
        console.error('No se encontraron cámaras.');
      }
    }).catch(function (e) {
      console.error(e);
    });
  } else {
    console.error('getUserMedia no está soportado en tu navegador.');
  }

  window.importQR = function () {
    let file = fileInput.files[0];
    if (file) {
      let reader = new FileReader();
      reader.onload = function (e) {
        let qrContent = e.target.result;
        alert('Código QR importado: ' + qrContent);
        verificarContenidoQR(qrContent);
      };
      reader.readAsText(file);
    } else {
      alert('Selecciona un archivo antes de importar.');
    }
  };

  function verificarContenidoQR(content) {
    // Puedes usar un servicio en línea o imprimir en la consola para comparar
    console.log('Contenido del código QR:', content);
    // Luego, compara el contenido con el resultado del servicio en línea
  }
});
