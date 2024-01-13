const { json } = require("body-parser");
const takeReq = require("./middlewareForHandlingUserReq.js")
// Install required packages using npm install
const express = require('express');
const http = require('http');
const socketIO = require('socket.io');

const app = express();
const server = http.createServer(app);
const io = socketIO(server);

let oldUserData = undefined
// Middleware to handle socket connections
io.on('connection', (socket) => {
    console.log('A user connected');

    // Listen for user data from the client
    socket.on('userData', (userData) => {
        console.log('Received user data:', userData);
        // Process the request using your logic
        takeReq(userData,oldUserData);

        // Acknowledge that the request is received
        socket.emit('requestReceived', 'Request received successfully');
    });

    // Handle disconnection
    socket.on('disconnect', () => {
        console.log('User disconnected');
    });
});

// Start the server
const PORT = 3000;
server.listen(PORT, () => {
    console.log(`Server is running on port ${PORT}`);
});
