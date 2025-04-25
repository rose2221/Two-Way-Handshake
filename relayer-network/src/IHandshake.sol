// Handshake.sol
pragma solidity ^0.8.24;

interface IHandshake {
    // ───── Events ──────────────────────────
    event Syn(uint256 indexed sessionId, address indexed from);
    event Ack(uint256 indexed sessionId);
    event SynAck(uint256 indexed sessionId);

    // ───── Functions ───────────────────────
    function initiateHandshake(uint256 sessionId) external; // SYN
    function acknowledge(uint256 sessionId) external;       // ACK
    function confirm(uint256 sessionId) external;           // SYN-ACK
}
