console.log("Starting");
document.querySelector("#publickey").remove();
document.querySelector("p#download a").remove();
const publicKey = document.createElement("p");
publicKey.classList.add("publickey");
publicKey.innerHTML = "abcabcab0101010";
document.querySelector(".input_container").append(publicKey);
chrome.runtime.sendMessage({ action: "getData" }, (response) => {
    console.log("Received data:", response);
});

const port = chrome.runtime.connect({ name: "content-background-channel" });
port.postMessage({ action: "startStream", details: "Continuous request" });
port.onMessage.addListener((message) => {
    console.log("Message to background:", message);
    if (message.data) {
        console.log("Received message:", message.data);
    }
});
port.onDisconnect.addListener(() => {
    console.error("Background disconnected.");
});
