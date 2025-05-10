let socket = new WebSocket("ws://localhost:8080/ws");
socket.onmessage = (message) => {console.log(message.data) };


socket.onopen = () => {

};

socket.onclose = (event) => {
    console.log("Socket closed connection: ", event);
};
socket.onerror = (error) => {
    console.log("Socket error: ", error)
};
