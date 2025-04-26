// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {Ownable} from "openzeppelin-contracts/contracts/access/Ownable.sol";


interface IContractA {
    function acknowledge(uint256 sessionId) external;
}

contract ContractB is Ownable {
 
    event Ack   (uint256 indexed sessionId);        
    event Done  (uint256 indexed sessionId);         

    error InvalidState();
    error OnlyRelayer();

    enum Status { None, AckSent, Complete }

    struct Session {
        Status status;
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

    function acknowledge(uint256 sessionId)
        external
       
        onlyRelayer
    {
        Session storage s = sessions[sessionId];
        if (s.status != Status.None) revert InvalidState();

        s.status = Status.AckSent;
        emit Ack(sessionId);
    }

 
    function confirm(uint256 sessionId)
        external
       
        onlyRelayer
    {
        Session storage s = sessions[sessionId];
        if (s.status != Status.AckSent) revert InvalidState();

        s.status = Status.Complete;
        emit Done(sessionId); 
    }

    function handshakeComplete(uint256 sessionId) external view returns (bool) {
        return sessions[sessionId].status == Status.Complete;
    }
}
