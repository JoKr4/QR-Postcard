// https://developer.mozilla.org/en-US/docs/Web/API/Media_Capture_and_Streams_API/Taking_still_photos

let videoElement = null;
let canvasCaptureElement = null;
let canvasSpacerElement = null;
let streaming = false;
const width = 240; // TODO proportional to device real display width
let height = 0; // This will be computed based on the input stream
let front = false;
let cachedPhoto = null;


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

function newphoto() {
    videoElement.play();
    const context = canvasCaptureElement.getContext('2d');
    context.clearRect(0, 0, width, height);
}

function takephoto() {
    const context = canvasCaptureElement.getContext("2d");
    canvasCaptureElement.width = width
    canvasCaptureElement.height = height
    context.drawImage(videoElement, 0, 0, width, height);
  
    canvasCaptureElement.toBlob((blob) => {
        cachedPhoto = blob;
    });

    videoElement.stop();
}

function uploadCachedPhoto() {
    var xmlHttp = new XMLHttpRequest();
    xmlHttp.open("POST", "/api/postcard/upload");
    xmlHttp.setRequestHeader('Content-type', 'application/octet-stream');
    xmlHttp.send(cachedPhoto);
}

// function togglecams() {
//     document.getElementById("flip-button").onclick = () => {
//         front = !front;
//         // videoElement.stop();
//     };
//     const constraints = {
//         video: { facingMode: front ? "user" : "environment" },
//     };
// }

 