// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.5.11;

import "./owned.sol";
import "./IERC20.sol";

contract ApplyToken is owned{

    uint public constant TokenNoPerApplication = 1000 * (10 ** 18);

    IERC20 public token;
    address public payer;

    constructor(address t, address p){
        token = IERC20(t);
        payer = p;
    }

    function applyEth(address payable userAddress) public{
        require(userAddress.balance < 0.05 ether);
        userAddress.transfer(0.1 ether);
    }

    function applyFreeToken(address payable userAddress) public{
        require(token.balanceOf(userAddress) < TokenNoPerApplication);
        token.transferFrom(payer, userAddress, TokenNoPerApplication);
    }

    function changePayer(address newPayer) public onlyOwner{
        payer = newPayer;
    }

    fallback() external payable {
        // custom function code
    }

    receive() external payable {
        // custom function code
    }
}