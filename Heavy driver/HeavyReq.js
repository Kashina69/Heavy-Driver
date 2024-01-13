const { Worker } = require('worker_threads');

// Function to run the heavy computation using a worker thread
function runHeavyComputation(data, callback) {
    const worker = new Worker('./worker.js', { workerData: data });

    worker.on('message', (result) => {
        callback(result);
    });

    worker.on('error', (error) => {
        console.error(error);
    });

    worker.on('exit', (code) => {
        if (code !== 0)
            console.error(`Worker stopped with exit code ${code}`);
    });
}

// Modify theHeavyReqProcessing function to use runHeavyComputation
function theHeavyReqProcessing(data) {
    theHeavyReqProcessingTimeout = setTimeout(() => {
        try {
            // you can also add a conditin that if it take more then 10 min it will close automaticallyy 
            runHeavyComputation(data, (result) => {sendResponse();});
        } catch (error) {
            console.log(error);
        }
        
    },0);
}

export default { theHeavyReqProcessing, theHeavyReqProcessingTimeout }