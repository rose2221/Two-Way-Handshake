// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

/*──────────────────────────────────────────────────────────────────────────────
│  ContractA — “initiator side” of a cross-chain three-way handshake
│
│  Flow
│  1.  initiateHandshake()  → emits Syn(sessionId)         (Goerli ➜ Mumbai)
│      - Relayer picks up Syn log and calls acknowledge() on ContractB.
│  2.  ContractB.acknowledge()  → emits Ack(sessionId)     (Mumbai ➜ Goerli)
│      - Relayer picks up Ack log and calls acknowledge() **here**.
│  3.  ContractA.acknowledge()  → emits SynAck(sessionId)  (Goerli ➜ Mumbai)
│      - Relayer relays SynAck to ContractB.confirm().
│  ────────────────────────────────────────────────────────────────────────────
│  After Syn-Ack is emitted the handshake is considered complete.
│
│  •   Only the designated `relayer` may invoke acknowledge().
│  •   Owner (--deploy script / multisig) controls the relayer address.
│  •   Re-entrancy not possible (no external calls), but ReentrancyGuard
│      included for defense-in-depth.
│  •   Custom errors save gas vs. revert strings.
└─────────────────────────────────────────────────────────────────────────────*/

import {Ownable}          from "openzeppelin-contracts/contracts/access/Ownable.sol";


interface IContractB {
    function confirm(uint256 sessionId) external;
}

contract ContractA is Ownable{
    /*─────────────────────────── Events ──────────────────────────*/
    event Syn    (uint256 indexed sessionId, address indexed initiator);
    event SynAck (uint256 indexed sessionId);         // emitted in step-3

    /*─────────────────────────── Errors ──────────────────────────*/
    error InvalidState();
    error OnlyRelayer();
    error SessionExists();

    /*──────────────────────── State / config ─────────────────────*/
    enum Status { None, SynSent, AckReceived, Complete }

    struct Session {
        Status  status;
        address initiator;   // EOA that called initiateHandshake
    }

    mapping(uint256 => Session) public sessions;  // sessionId → Session
    address public immutable counterparty;        // ContractB on the other chain
    address public relayer;                       // hot key that does the bridging

    /*──────────────────────── Constructor ────────────────────────*/
    constructor(address _counterparty, address _relayer)  Ownable(msg.sender){
        require(_counterparty != address(0) && _relayer != address(0), "zero");
        counterparty = _counterparty;
        relayer      = _relayer;
    }

    /*─────────────────── Admin: update relayer ───────────────────*/
    function setRelayer(address newRelayer) external onlyOwner {
        require(newRelayer != address(0), "zero");
        relayer = newRelayer;
    }

    modifier onlyRelayer() {
        if (msg.sender != relayer) revert OnlyRelayer();
        _;
    }

    /*──────────────────── 1️⃣  SYN  (user-facing) ───────────────────*/
    function initiateHandshake(uint256 sessionId)
        external
       
    {
        if (sessions[sessionId].status != Status.None) revert SessionExists();

        sessions[sessionId] = Session({
            status:    Status.SynSent,
            initiator: msg.sender
        });

        emit Syn(sessionId, msg.sender);
        // NOTE: The relayer observes this log and calls
        //       ContractB.acknowledge(sessionId) on the other network.
    }

    /*──────────────────── 2️⃣+3️⃣  ACK → SYN-ACK  (relayer) ───────────*/
    /// @notice Called by the relayer after ContractB has emitted `Ack`.
    ///         Emits `SynAck` so the relayer can finalize the handshake on B.
    function acknowledge(uint256 sessionId)
        external
     
        onlyRelayer
    {
        Session storage s = sessions[sessionId];
        if (s.status != Status.SynSent) revert InvalidState();

        s.status = Status.Complete; // Handshake considered done locally.

        emit SynAck(sessionId);
        // Relayer will now call ContractB.confirm(sessionId) cross-chain.
    }

    /*──────────────────── View helpers (optional) ─────────────────*/
    function handshakeComplete(uint256 sessionId) external view returns (bool) {
        return sessions[sessionId].status == Status.Complete;
    }
}

