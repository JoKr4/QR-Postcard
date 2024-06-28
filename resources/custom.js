// https://developer.mozilla.org/en-US/docs/Web/API/Media_Capture_and_Streams_API/Taking_still_photos

let videoElement = null;
let canvasCaptureElement = null;
let canvasSpacerElement = null;
let streaming = false;
const width = 240; // TODO proportional to device real display width
let height = 0; // This will be computed based on the input stream
let front = true;
let cachedPhoto = null;


function startCamera() {
    videoElement = document.getElementById('camera-feed');
    canvasCaptureElement = document.getElementById("canvas-capture");
    canvasSpacerElement = document.getElementById("canvas-spacer");

    navigator.mediaDevices.getUserMedia({ video: { facingMode: front ? "user" : "environment" } })
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

function stopCamera() {
    const stream = videoElement.srcObject;
    const tracks = stream.getTracks();

    tracks.forEach((track) => {
        track.stop();
    });

    videoElement.srcObject = null;
    streaming = false;
}

function restartCamera() {

    stopCamera();

    navigator.mediaDevices.getUserMedia({ video: { facingMode: front ? "user" : "environment" } })
        .then(stream => {
                videoElement.srcObject = stream;
                videoElement.play();
            })
        .catch(error => {
            console.error('Error accessing camera:', error);
        }
    );
}

function newphoto() {

    resetSendtextButton()

    // reset canvas overlay of previously captured pixels
    const context = canvasCaptureElement.getContext('2d');
    context.clearRect(0, 0, width, height);
}

function resetSendtextButton() {
    sendtextButton = document.getElementById("sendtext");
    sendtextButton.className = "btn mt-2 btn-secondary"
}

function takephoto() {

    resetSendtextButton()
    
    const context = canvasCaptureElement.getContext("2d");
    canvasCaptureElement.width = width
    canvasCaptureElement.height = height
    context.drawImage(videoElement, 0, 0, width, height);
  
    canvasCaptureElement.toBlob((blob) => {
        cachedPhoto = blob;
    });
}

function uploadCachedPhoto() {
    var xmlHttp = new XMLHttpRequest();
    xmlHttp.open("POST", "/api/postcard/upload");
    xmlHttp.setRequestHeader('Content-type', 'application/octet-stream');
    xmlHttp.send(cachedPhoto);
}

function togglecams() {
    front = !front;
    restartCamera();
}
 