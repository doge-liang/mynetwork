package main

import (
	"log"

	. "mynetwork/chaincode/strategy/contract"
	"mynetwork/chaincode/strategy/model/list"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	contract := new(SmartContract)
	contract.TransactionContextHandler = new(list.TransactionContext)
	contract.Name = "org.mynetwork.strategy"

	strategyChaincode, err := contractapi.NewChaincode(contract)

	if err != nil {
		log.Panicf("Error creating strategy-subscribe chaincode: %v", err)
	}

	if err := strategyChaincode.Start(); err != nil {
		log.Panicf("Error starting strategy-subscribe chaincode: %v", err)
	}
}
