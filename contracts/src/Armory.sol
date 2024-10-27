// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "./interfaces/ISTETH.sol";

contract ArmoryEntry {
    event Deposited(address indexed depositer, uint256 amount);
    event WithdrawalRequested(uint256 indexed tokenId, address indexed withdrawer, uint256 amount);

    ISTETH public stETH = ISTETH(0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84);
    WithdrawalQueue public withdrawalQueue = WithdrawalQueue(0x889edC2eDab5f40e902b864aD4d7AdE8E412F9B1);
    address public owner;
    constructor() {
        owner = msg.sender;
    }
    
    struct PermitInput {
        uint256 value;
        uint256 deadline;
        uint8 v;
        bytes32 r;
        bytes32 s;
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can call this function");
        _;
    }

    // Request a withdrawal from the LST contract
    function requestWithdrawalsWithPermit(uint256 _amount, address withdrawer, PermitInput calldata _permit) public onlyOwner returns (uint256[] memory requestIds) {
        ids = withdrawalQueue.requestWithdrawalsWithPermit([_amount], address(this), _permit);
        require(ids.length == 1, "Only one withdrawal request at a time");
        emit WithdrawalRequested(ids[0], withdrawer, _amount);
    }

    // Send ETH in order to deposit
    // This will mint LST
    receive() external payable {
        stETH.submit{value: msg.value}(address(0));
        emit Deposited(msg.sender, msg.value);
    }
}

contract ArmoryExit {
    event Withdrawn(uint256 indexed tokenId, uint256 amount);

    constructor() {}
    receive() external payable {

    }
}