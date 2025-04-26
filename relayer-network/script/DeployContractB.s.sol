// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "forge-std/Script.sol";
import "../src/ContractB.sol";

contract DeployB is Script {

    function run() external {
        uint256 pk       = vm.envUint("PRIVATE_KEY2");
        address relayer  = vm.envAddress("RELAYER_ADDR");
        address ctrA     = vm.envAddress("CONTRACT_A_ADDR");

        vm.startBroadcast(pk);

        ContractB b = new ContractB(ctrA, relayer);

        vm.stopBroadcast();

    }
}
