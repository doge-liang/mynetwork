package main

import (
	"log"
	"mynetwork/chaincode/subscriber/strategy/chaincode"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	strategyChaincode, err := contractapi.NewChaincode(&chaincode.SmartContract{})
	if err != nil {
		log.Panicf("Error creating strategy-subscribe chaincode: %v", err)
	}

	if err := strategyChaincode.Start(); err != nil {
		log.Panicf("Error starting strategy-subscribe chaincode: %v", err)
	}
}
