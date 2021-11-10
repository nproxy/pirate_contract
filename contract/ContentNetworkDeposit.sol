// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

import "./owned.sol";
import "./IERC20.sol";

contract ContentNetworkDeposit is owned{

    struct Bill{
        address user;
        address pool;
        uint256 tokenNo;
        uint256 dueDay;
    }

    constructor(address t){
        token = IERC20(t);
    }

    /*************************************************
    *           system parameters
    *************************************************/
    uint256 public CurrentBID = 1000000;
    bool public isOpen = true;
    uint256 public userRewardFor100Hops = 14;
    uint256 public poolRewardFor100Hops = 1;
    uint256 public billLifeInDays = 90;
    uint256 public constant TokenDecimal = 10 ** 18;


    /*************************************************
    *           variables
    *************************************************/
    IERC20 public token;
    mapping(address=>uint256[]) public userBillIDList;
    mapping(uint256=>Bill) public billDetails;
    mapping(address=>uint256[]) public poolsBillIDList;
    mapping(address=>uint256)public poolsFeeIncome;

    /*************************************************
    *           events
    *************************************************/
    event UserDeposit(address indexed pool, address indexed user, uint256 tokenNo);
    event UserWithDraw(address indexed pool, address indexed user, uint256 tokenNo, uint256 rewards);
    event PoolWithDraw(address indexed pool, uint256 tokenNo);

    /*************************************************
    *           modifier
    *************************************************/
    modifier usableBill(uint256 bid) {
        Bill memory bill = billDetails[bid];
        require(bill.dueDay != 0 && bill.tokenNo > 0, "invalid bill");
        require(bill.pool != address(0) && bill.user == msg.sender, "no right");
        require(block.timestamp > bill.dueDay, "not be due");
        _;
    }
    modifier serviceOpen(){
        require(isOpen, "service closed now");
        _;
    }
    function removeValue(uint256[] storage self, uint256 value) private{

        for (uint idx = 0; idx < self.length; idx++){
            if (self[idx] == value){
                self[idx]=0;
            }
        }
    }

    /*************************************************
    *           owner operation
    *************************************************/
    function openOperation(bool op) public onlyOwner{
        isOpen = op;
    }
    function setUserRate(uint256 newRate) public onlyOwner{
        userRewardFor100Hops = newRate;
    }
    function setPoolRate(uint256 newRate) public onlyOwner{
        poolRewardFor100Hops = newRate;
    }
    function setBillLifeInDay(uint256 newDays) public onlyOwner{
        billLifeInDays = newDays;
    }


    /*************************************************
    *           normal operation
    *************************************************/
    function depositForPool(address poolAddress, uint256 tokenNo) public serviceOpen{
        require(address(0) != poolAddress, "invalid pool address");
        require(tokenNo != 0, "invalid token no");

        uint256 noInDecimal = TokenDecimal * tokenNo;
        token.transferFrom(msg.sender, address(this), noInDecimal);

        CurrentBID++;
        billDetails[CurrentBID] = Bill(msg.sender, poolAddress, noInDecimal, block.timestamp + billLifeInDays * 1 days);

        userBillIDList[msg.sender].push(CurrentBID);

        poolsBillIDList[poolAddress].push(CurrentBID);

        emit UserDeposit(poolAddress, msg.sender, noInDecimal);
    }


    function withDrawFromPool(uint256 bid)  public serviceOpen usableBill(bid){

        Bill memory bill = billDetails[bid];

        removeValue(userBillIDList[bill.user], bid);
        removeValue(poolsBillIDList[bill.pool], bid);

        delete billDetails[bid];
        uint256 rewards = bill.tokenNo / 100 * userRewardFor100Hops;
        token.transfer(bill.user, bill.tokenNo + rewards);
        uint256 fee = bill.tokenNo / 100 * poolRewardFor100Hops;
        poolsFeeIncome[bill.pool] = poolsFeeIncome[bill.pool] + fee;

        emit UserWithDraw(bill.pool, msg.sender, bill.tokenNo, rewards);
    }

    function poolIncomeWithDraw() public serviceOpen{
        uint256 tokenNo = poolsFeeIncome[msg.sender];
        require(tokenNo > 0, "no founds");

        poolsFeeIncome[msg.sender] = 0;
        token.transfer(msg.sender, tokenNo);

        emit PoolWithDraw(msg.sender, tokenNo);
    }

    /*************************************************
    *           query
    *************************************************/
    function queryBillInterest(uint256 bid) view public returns(uint256){
        Bill memory bill = billDetails[bid];
        if (bill.tokenNo == 0){
            return 0;
        }

        return bill.tokenNo / 100 * userRewardFor100Hops;
    }

    function hopBalanceOf(address addr)view public returns(uint256){
        return token.balanceOf(addr);
    }

    function billListOfUser(address usr) view public returns (uint256[] memory){
        return userBillIDList[usr];
    }

    function billListOfPool(address pool) view public returns (uint256[] memory){
        return poolsBillIDList[pool];
    }
}