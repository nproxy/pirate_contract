pragma solidity ^0.8.0;

// SPDX-License-Identifier: MIT

contract Ownable {
    address public owner;

    constructor() {
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
    function totalSupply() external view returns (uint256);

    function balanceOf(address tokenOwner) external view returns (uint256);

    function allowance(address tokenOwner, address spender)
    external
    view
    returns (uint256);

    function transfer(address to, uint256 value) external;

    function approve(address spender, uint256 value) external;

    function transferFrom(
        address from,
        address to,
        uint256 value
    ) external;

    event Transfer(address indexed from, address indexed to, uint256 tokens);
    event Approval(
        address indexed tokenOwner,
        address indexed spender,
        uint256 tokens
    );
}

contract HopSellInvestor is Ownable {
    IERC20 public USDT;
    IERC20 public HOP;
    address public HOP_FUND;
    address public beneficiary;

    struct detail {
        uint256 balance;
        uint256 mutiplier;
        uint256 claimable;
    }

    uint256 public totalLocked;
    mapping(address => detail) public balanceDetail;

    constructor(
        address HOPAddr,
        address USDTAddr,
        address fundAddr,
        address beneAddr
    ) {
        HOP = IERC20(HOPAddr);
        USDT = IERC20(USDTAddr);
        HOP_FUND = fundAddr;
        beneficiary = beneAddr;
    }

    event Exchange(address indexed user, uint256 HOPAmount);
    event Claim(address user, uint256 HOPAmount);

    //set how many HOP will be exchanged base on 1e6 USDT
    function setRateFor(address account, uint256 r) external onlyOwner {
        require(r > 0, "can't set rate to 0");
        balanceDetail[account].mutiplier = (uint256(1e12) * r) / 1e6; // 1e12 = 1e18 / 1e6
    }

    function changeAddress(address fundAddr, address beneAddr)
    external
    onlyOwner
    {
        HOP_FUND = fundAddr;
        beneficiary = beneAddr;
    }

    function addBalance(address[] memory addrs, uint256[] memory values)
    external
    onlyOwner
    {
        require(addrs.length == values.length, "array length not match");
        for (uint256 i = 0; i < addrs.length; i++) {
            require(
                HOP.balanceOf(address(this)) >= totalLocked + values[i],
                "not enough token to lock"
            );
            balanceDetail[addrs[i]].balance += values[i];
            totalLocked += values[i];
        }
    }

    function exchangeForHOP(uint256 inputUSDT) public {
        uint256 mutiplier = balanceDetail[msg.sender].mutiplier;
        require(mutiplier > 0, "rate not set");
        USDT.transferFrom(msg.sender, beneficiary, inputUSDT);
        uint256 amount = inputUSDT * mutiplier;
        require(
            HOP.balanceOf(address(this)) >= totalLocked + amount,
            "not enough token to sell"
        );
        totalLocked += amount;
        balanceDetail[msg.sender].balance += amount;
        emit Exchange(msg.sender, inputUSDT);
    }

    function receiveUSDT(uint256 inputUSDT) public{
        USDT.transferFrom(msg.sender, beneficiary, inputUSDT);
    }

    function releaseFor(address account, uint256 amount) external onlyOwner {
        require(
            balanceDetail[account].claimable + amount <=
            balanceDetail[account].balance,
            "release exceeded"
        );
        balanceDetail[account].claimable += amount;
    }

    function claimHOP(uint256 amount) public {
        require(amount > 0, "claim zero hop");
        require(amount <= balanceDetail[msg.sender].claimable, "claim too much");
        balanceDetail[msg.sender].balance -= amount;
        balanceDetail[msg.sender].claimable -= amount;
        totalLocked -= amount;
        HOP.transfer(msg.sender, amount);
        emit Claim(msg.sender, amount);
    }

    function fundWithdraw(uint256 amount) public {
        require(HOP_FUND == msg.sender, "not fund account");
        require(
            amount <= HOP.balanceOf(address(this)) - totalLocked,
            "can't withdraw"
        );
        HOP.transfer(HOP_FUND, amount);
    }

    function EmergencyWithdraw(uint256 amount, address to) external onlyOwner {
        HOP.transfer(to, amount);
    }
}

