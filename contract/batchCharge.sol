pragma solidity ^0.5.12;

import "./IERC20.sol";
import "./TrafficMarket.sol";

contract BatchCharge {
    
    IERC20 public token;
    TrafficMarket public market;
    
    constructor (address t, address m) public {
        token = IERC20(t);
        market = TrafficMarket(m);
        token.approve(m, token.totalSupply());
    }
    
    function batchCharge(address[] calldata user, uint256[] calldata charge, address[] calldata pool, uint256[] calldata index) external {
        uint256 l = user.length;
        require(l == charge.length && l == pool.length && l == index.length, "length mismatch");
        uint256 cost = 0;
        for(uint256 i = 0; i<l; i++){
            cost += charge[i];
        }
        token.transferFrom(msg.sender, address(this), cost);
        for(uint256 i = 0; i<l; i++){
            market.charge(user[i], charge[i], pool[i], index[i]);
        }
    }
}
