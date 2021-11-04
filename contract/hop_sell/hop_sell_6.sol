pragma solidity  ^0.8.0;
// SPDX-License-Identifier: MIT

contract Ownable {
    address public owner;
    constructor(){
        owner = msg.sender;
    }
    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }

    function transferOwnership(address newOwner) public onlyOwner {
        if (newOwner != address(0)) {
            owner = newOwner;
        }
    }
}

interface IERC20 {
    function totalSupply() external view returns (uint);
    function balanceOf(address tokenOwner) external view returns (uint);
    function allowance(address tokenOwner, address spender) external view returns (uint);
    function transfer(address to, uint value) external;
    function approve(address spender, uint value) external;
    function transferFrom(address from, address to, uint value) external;
    event Transfer(address indexed from, address indexed to, uint tokens);
    event Approval(address indexed tokenOwner, address indexed spender, uint tokens);
}

contract HopSell6 is Ownable{
    
    IERC20 public USDT;
    IERC20 public HOP;
    address public HOP_FUND;
    address public beneficiary;
    
    uint256 public ON_EXCHANGE_TIME;
    bool public ON_EXCHANGE;
    
    
    mapping(address=>detail) public balanceDetail;
    
    //convert USDT uint to HOP unit, init 1 USDT = 2 HOP
    uint256 public mutiplier = 1e18 / 1e6 * 2; 
    
    constructor(address HOPAddr,
                address USDTAddr,
                address fundAddr,
                address beneAddr){
        HOP = IERC20(HOPAddr);
        USDT = IERC20(USDTAddr);
        HOP_FUND = fundAddr;
        beneficiary = beneAddr;
    }
    
    struct detail{
        uint256 totalBalance;
        uint256 claimed;
    }
    
    event Exchange(address indexed user, uint256 HOPAmount); 
    event Claim(address user, uint256 HOPAmount);
    
    //set how many HOP will be exchanged base on 1e6 USDT
    function setRate(uint256 r) external onlyOwner {
        require(r > 0, "can't set rate to 0");
        mutiplier = uint256(1e12) * r / 1e6; // 1e12 = 1e18 / 1e6
    }

    function onExchange() external onlyOwner {
        ON_EXCHANGE = true;
        ON_EXCHANGE_TIME = block.timestamp;
    }

    function changeAddress(address fundAddr, address beneAddr) external onlyOwner {
        HOP_FUND = fundAddr;
        beneficiary = beneAddr;
    }
    
    
    function editBalance(address[] memory addrs, uint256[] memory values) external onlyOwner {
        require(addrs.length == values.length, "array length not match");
        for(uint256 i=0; i< addrs.length; i++){
            balanceDetail[addrs[i]].totalBalance = values[i];
        }
    }
    
    function estimateCost(uint256 inputUSDT) public view returns (uint256){
        return inputUSDT * mutiplier;
    }
    
    function exchangeForHOP(uint256 inputUSDT) public {
        require(!ON_EXCHANGE, "already on exchange");
        uint256 half = inputUSDT * mutiplier / 2;
        require(half > 0, "input too small");
        USDT.transferFrom(msg.sender, beneficiary, inputUSDT);
        HOP.transferFrom(HOP_FUND, msg.sender, half);
        balanceDetail[msg.sender].totalBalance += half;
        emit Exchange(msg.sender, inputUSDT);
    }
    
    function accountInfo(address payer, uint256 time) public view 
                                returns (uint256 totalBalance, 
                                        uint256 claimed,
                                        uint256 claimable){
        detail memory dtl = balanceDetail[payer];
        totalBalance = dtl.totalBalance;
        claimed = dtl.claimed;
        if(!ON_EXCHANGE){
            return (totalBalance, 0, 0);
        }
        uint256 period = (time - ON_EXCHANGE_TIME) / 30 days + 1;
        if(period >= 6){
            return (totalBalance, claimed, totalBalance - claimed);
        }
        claimable = totalBalance * period / 6 - claimed;
    }
    
    // this implement can't ensure fund is surfficient for claiming
    function claimHOP(uint256 amount) public {
        require(amount > 0, "claim zero hop");
        (uint256 totalBalance, uint256 claimed, uint256 claimable) = accountInfo(msg.sender, block.timestamp);
        require(claimable + claimed <= totalBalance, "error in calculation");
        require(amount <= claimable, "claim to much");
        balanceDetail[msg.sender].claimed += amount; 
        HOP.transferFrom(HOP_FUND, msg.sender, amount);
        emit Claim(msg.sender, amount);
    }
    
}