pragma solidity >=0.5.11;

import "./owned.sol";
import "./IERC20.sol";
import "./TrafficMarket.sol";

contract PirateDeposit is owned{
    using SafeMath for uint256;

    uint private Decimal = 18;
    uint private minDepositInterval = 20 days;
    uint private poolDrawInterval = 50 days;
    uint private minOneMonth = 28 days;
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
        uint256 lastRateIndex;
    }

    struct drawRate{
        uint256 rate;
        uint256 drawRateTime;
        uint256 totalReward;
        uint256 leftReward;
    }

    event UserDepositEvent(address indexed user, address indexed pool, uint256 tokenNumber, uint256 depositTime);
    event UserWithDrawDepositEvent(address indexed user, address indexed pool, uint256 depositTime);
    event UserDrawRewardEvent(address indexed user, address indexed pool, uint256 reward, uint256 lastIndex);
    event AddRewardEvent(address pool, uint256 rate, uint256 totalReward, uint256 rewardTime);
    event PoolDrawRewardEvent(address pool, uint256 profits);

    function changeCoordinator(address coorD) external onlyOwner{
        coordinator = coorD;
    }
    function changeDepositInterval(uint d) public mustCoordinator{
        minDepositInterval = d * (1 days);
    }
    function changeMinDeposit(uint256 min) public mustCoordinator{
        minDeposit = min;
    }

    function doUserDeposit(address pool, uint256 tokenNumber,uint256 index) public{
        require(market.LegalPool(pool,index),"pool not found");
        require(tokenNumber >= minDeposit,"token must large than 100" );

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
            if ((now - DrawRates[pool][DrawRates[pool].length-1].drawRateTime) < (minOneMonth)){
                revert("must large than one month");
            }
        }

        token.transferFrom(msg.sender,address(this),totalReward);

        uint256 NowTime = now;
        DrawRates[pool].push(drawRate(rate,NowTime, totalReward,totalReward));

        emit AddRewardEvent(pool,rate,totalReward,NowTime);
    }

    function drawReward(address pool, uint256 poolIndex) public {
        require(market.LegalPool(pool,poolIndex),"pool not found");
        userDeposits memory uds = DepositDatas[pool][msg.sender];
        require(DrawRates[pool].length > 0, "no withdraw reward rate");
        require(uds.uds.length >0, "no deposit");
        require(uds.lastRateIndex < DrawRates[pool].length-1, "no reward for withdraw");

        uint256 sumDs = 0;

        for(uint256 i=uds.lastRateIndex.add(1);i<DrawRates[pool].length; i++){
            uint256 sumT = 0;
            for(uint256 j=0;j<uds.uds.length;j++){
                if (DrawRates[pool][i].drawRateTime - uds.uds[j].depositTime > minDepositInterval){
                    sumT.add(uds.uds[j].totalDeposit);
                }
            }
            if (sumT > 0){
                sumDs.add(sumT*DrawRates[pool][i].rate);
                DrawRates[pool][i].leftReward.sub(sumT);
            }
        }

        if (sumDs > 0){
            token.transfer(msg.sender, sumDs);
        }

        DepositDatas[pool][msg.sender].lastRateIndex = DrawRates[pool].length-1;
//        DrawRates[pool][msg.sender].leftReward.sub(sumDs);

        emit UserDrawRewardEvent(msg.sender, pool, sumDs, DepositDatas[pool][msg.sender].lastRateIndex);
    }

    function poolDrawReward(uint256 poolIndex) public{
        require(market.LegalPool(msg.sender,poolIndex),"pool not found");
        require(now > (DrawRates[msg.sender][DrawRates[msg.sender].length-1].drawRateTime + poolDrawInterval), "pool draw profit error");

        uint256 sumDs = 0;
        for(uint256 i=0;i<DrawRates[msg.sender].length;i++){
            sumDs.add(DrawRates[msg.sender][i].leftReward);
            DrawRates[msg.sender][i].leftReward = 0;
        }

        if (sumDs > 0){
            token.transfer(msg.sender, sumDs);
        }
        emit PoolDrawRewardEvent(msg.sender, sumDs);
    }

}
