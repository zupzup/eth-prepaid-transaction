(function() {
    var web3 = new Web3();
    web3.setProvider(new web3.providers.HttpProvider("http://localhost:8545"));

    var address = document.getElementById("address");
    address.innerText = "0x000";

    var agreement = document.getElementById("agreement");
    agreement.innerText = "Some weird Text";

    var signed = document.getElementById("signed");
    signed.innerText = "false";

    var balance = document.getElementById("balance");
    balance.innerText = "0";

    var signbutton = document.getElementById("signbutton");
    signbutton.addEventListener("click", function(event) {
        console.log("clicked sign");
    });
})();
