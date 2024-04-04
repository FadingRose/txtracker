// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

contract FunctionTest {
    mapping(address => uint256) public balance;
    address public owner;
    uint256 public credits = 100;

    constructor() {
        owner = msg.sender;
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can call this function");
        _;
    }

    function deposit(uint256 amount) public {
        balance[msg.sender] += amount;
    }

    function withdraw(uint256 amount) public {
        require(balance[msg.sender] >= amount, "Insufficient balance");
        balance[msg.sender] -= amount;
    }
}
