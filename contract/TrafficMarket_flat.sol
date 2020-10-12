pragma solidity >=0.4.24;

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

contract owned {
    address public owner;

    constructor() public {
        owner = msg.sender;
    }

    modifier onlyOwner {
        require(msg.sender == owner);
        _;
    }

    function transferOwnership(address newOwner) onlyOwner public {
        owner = newOwner;
    }
}





contract TrafficMarket is owned {
    using SafeMath for uint256;

    IERC20 public token;

    uint public decimal = 18;

    //sys settings
    uint256 public MBytesPerToken = 1000;
    uint256 public PoolDeposit = 200000 * (10 ** decimal);
    uint256 public MinerDeposit = 2000 * (10 ** decimal);

    //storage
    address[] public Pools;
    //      pool              user
    mapping(address=>mapping(address=>userData)) public UserData;
    //      pool          miner
    mapping(address=>mapping(bytes32=>address)) public payerForMiner;

    event SettingsChanged();
    event PoolReg(address poolAddr, uint8);
    event PoolClaim(
        address indexed pool,
        address indexed user,
        uint256 packet,
        uint256 tonken,
        uint256 micrNonce,
        uint256 claimNonce);
    event Charge(
        address indexed user,
        address indexed pool,
        uint256 tokenAmount);
    event MinerEvent(
        bytes32 indexed subAddr,
        address indexed addr1,
        address addr2,
        uint8 eventType);  // 0 for join, 1 for change, 2 for retire

    struct userData {
        uint256 chargeBalance;
        uint256 epoch;
    }

    struct minerData {
        address payerAddr;
        bytes32 subAddr;
    }

    modifier poolFind(address poolAddr, uint256 index){
        require(Pools[index] == poolAddr);
        _;
    }

    constructor(address tokenAddr) public {
        token = IERC20(tokenAddr);
    }

    function emergency() public onlyOwner{
        token.transfer(msg.sender, token.balanceOf(address(this)));
    }

    function changeSettings(uint256 price, uint256 pDpos, uint256 mDpos) external onlyOwner {
        require(price >0, "invalid price");
        MBytesPerToken = price;
        PoolDeposit = pDpos * (10 ** decimal);
        MinerDeposit = mDpos * (10 ** decimal);

        emit SettingsChanged();
    }

    function getPoolList() external view returns (address[] memory) {
        return Pools;
    }

    function tokenBalance(address userAddr) public view returns (uint256, uint256, uint256){
        return (token.balanceOf(userAddr), userAddr.balance, token.allowance(userAddr, address(this)));
    }


    //pool action
    function regPool() public {
        for(uint256 i=0;i< Pools.length;i++){
            if (Pools[i] == msg.sender){
                revert("already registered");
            }
        }
        Pools.push(msg.sender);
        token.transferFrom(msg.sender, address(this), PoolDeposit);

        emit PoolReg(msg.sender, 0);
    }

    function destroyPool(uint256 index) public {
        require(Pools[index] == msg.sender, "wrong index");
        Pools[index] = Pools[Pools.length-1];
        Pools.pop();
        token.transfer(msg.sender, PoolDeposit);
        emit PoolReg(msg.sender, 1);
    }

    function claim(address user, address pool, uint256 credit, uint256 amount, uint256 micNonce, uint256 cn, bytes memory signature) public {
        uint total = amount + credit;
        userData memory ud = UserData[pool][user];
        require( ud.epoch == cn, "Claim number mismatch");
        bytes32 message = keccak256(abi.encode(this, token, user, pool, credit, amount, micNonce, cn));
        bytes32 msgHash = prefixed(message);
        require(recoverSigner(msgHash, signature) == user);

        uint tn = total.mul(1 szabo).div(MBytesPerToken);
        if (tn > ud.chargeBalance){
            tn = ud.chargeBalance;
        }
        token.transfer(msg.sender, tn);
        UserData[pool][user] = userData(ud.epoch +1, ud.chargeBalance.sub(tn));

        emit PoolClaim(pool, user, total, tn, micNonce, cn);

    }

    /// builds a prefixed hash to mimic the behavior of eth_sign.
    function prefixed(bytes32 hash) internal pure returns (bytes32) {
        return keccak256(abi.encode("\x19Ethereum Signed Message:\n32", hash));
    }
    function recoverSigner(bytes32 message, bytes memory sig) internal pure  returns (address) {
        (uint8 v, bytes32 r, bytes32 s) = splitSignature(sig);
        return ecrecover(message, v, r, s);
    }
    /// signature methods.
    function splitSignature(bytes memory sig) internal pure returns (uint8 v, bytes32 r, bytes32 s) {
        require(sig.length == 65);
        assembly {
        // first 32 bytes, after the length prefix.
            r := mload(add(sig, 32))
        // second 32 bytes.
            s := mload(add(sig, 64))
        // final byte (first byte of the next 32 bytes).
            v := byte(0, mload(add(sig, 96)))
        }
        return (v, r, s);
    }

    //user action
    function charge(address user, uint256 tokenNo, address poolAddr, uint256 index) public poolFind(poolAddr, index){
        token.transferFrom(msg.sender, address(this), tokenNo);
        UserData[poolAddr][user].chargeBalance = UserData[poolAddr][user].chargeBalance.add(tokenNo);

        emit Charge(user, poolAddr, tokenNo);
    }

    //miner action
    function joinPool(address poolAddr, uint256 index, bytes32 subAddr) public poolFind(poolAddr, index){
        token.transferFrom(msg.sender, address(this), MinerDeposit);
        require(payerForMiner[poolAddr][subAddr] == address(0), "miner registered");
        payerForMiner[poolAddr][subAddr] = msg.sender;
        emit MinerEvent(subAddr, poolAddr, address(0), 0);
    }

    function changePool(address from, address to, bytes32 subAddr) public {
        require(from != to, "invalid opt");
        require(payerForMiner[from][subAddr] == msg.sender, "not allowed");
        delete payerForMiner[from][subAddr];
        payerForMiner[to][subAddr] = msg.sender;
        emit MinerEvent(subAddr, from, to, 1);
    }

    function retireFromPool(address from, bytes32 subAddr) public {
        require(payerForMiner[from][subAddr] ==  msg.sender, "not allowed");
        delete payerForMiner[from][subAddr];
        token.transfer(msg.sender, MinerDeposit);
        emit MinerEvent(subAddr, from, address(0), 2);
    }

}

