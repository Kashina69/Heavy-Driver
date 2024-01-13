// this will be a middleware 
function takeReq() {
    if (user.CurrentRequestData === user.oldRequestData ) {
        return 0;
    }
    else{
        clearTimeout(theHeavyReqProcessing)
        theHeavyReqProcessing(user.CurrentRequestData)
    }
}

// you can do this same thing in the client side like 

function sendRequest() {
    if (user.CurrentRequestData === user.oldRequestData ) {
        console.log("Dont hit the button like your father beats you up")
    }
    else{
        fetch("link",{
            method:"idk what will be used",
            body: JSON.stringify({
                userId: 2313,
                userName: "sdfhls",
                otherBigDataObj:{}
            })
        })
    }
}

function theHeavyReqProcessing(data) {
    // doing some shit 
    sendResponce()
}

// if the user is making different req and you dont want to cancel the old request you have to add it in the queue 

// you can also run multiple task simenteniously by using async or if we can use something like multithreading it is available i checked it 

// if you go with multithreatding you should use function channing to make the proper queue like fashion 

// if the computation taking 10 min broo just use rust or goland what are you doing with node.js use microsystem architecture

// this will also prevent that if one computation take less time but it was the later one will be sent after the first one 

// Do it in the client too if needed  

// if dataTransfere is high use steams idk about it that much but ya 

// also use golang or rust or socket.io for websockets to reduce the aws bill -_-

userDataQueue = []

function addUserdataInQueue (CurrentRequestData){
    userDataQueue.push(user.CurrentRequestData)
}

function runNextTaskFromQueue() {
    theHeavyReqProcessing(userDataQueue[index])
}

check 