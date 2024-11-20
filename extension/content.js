document.querySelector("#publickey").remove();
document.querySelector("p#download a").remove();
const publicKey = document.createElement("p");
publicKey.classList.add("publickey");
publicKey.innerHTML = "abcabcab0101010";
document.querySelector(".input_container").append(publicKey);