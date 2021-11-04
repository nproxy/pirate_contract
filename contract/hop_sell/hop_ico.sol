pragma solidity  ^0.5.0;

library SafeMath {

  function mul(uint256 a, uint256 b) internal pure returns (uint256) {

    if (a == 0) {
      return 0;
    }

    uint256 c = a * b;
    require(c / a == b);

    return c;
  }


  function div(uint256 a, uint256 b) internal pure returns (uint256) {
    require(b > 0); // Solidity only automatically asserts when dividing by 0
    uint256 c = a / b;

    return c;
  }

  function sub(uint256 a, uint256 b) internal pure returns (uint256) {
    require(b <= a);
    uint256 c = a - b;

    return c;
  }

  function add(uint256 a, uint256 b) internal pure returns (uint256) {
    uint256 c = a + b;
    require(c >= a);

    return c;
  }

  function mod(uint256 a, uint256 b) internal pure returns (uint256) {
    require(b != 0);
    return a % b;
  }
}

interface IERC20 {
  function totalSupply() external view returns (uint256);

  function balanceOf(address who) external view returns (uint256);

  function allowance(address owner, address spender)
    external view returns (uint256);

  function transfer(address to, uint256 value) external returns (bool);

  function approve(address spender, uint256 value)
    external returns (bool);

  function transferFrom(address from, address to, uint256 value)
    external returns (bool);

  event Transfer(
    address indexed from,
    address indexed to,
    uint256 value
  );

  event Approval(
    address indexed owner,
    address indexed spender,
    uint256 value
  );
}

contract HOPICO {
    
    using SafeMath for uint256;
    
    uint256 public e9 = 1e9;
    address public HOP;
    address public owner;
    
    
    struct ERC20Pair{
        address tokenAddr;
        uint256 rate; // {number} <=> 1e9 HOP
    }
    
    event EXCHANGE(address indexed user, 
                    string indexed tokenName,
                    uint256 HOPAmount); 
    event PAIR(string name, bool remove);
    event DRAWFUND(address to);
    
    modifier onlyOwner {
        require(msg.sender == owner, "only owner allowed");
        _;
    }
    
    mapping(string=>ERC20Pair) public Pairs;
    
    constructor(address HOPAddr) public{
        HOP = HOPAddr;
        owner = msg.sender;
    }
    
    function checkBalance(address tokenAddr) public view returns (uint256) {
        return IERC20(tokenAddr).balanceOf(address(this));
    }
    
    function estimateCost(string memory useToken, uint256 exchangeHOPAmount) public view returns (uint256){
        ERC20Pair memory pair = Pairs[useToken];
        if (pair.tokenAddr == address(0)){
            return 0;
        }
        return exchangeHOPAmount.mul(pair.rate).div(e9);
    }
    
    function exchangeForHOP(string memory useToken, uint256 exchangeHOPAmount) public {
        ERC20Pair memory pair = Pairs[useToken];
        require(pair.tokenAddr != address(0), "can't find target token pair");
        IERC20 token = IERC20(pair.tokenAddr);
        uint256 cost = exchangeHOPAmount.mul(pair.rate).div(e9);
        require(cost >0 , "can't transfer 0 amount");
        token.transferFrom(msg.sender, address(this), cost);
        IERC20(HOP).transfer(msg.sender, exchangeHOPAmount);
        emit EXCHANGE(msg.sender, useToken, exchangeHOPAmount);
    }
    
    function editPair(string calldata symbol, address tokenAddr, uint256 rate) external onlyOwner {
        Pairs[symbol] = ERC20Pair(tokenAddr, rate);
        emit PAIR(symbol, false);
    }
    
    function deletePair(string calldata symbol) external onlyOwner {
        delete Pairs[symbol];
        emit PAIR(symbol, true);
    }
    
    function transferOwner(address newOwner) external onlyOwner {
        owner = newOwner;
    }
    
    function drawToken(address tokenAddr, address to, uint256 amount) external onlyOwner {
        IERC20(tokenAddr).transfer(to, amount);
        emit DRAWFUND(to);
    }
    
}














