// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

import "./owned.sol";
import "./IERC20.sol";
import "./safemath.sol";

contract PoolDeposit is owned{
using SafeMath for uint256;

IERC20 public token;
mapping(address=>mapping(address=>uint256)) public depositRecord;


event UserDeposit(address indexed pool, address indexed user, uint256 tokenNo);
event UserWithDraw(address indexed pool, address indexed user, uint256 tokenNo);

constructor(address t){
token = IERC20(t);
}

function depositForPool(address poolAddress, uint256 tokenNo) public{

token.transferFrom(msg.sender, address(this), tokenNo);

depositRecord[poolAddress][msg.sender] = depositRecord[poolAddress][msg.sender].add(tokenNo);

emit UserDeposit(poolAddress, msg.sender, tokenNo);
}

function withDrawFromPool(address poolAddress) public{

uint256 tokenNo = depositRecord[poolAddress][msg.sender];
require(tokenNo > 0, "insufficient fouds");

depositRecord[poolAddress][msg.sender] = 0;
token.transferFrom(address(this), msg.sender, tokenNo);

emit UserWithDraw(poolAddress, msg.sender, tokenNo);
}
}