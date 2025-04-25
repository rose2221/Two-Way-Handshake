// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "forge-std/Script.sol";
import "../src/ContractA.sol";

contract DeployA is Script {
    /**
     * env vars used
     *  - PRIVATE_KEY              → EOA that broadcasts the tx
     *  - RELAYER_ADDR             → same relayer key as on chain-B
     *  - CONTRACT_B_ADDR          → address printed by DeployB
     */
    function run() external {
        uint256 pk       = vm.envUint("PRIVATE_KEY1");
        address relayer  = vm.envAddress("RELAYER_ADDR");
        address ctrB     = vm.envAddress("CONTRACT_B_ADDR");

        vm.startBroadcast(pk);

        ContractA a = new ContractA(ctrB, relayer);

        vm.stopBroadcast();

        console2.log("ContractA deployed:", address(a));
    }
}
