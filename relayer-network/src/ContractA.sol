// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;


import {Ownable} from "openzeppelin-contracts/contracts/access/Ownable.sol";

interface IContractB {
    function confirm(uint256 sessionId) external;
}

contract ContractA is Ownable{

    event Syn    (uint256 indexed sessionId, address indexed initiator);
    event SynAck (uint256 indexed sessionId);

    error InvalidState();
    error OnlyRelayer();
    error SessionExists();

    enum Status { None, SynSent, AckReceived, Complete }

    struct Session {
        Status  status;
        address initiator;  
    }

    mapping(uint256 => Session) public sessions;  
    address public immutable counterparty;        
    address public relayer;                  


    constructor(address _counterparty, address _relayer)  Ownable(msg.sender){
        require(_counterparty != address(0) && _relayer != address(0), "zero");
        counterparty = _counterparty;
        relayer      = _relayer;
    }

    function setRelayer(address newRelayer) external onlyOwner {
        require(newRelayer != address(0), "zero");
        relayer = newRelayer;
    }

    modifier onlyRelayer() {
        if (msg.sender != relayer) revert OnlyRelayer();
        _;
    }

    function initiateHandshake(uint256 sessionId)
        external
       
    {
        if (sessions[sessionId].status != Status.None) revert SessionExists();

        sessions[sessionId] = Session({
            status:    Status.SynSent,
            initiator: msg.sender
        });

        emit Syn(sessionId, msg.sender);

    }

    function acknowledge(uint256 sessionId)
        external
     
        onlyRelayer
    {
        Session storage s = sessions[sessionId];
        if (s.status != Status.SynSent) revert InvalidState();

        s.status = Status.Complete; 

        emit SynAck(sessionId);
    }

    function handshakeComplete(uint256 sessionId) external view returns (bool) {
        return sessions[sessionId].status == Status.Complete;
    }
}

