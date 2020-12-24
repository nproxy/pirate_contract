pragma solidity >=0.5.11;

import "./owned.sol";
import "./IERC20.sol";
import "./TrafficMarket.sol";

contract PirateDeposit is owned{
    using SafeMath for uint256;

    uint private Decimal = 18;
    uint private days20 = 20 days;
    uint256 private minDeposit = 100**Decimal;
    IERC20 public token;
    TrafficMarket public market;
    address public coordinator;

    modifier mustCoordinator(){
        require(msg.sender == coordinator, "not a coordinator");
        _;
    }

    constructor (address t, address m) public{
        token = IERC20(t);
        market = TrafficMarket(m);
        coordinator =  msg.sender;
    }
    //pool                user
    mapping(address=>mapping(address=>userDeposits)) public DepositDatas;
    //pool
    mapping(address=>drawRate[]) public DrawRates;

    struct userDeposit{
        uint256 totalDeposit;
        uint256 depositTime;
    }

    struct userDeposits{
        userDeposit[] uds;
        uint256 lastDrawRateTime;
    }

    struct drawRate{
        uint256 rate;
        uint256 drawRateTime;
        uint256 totalReward;
    }

    event UserDepositEvent(
        address indexed user,
        address indexed pool,
        uint256 tokenNumber,
        uint256 depositTime
    );

    event UserWithDrawDepositEvent(
        address indexed user,
        address indexed pool,
        uint256 depositTime
    );

    event UserDrawRewardEvent(
        address indexed user,
        address indexed pool,
        uint256 reward,
        uint256 lastRewardTime
    );

    event AddRewardEvent(address pool, uint256 rate, uint256 totalReward, uint256 rewardTime);

    function changeCoordinator(address coorD) external onlyOwner{
        coordinator = coorD;
    }
    function changeDepositInterval(uint d) public mustCoordinator{
        days20 = d * (1 days);
    }
    function changeMinDeposit(uint256 min) public mustCoordinator{
        minDeposit = min;
    }

    function doUserDeposit(address pool, uint256 tokenNumber,uint256 index) public{
        require(market.LegalPool(pool,index),"pool not found");
        require(tokenNumber > minDeposit,"token must large than 100" );

        token.transferFrom(msg.sender, address(this), tokenNumber);

        uint256 NowTime = now;

        DepositDatas[pool][msg.sender].uds.push(userDeposit(tokenNumber,NowTime));

        emit UserDepositEvent(msg.sender,pool,tokenNumber,NowTime);
    }


    function userWithDrawDeposit(address pool, uint256 recordTime, uint256 index, uint256 poolIndex) public {
        require(market.LegalPool(pool,poolIndex),"pool not found");
        userDeposits memory uds = DepositDatas[pool][msg.sender];
        require(uds.uds[index].depositTime == recordTime, "not a valid Deposit");

        token.transfer(msg.sender, uds.uds[index].totalDeposit);

        DepositDatas[pool][msg.sender].uds[index] = DepositDatas[pool][msg.sender].uds[uds.uds.length-1];
        DepositDatas[pool][msg.sender].uds.pop();

        emit UserWithDrawDepositEvent(msg.sender,pool,recordTime);
    }

    function addReward(address pool, uint256 rate, uint256 totalReward) public mustCoordinator{
        if (DrawRates[pool].length > 0){
            if ((now - DrawRates[pool][DrawRates[pool].length-1].drawRateTime) < (28 days)){
                revert("must large than one month");
            }
        }
        uint256 NowTime = now;
        DrawRates[pool].push(drawRate(rate,NowTime, totalReward));

        emit AddRewardEvent(pool,rate,totalReward,NowTime);
    }

    function drawReward(address pool, uint256 poolIndex) public {
        require(market.LegalPool(pool,poolIndex),"pool not found");
        userDeposits memory uds = DepositDatas[pool][msg.sender];
        require(DrawRates[pool].length > 0, "no withdraw reward rate");
        require(uds.uds.length >0, "no deposit");
        require((uds.lastDrawRateTime + (28 days)) <=  DrawRates[pool][DrawRates[pool].length-1].drawRateTime, "no reward for withdraw");

        uint256 sumDs = 0;

        for(uint256 i=0;i<DrawRates[pool].length; i++){
            if (uds.lastDrawRateTime > DrawRates[pool][i].drawRateTime){
                continue;
            }

            uint256 sumT = 0;
            for(uint256 j=0;j<uds.uds.length;j++){
                if (DrawRates[pool][i].drawRateTime - uds.uds[j].depositTime > days20){
                    sumT += uds.uds[j].totalDeposit;
                }
            }
            if (sumT > 0){
                sumDs += sumT*DrawRates[pool][i].rate;
            }
        }

        if (sumDs > 0){
            token.transfer(msg.sender, sumDs);
        }

        DepositDatas[pool][msg.sender].lastDrawRateTime = DrawRates[pool][DrawRates[pool].length-1].drawRateTime;

        emit UserDrawRewardEvent(msg.sender, pool, sumDs, DepositDatas[pool][msg.sender].lastDrawRateTime);
    }
}
