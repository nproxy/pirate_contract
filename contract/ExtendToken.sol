pragma solidity >=0.5.11;

import "./owned.sol";
import "./safemath.sol";
import "./IERC20.sol";

contract ExtendToken is owned{

    struct Token{
        address admin;
        address paymentContract;
        IERC20  tokenI;
        bytes8  symbol;
        uint    decimal;
        uint index;
    }

    address[] public AllTokenList;
    address[] public AllPayContractList;
    mapping(address=>Token) public AllTokenData;

    constructor() public{
    }

    function AddNewToken(address admin, address payContract, address tokenAddr, uint8 decimal,
        bytes8 symbol) public onlyOwner{

        require(admin != address(0));
        require(tokenAddr != address(0));
        require(decimal <= 18 && decimal > 0);
        require(AllTokenData[tokenAddr].admin == address(0));
        IERC20 token = IERC20(tokenAddr);
        require(token.totalSupply() > 0 );

        AllTokenData[tokenAddr] = Token(admin, payContract, token, symbol, decimal, AllTokenList.length );
        AllTokenList.push(tokenAddr);
        AllPayContractList.push(payContract);
    }

    function DelToken(address tokenAddr) public onlyOwner{
        Token memory t = AllTokenData[tokenAddr];
        require(t.admin != address(0));

        delete AllTokenList[t.index];
        delete AllTokenData[tokenAddr];
    }


    function TokenBalances(address userAddr) public view returns(address[] memory, uint[] memory){
        uint token_len = AllTokenList.length;
        uint[] memory balance = new uint[](token_len);
        address[] memory addr = new address[](token_len);
        for (uint i = 0; i < AllTokenList.length; i++){
            address token_addr = AllTokenList[i];
            Token memory token_info = AllTokenData[token_addr];
            balance[i] = token_info.tokenI.balanceOf(userAddr);
            addr[i] = token_addr;
        }

        return (addr, balance);
    }

    function TokenBalance(address userAddr, address tokenAddr) public view returns (uint){
        return IERC20(tokenAddr).balanceOf(userAddr);
    }

    function TokensList() public view returns(address[] memory){
        return AllTokenList;
    }

    function PayContractsList() public view returns(address[] memory){
        return AllPayContractList;
    }
}