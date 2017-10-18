var abi = "[{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getAgreement\",\"outputs\":[{\"name\":\"stringToAgreeOn\",\"type\":\"string\"},{\"name\":\"signed\",\"type\":\"bool\"},{\"name\":\"initialized\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_stringToAgreeOn\",\"type\":\"string\"},{\"name\":\"customer\",\"type\":\"address\"}],\"name\":\"createAgreement\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"signAgreement\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]";

var host = "http://localhost:8080";

(function() {
    var web3 = new Web3();
    web3.setProvider(new web3.providers.HttpProvider("http://localhost:8545"));
    var accounts = web3.eth.accounts;
    var currentAccount = accounts[1];
    var balance = web3.eth.getBalance(currentAccount);

    var Signer = web3.eth.contract(JSON.parse(abi));
    var instance = Signer.at("0x34c349e99ccc2f258cbc663ca508c805194cdf10")

    var agreement = document.getElementById("agreement");
    agreement.innerText = "";

    var signed = document.getElementById("signed");
    signed.innerText = "";

    var refresh = document.getElementById("refresh");
    refresh.addEventListener("click", function(event) {
        instance.getAgreement(currentAccount, function(err, results) {
            if (err || !results[2]) {
                agreement.innerText = ""
                signed.innerText = "";
            } else {
                agreement.innerText = results[0];
                signed.innerText = results[1];
            }
        });
    });

    var refreshbalance = document.getElementById("refreshbalance");
    refreshbalance.addEventListener("click", function(event) {
        var newBalance = web3.eth.getBalance(currentAccount);
        var balanceElement = document.getElementById("balance");
        balanceElement.innerText = newBalance.toNumber();
    });

    var address = document.getElementById("address");
    address.innerText = currentAccount;

    var balanceElement = document.getElementById("balance");
    balanceElement.innerText = balance.toNumber();


    var signbutton = document.getElementById("signbutton");
    signbutton.addEventListener("click", function(event) {
        var tx = {
            from: currentAccount,
            value: 0,
            gas: 100000
        }
        instance.signAgreement.sendTransaction(tx, function(err, res) {
            if (err) {
                return alert(err)
            }
        })
    });

    var agreementButton = document.getElementById("agreementButton");
    agreementButton.addEventListener("click", function(event) {
        var agreementInput = document.getElementById("agreementInput");
        agreementValue = agreementInput.value
        superagent.post(host + "/agreement").send({
            account: currentAccount,
            agreement: agreementValue,
        }).end(function(err, res) {
            if (err) {
                return alert(res.text)
            }
            console.log(res);
        });
    });
})();
