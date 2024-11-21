var messages = [];

function sendMessageToNativeApp(message) {
    const port = chrome.runtime.connectNative("com.example.nativeapp");
    port.postMessage(message);

    port.onMessage.addListener((response) => {
        messages.push(response)
    });

    port.onDisconnect.addListener(() => {
        if (chrome.runtime.lastError) {
            console.error("Error:", chrome.runtime.lastError.message);
        } else {
            console.log("Disconnected from native app");
        }
    });
}

chrome.runtime.onMessage.addListener((message, sender, sendResponse) => {
    if (message.action === "getData") {
        sendMessageToNativeApp({ text: "Hello from Chrome extension!" });
        sendResponse(messages);
        return true;
    }
    return false;
});

chrome.runtime.onConnect.addListener((port) => {
    if (port.name === "content-background-channel") {
        console.log("Connection established with content_script:", port);
        const sendData = () => {
            port.postMessage({ data: `Data update: ${new Date().toISOString()}` });
        };

        const intervalId = setInterval(sendData, 5000);

        port.onMessage.addListener((message) => {
            console.log("Message received from content_script:", message);

            if (message.action === "stopStream") {
                console.log("Request to stop the stream.");
                clearInterval(intervalId); port.disconnect();
            }
        });

        port.onDisconnect.addListener(() => {
            console.log("Connection with content_script closed.");
            clearInterval(intervalId);
        });
    }
});
