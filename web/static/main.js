//<hostname>:<port>
var currentLocation = window.location.hostname+":"+window.location.port;
var client = new WebSocket('ws://'+currentLocation+'/ws');

client.onopen = () => {
    console.log('WebSocket Client Connected');
};
client.onmessage = (message) => {
    loadTimer()
};

async function loadTimer() {
    let url = 'http://'+currentLocation+'/frame';

    fetch(url)
        .then(response=>response.json())
        .then(t => {
            renderTimer(t)
        });
}

function renderTimer(t) {
    document.getElementById("timer_span").textContent=t.time;
}
