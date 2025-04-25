// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

/*──────────────────────────────────────────────────────────────────────────────
│  ContractB — “responder side” of a cross-chain three-way handshake
│
│  Deployed on a different network from ContractA (e.g., Mumbai if A is Goerli)
│
│  Handshake timeline
│  ┌──────────────────────────────────────────────────────────────────────────┐
│  │ 1. ContractA emits Syn(sessionId)      ──▶ (relayer) acknowledge()  │B│ │
│  │ 2. ContractB.acknowledge() emits Ack   ◀── (relayer) Ack log        │ │ │
│  │ 3. ContractA.acknowledge() emits SynAck──▶ (relayer) confirm()   │B│ │
│  │ 4. ContractB.confirm() marks Complete and optionally emits Done        │
│  └──────────────────────────────────────────────────────────────────────────┘
│
│  •  Only the designated `relayer` may call acknowledge() & confirm().
│  •  `Ownable` allows the deployer (or a multisig) to rotate the relayer key.
│  •  ReentrancyGuard is included for completeness (no external calls here).
└─────────────────────────────────────────────────────────────────────────────*/

import {Ownable}          from "openzeppelin-contracts/contracts/access/Ownable.sol";


interface IContractA {
    function acknowledge(uint256 sessionId) external;
}

contract ContractB is Ownable {
    /*────────────────────────── Events ──────────────────────────*/
    event Ack   (uint256 indexed sessionId);          // step-2
    event Done  (uint256 indexed sessionId);          // optional final marker

    /*────────────────────────── Errors ──────────────────────────*/
    error InvalidState();
    error OnlyRelayer();

    /*──────────────────────── State/config ──────────────────────*/
    enum Status { None, AckSent, Complete }

    struct Session {
        Status status;
    }

    mapping(uint256 => Session) public sessions;  // sessionId → Session
    address public immutable counterparty;        // ContractA on the other chain
    address public relayer;                       // hot key that sends txs

    /*──────────────────────── Constructor ───────────────────────*/
    constructor(address _counterparty, address _relayer)  Ownable(msg.sender){
        require(_counterparty != address(0) && _relayer != address(0), "zero");
        counterparty = _counterparty;
        relayer      = _relayer;
    }

    /*────────────────── Admin: rotate relayer ───────────────────*/
    function setRelayer(address newRelayer) external onlyOwner {
        require(newRelayer != address(0), "zero");
        relayer = newRelayer;
    }

    modifier onlyRelayer() {
        if (msg.sender != relayer) revert OnlyRelayer();
        _;
    }

    /*──────────────────── 2️⃣  ACK (relayer) ────────────────────*/
    /// @dev Called by the relayer after it observes Syn on ContractA.
    ///      Emits `Ack`, which the relayer will pick up to drive step-3.
    function acknowledge(uint256 sessionId)
        external
       
        onlyRelayer
    {
        Session storage s = sessions[sessionId];
        if (s.status != Status.None) revert InvalidState();

        s.status = Status.AckSent;
        emit Ack(sessionId);
    }

    /*──────────────────── 4️⃣  confirm  (relayer) ───────────────*/
    /// @dev Finalises the handshake after ContractA emits SynAck.
    function confirm(uint256 sessionId)
        external
       
        onlyRelayer
    {
        Session storage s = sessions[sessionId];
        if (s.status != Status.AckSent) revert InvalidState();

        s.status = Status.Complete;
        emit Done(sessionId);   // optional, but nice for explorers/metrics
    }

    /*───────────────── View helper (optional) ───────────────────*/
    function handshakeComplete(uint256 sessionId) external view returns (bool) {
        return sessions[sessionId].status == Status.Complete;
    }
}
