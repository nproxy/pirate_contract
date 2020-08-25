pragma solidity >=0.4.24;

import "./owned.sol";
import "./safemath.sol";
import "./IERC20.sol";


contract HopShop is owned{
    using SafeMath for uint256;

    event TokensPurchased(
        address indexed purchaser,
        uint256 value,
        uint256 amount
    );

    address payable public  _adminWallet;
    uint256 public _rateLevel1 = 300;
    uint256 public _rateLevel2 = 240;
    uint256 public TokenNoToSellInLevel1;
    uint256 public hasRaised = 0;
    IERC20 public _token;

    constructor(address payable wallet, IERC20 token) public{
        require(wallet != address(0));
        require(address(token) != address(0));
        _adminWallet = wallet;
        _token = token;
        TokenNoToSellInLevel1 = 2.1e7 * (10 ** 18);
    }

    function () payable external{
        require(msg.value > 0);

        uint256 tokenNo = 0;
        if (hasRaised < TokenNoToSellInLevel1){
            tokenNo = msg.value.mul(_rateLevel1);
        }else{
            tokenNo = msg.value.mul(_rateLevel2);
        }
        require(remainingTokens() >= tokenNo);

        _token.transferFrom(_adminWallet, msg.sender, tokenNo);
        emit TokensPurchased(msg.sender, msg.value, tokenNo);
        _adminWallet.transfer(msg.value);
        hasRaised += tokenNo;
    }

    function SetLevelRate(uint newRate, uint8 level) public onlyOwner{
        if (level == 1){
            _rateLevel1 = newRate;
        }
        if (level == 2){
            _rateLevel2 = newRate;
        }
    }

    function remainingTokens() public view returns (uint256) {
        return _token.allowance(_adminWallet, address(this));
    }

    function ChangeAdmin(address payable wallet) public onlyOwner{
        require(wallet != _adminWallet);
        _adminWallet = wallet;
    }
}