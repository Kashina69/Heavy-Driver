const heavy = require("./HeavyReq")
heavy = {theHeavyReqProcessing, theHeavyReqProcessingTimeout}
const queue = require("./queue")

// Middleware to handle user requests
function takeReq(userData,oldUserData) {
    // Your logic to check if the request is duplicate
    if (JSON.stringify(userData) === JSON.stringify(oldUserData)) {
        return;
    } else {
        // if you want to clear old request 
        clearTimeout(theHeavyReqProcessingTimeout);

        // Call the function for heavy request processing
        queue(userData);
    }
}
export default takeReq;