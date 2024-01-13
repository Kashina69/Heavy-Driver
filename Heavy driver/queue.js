theHeavyReqProcessing= require("./HeavyReq")
// Queue for user data
const userDataQueue = [];

// Function to add user data to the queue
function addUserdataInQueue(CurrentRequestData) {
    userDataQueue.push(CurrentRequestData);
}

// Function to run the next task from the queue
function runNextTaskFromQueue() {
    if (userDataQueue.length > 0) {
        theHeavyReqProcessing(userDataQueue.shift());
    }
}
