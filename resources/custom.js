
function startCamera() {
    navigator.mediaDevices.getUserMedia({ video: true })
        .then(stream => {
                const videoElement = document.getElementById('camera-feed');
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
}

function takephoto() {
    const canvasElement = document.getElementById('canvas');
    const videoElement = document.getElementById('camera-feed');

    const context = canvasElement.getContext("2d");
    width = 320
    height = 240
    canvasElement.width = width
    canvasElement.height = height
    context.drawImage(videoElement, 0, 0, width, height);
  
    canvasElement.toBlob((blob) => {
        var xmlHttp = new XMLHttpRequest();
        xmlHttp.open("POST", "/upload");
        xmlHttp.setRequestHeader('Content-type', 'application/octet-stream');
        xmlHttp.send(blob);
    });
    //photo.setAttribute("src", data);
}

function clearphoto() {
    const canvasElement = document.getElementById('canvas');
    const context = canvasElement.getContext("2d");
    context.fillStyle = "#AAA";
    context.fillRect(0, 0, canvasElement.width, canvasElement.height);
  
    const data = canvasElement.toDataURL("image/png");
    photo.setAttribute("src", data);
}
  
  