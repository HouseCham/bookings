function sendBadToast(message){
    Toastify({
        text: message,
        duration: 3000,
        gravity: "top", // top - bottom
        position: "right",  // left - center - right
        className: "info",
        style: {
            background: "linear-gradient(to right, rgb(255, 65, 108), rgb(255, 75, 43))",
        }
    }).showToast();
}

function sendGoodToast(message){
    Toastify({
        text: message,
        duration: 3000,
        gravity: "top", // top - bottom
        position: "right",  // left - center - right
        className: "info",
        style: {
            background: "linear-gradient(to right, #00b09b, #96c93d)",
        }
    }).showToast();
}