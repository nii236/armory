// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Test, console} from "forge-std/Test.sol";
import {Armory} from "src/Armory.sol";
import "src/interfaces/ISTETH.sol";


contract ArmoryTest is Test {
    function test() public {
        address alice = address(1);
        deal(alice, 1 ether);
        Armory armory = new Armory();
        ISTETH stETH = ISTETH(0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84);


        // alice sends 1 ether to armory
        vm.prank(alice);
        (bool success, ) = payable(address(armory)).call{value: 1 ether}("");
        require(success, "Transfer failed");
        vm.stopPrank();

      // Check that the Armory received 1 stETH
        uint256 armoryBalance = stETH.balanceOf(address(armory));
        assertApproxEqAbs(armoryBalance, 1 ether, 2);

        vm.warp(block.timestamp + 1 days);
        
        // bob rebases the stETH
        address bob = address(2);
        deal(bob, 1 ether);
        stETH.submit{value: 1 ether}(address(0));

        // armory balance increases
        uint256 armoryBalance2 = stETH.balanceOf(address(armory));
        assertGt(armoryBalance2, armoryBalance);
    }
}
