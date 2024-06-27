// https://developer.mozilla.org/en-US/docs/Web/API/Media_Capture_and_Streams_API/Taking_still_photos

let videoElement = null;
let canvasCaptureElement = null;
let canvasSpacerElement = null;
let streaming = false;
const width = 640; // We will scale the photo width to this
let height = 0; // This will be computed based on the input stream

function startCamera() {
    videoElement = document.getElementById('camera-feed');
    canvasCaptureElement = document.getElementById("canvas-capture");
    canvasSpacerElement = document.getElementById("canvas-spacer");

    navigator.mediaDevices.getUserMedia({ video: true })
        .then(stream => {
                videoElement.srcObject = stream;
                videoElement.play();
                // videoElement.onloadedmetadata = () => {
                // };
            })
        .catch(error => {
            console.error('Error accessing camera:', error);
            // var xmlHttp = new XMLHttpRequest();
            // xmlHttp.open("GET", "/testfailed");
            // xmlHttp.send();
        }
    );
    videoElement.addEventListener(
        "canplay",
        (ev) => {
          if (!streaming) {
            height = (videoElement.videoHeight / videoElement.videoWidth) * width;
      
            videoElement.setAttribute("width", width);
            videoElement.setAttribute("height", height);
            canvasCaptureElement.setAttribute("width", width);
            canvasCaptureElement.setAttribute("height", height);
            canvasSpacerElement.setAttribute("width", width);
            canvasSpacerElement.setAttribute("height", height);
            streaming = true;
          }
        },
        false,
      );
      
}

function takephoto() {
    const context = canvasCaptureElement.getContext("2d");
    canvasCaptureElement.width = width
    canvasCaptureElement.height = height
    context.drawImage(videoElement, 0, 0, width, height);
  
    canvasCaptureElement.toBlob((blob) => {
        var xmlHttp = new XMLHttpRequest();
        xmlHttp.open("POST", "/upload");
        xmlHttp.setRequestHeader('Content-type', 'application/octet-stream');
        xmlHttp.send(blob);
    });
    //photo.setAttribute("src", data);
}

// function clearphoto() {
//     const canvasElement = document.getElementById('canvas');
//     const context = canvasElement.getContext("2d");
//     context.fillStyle = "#AAA";
//     context.fillRect(0, 0, canvasElement.width, canvasElement.height);
  
//     const data = canvasElement.toDataURL("image/png");
//     photo.setAttribute("src", data);
// }
  
  