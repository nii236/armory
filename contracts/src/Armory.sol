// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "./interfaces/ISTETH.sol";

contract ArmoryEntry {
    event Deposited(address indexed depositer, uint256 amount);
    event WithdrawalRequested(uint256 indexed tokenId, address indexed withdrawer, uint256 amount);

    ISTETH public stETH = ISTETH(0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84);
    WithdrawalQueue public withdrawalQueue = WithdrawalQueue(0x889edC2eDab5f40e902b864aD4d7AdE8E412F9B1);

    ArmoryBuffer public armoryBuffer;

    address public owner;
    constructor(
        address _armoryBuffer
    ) {
        owner = msg.sender;
        armoryBuffer = ArmoryBuffer(_armoryBuffer);
    }
    
    struct PermitInput {
        uint256 value;
        uint256 deadline;
        uint8 v;
        bytes32 r;Oh the GitHub minutes issue, looks like we had a few builds that
        bytes32 s;
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can call this function");
        _;
    }

    // Send ETH in order to deposit
    // This will mint LST
    receive() external payable {
        stETH.submit{value: msg.value}(address(0));
        // Send the stETH to the buffer
        stETH.transfer(address(armoryBuffer), stETH.balanceOf(address(this)));
        emit Deposited(msg.sender, msg.value);
    }
}

contract ArmoryBuffer {
    ArmoryExit public armoryExit 

    constructor(    ) {    }

    function setArmoryExit(address _armoryExit) public onlyOwner {
        armoryExit = ArmoryExit(_armoryExit);
    }

    modifier onlyArmoryExit() {
        require(msg.sender == address(armoryExit), "Only ArmoryExit can call this function");
        _;
    }

    // Request a withdrawal from the LST contract
    function requestWithdrawalsWithPermit(uint256 _amount, address withdrawer, PermitInput calldata _permit) public onlyOwner returns (uint256[] memory requestIds) {
        ids = withdrawalQueue.requestWithdrawalsWithPermit([_amount], address(this), _permit);
        require(ids.length == 1, "Only one withdrawal request at a time");
        emit WithdrawalRequested(ids[0], withdrawer, _amount);
    }

    function sendToExit(uint256 tokenId, uint256 amount) onlyArmoryExit {
        require(armoryExit != address(0), "ArmoryExit not set");
        // Send the stETH to the exit contract
        stETH.transfer(address(armoryExit), amount);
        emit Withdrawn(tokenId, amount);
    }
}
contract ArmoryExit {
    ArmoryBuffer public armoryBuffer;
    constructor( 
        address _armoryBuffer
    ) {
        armoryBuffer = ArmoryBuffer(_armoryBuffer);
    }

    event Withdrawn(uint256 indexed tokenId, uint256 amount);

    function withdraw(uint256 tokenId, uint256 amount) public {
        armoryBuffer.sendToExit(tokenId, amount);
    }
}