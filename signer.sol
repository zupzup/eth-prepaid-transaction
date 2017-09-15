pragma solidity ^0.4.6;

contract Signer {
    address public owner = msg.sender;
    struct Agreement {
        string stringToAgreeOn;
        bool signed;
        bool initialized;
    }
    mapping (address=> Agreement) agreements;

    modifier onlyBy(address _account)
    {
        require(msg.sender == _account);
        _;
    }

    function createAgreement(string _stringToAgreeOn, address customer) payable public onlyBy(owner) returns (bool success) {
        agreements[customer] = Agreement(_stringToAgreeOn, false, true);
        return true;
    }

    function signAgreement() payable public returns (bool success) {
        var agreement = agreements[msg.sender];
        require(agreement.initialized == true);
        require(agreement.signed == false);
        agreement.signed = true;
        return true;
    }

    function getAgreement() constant public returns(string stringToAgreeOn, bool signed) {
        var agreement = agreements[msg.sender];
        require(agreement.initialized == true);
        return (agreement.stringToAgreeOn, agreement.signed);
    }
}
