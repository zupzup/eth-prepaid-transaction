var abi = "[{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAgreement\",\"outputs\":[{\"name\":\"stringToAgreeOn\",\"type\":\"string\"},{\"name\":\"signed\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_stringToAgreeOn\",\"type\":\"string\"},{\"name\":\"customer\",\"type\":\"address\"}],\"name\":\"createAgreement\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"signAgreement\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"type\":\"function\"}]";

(function() {
    var web3 = new Web3();
    web3.setProvider(new web3.providers.HttpProvider("http://localhost:8545"));
    var accounts = web3.eth.accounts;
    var currentAccount = accounts[1];
    var balance = web3.eth.getBalance(currentAccount);
    console.log(accounts)

    var Signer = web3.eth.contract(JSON.parse(abi));
    var instance = Signer.at("0x34c349e99ccc2f258cbc663ca508c805194cdf10")
    console.log(instance)

    instance.owner(function(err, results) {
        if (err) {
            console.log(err, "error happened")
        } else {
            console.log(results, "results happened")
        }
    });

    var address = document.getElementById("address");
    address.innerText = currentAccount;

    var balanceElement = document.getElementById("balance");
    balanceElement.innerText = balance.toNumber();

    var agreement = document.getElementById("agreement");
    agreement.innerText = "Some weird Text";

    var signed = document.getElementById("signed");
    signed.innerText = "false";

    var signbutton = document.getElementById("signbutton");
    signbutton.addEventListener("click", function(event) {
        console.log("clicked sign");
    });
})();
