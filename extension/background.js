function sendMessageToNativeApp(message) {
    const port = chrome.runtime.connectNative("com.example.nativeapp");
    port.postMessage(message);

    port.onMessage.addListener((response) => {
        console.log("Response from native app:", response);
    });

    port.onDisconnect.addListener(() => {
        if (chrome.runtime.lastError) {
            console.error("Error:", chrome.runtime.lastError.message);
        } else {
            console.log("Disconnected from native app");
        }
    });
}

sendMessageToNativeApp({ text: "Hello from Chrome extension!" });