// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "forge-std/Script.sol";
import "../src/ContractA.sol";

contract DeployA is Script {
 
    function run() external {
        uint256 pk       = vm.envUint("PRIVATE_KEY1");
        address relayer  = vm.envAddress("RELAYER_ADDR");
        address ctrA     = vm.envAddress("CONTRACT_B_ADDR");

        vm.startBroadcast(pk);

        ContractA a = new ContractA(ctrA, relayer);

        vm.stopBroadcast();
    }
}
