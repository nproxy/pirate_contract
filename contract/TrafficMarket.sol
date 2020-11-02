pragma solidity >=0.5.11;

import "./owned.sol";
import "./safemath.sol";
import "./IERC20.sol";

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

    event PoolReg(address poolAddr, uint8 eventType);  //0 for join, 1 for quit

    event PoolClaim(
        address indexed pool,
        address indexed user,
        uint256 minerUsedPacket,
        uint256 minerPacket,
        uint256 claimedBalance,
        uint256 poolTotalPacket,
        uint8 eventTyp);
    event Charge(
        address indexed user,
        address indexed pool,
        uint256 tokenAmount);
    event MinerEvent(
        bytes32 indexed subAddr,
        address indexed addr1,
        address addr2,
        address payAddr,
        uint8 eventType);  // 0 for join, 1 for change, 2 for retire

    struct userData {
        uint256 totalChargeBalance;
        uint256 totalTraffic;
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

    function blockChainSettings() public view returns (uint256, uint256, uint256){
        return (MBytesPerToken, PoolDeposit, MinerDeposit);
    }


    //pool action
    function regPool() public {
        for(uint256 i=0;i< Pools.length;i++){
            if (Pools[i] == msg.sender){
                revert("already registered");
            }
        }
        token.transferFrom(msg.sender, address(this), PoolDeposit);
        Pools.push(msg.sender);

        emit PoolReg(msg.sender,0);

    }

    function destroyPool(uint256 index) public {
        require(Pools[index] == msg.sender, "wrong index");
        Pools[index] = Pools[Pools.length-1];
        Pools.pop();
        token.transfer(msg.sender, PoolDeposit);
        emit PoolReg(msg.sender,1);

    }
   //
    function pclaim(address user, address pool, uint256 minerCredit, uint256 minerAmount, uint256 usedTraffic, bytes memory signature) public {
        userData memory ud = UserData[pool][user];
        require( ud.totalTraffic >= usedTraffic, "Claim number mismatch");
        bytes32 message = keccak256(abi.encode(this, token, user, pool, minerCredit, minerAmount, usedTraffic));
        bytes32 msgHash = prefixed(message);
        require(recoverSigner(msgHash, signature) == user);

        uint256 tn = (usedTraffic.sub(ud.totalTraffic)).mul(10**12).div(MBytesPerToken);
        uint256 left = ud.totalChargeBalance.sub(ud.totalTraffic.mul(10**12).div(MBytesPerToken));
        uint256 trs = tn;
        if (tn > left){
            trs = left;
            UserData[pool][user] = userData(ud.totalChargeBalance,usedTraffic.sub((tn.sub(left)).mul(MBytesPerToken).div(10**12)));
        }else{
            UserData[pool][user] = userData(ud.totalChargeBalance,usedTraffic);
        }
        token.transfer(msg.sender, trs);

        emit PoolClaim(pool, user, minerCredit,minerAmount ,trs, usedTraffic,0);

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
        UserData[poolAddr][user].totalChargeBalance = UserData[poolAddr][user].totalChargeBalance.add(tokenNo);

        emit Charge(user, poolAddr, tokenNo);
        emit PoolClaim(poolAddr,user,0,0,tokenNo,0,1);
    }

    //miner action
    function joinPool(address poolAddr, uint256 index, bytes32 subAddr) public poolFind(poolAddr, index){
        require(payerForMiner[poolAddr][subAddr] == address(0), "miner registered");
        token.transferFrom(msg.sender, address(this), MinerDeposit);
        payerForMiner[poolAddr][subAddr] = msg.sender;
        emit MinerEvent(subAddr, poolAddr, address(0),msg.sender, 0);
    }

    function changePool(address from, address to, bytes32 subAddr) public {
        require(from != to, "invalid opt");
        require(payerForMiner[from][subAddr] == msg.sender, "not allowed");
        delete payerForMiner[from][subAddr];
        payerForMiner[to][subAddr] = msg.sender;
        emit MinerEvent(subAddr, from, to,msg.sender, 1);
    }

    function retireFromPool(address from, bytes32 subAddr) public {
        require(payerForMiner[from][subAddr] ==  msg.sender, "not allowed");
        delete payerForMiner[from][subAddr];
        token.transfer(msg.sender, MinerDeposit);
        emit MinerEvent(subAddr, from, address(0),msg.sender, 2);
    }

}