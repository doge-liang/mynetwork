package main

import (
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-contract-api-go/metadata"
)

// SmartContract definition
type MyContract struct {
	contractapi.Contract
}

func (mc *MyContract) Init(ctx contractapi.TransactionContextInterface) error {
	return ctx.GetStub().PutState("Name", []byte("Fabric@Golang"))
}

func (mc *MyContract) Hi(ctx contractapi.TransactionContextInterface) string {
	name, _ := ctx.GetStub().GetState("Name")
	return "Hello, I'm " + string(name)
}

func main() {
	assetContract := new(MyContract)
	assetContract.Name = "example01.MyContract"
	assetContract.Info = metadata.InfoMetadata{
		Title: "MyContract",
		Description: "SmartContract Example 01 - Blockchain Workshop",
		Version: "1.0.0",
		Contact: &metadata.ContactMetadata{
			Name: "Bing",
			Email: "23227732@qq.com",
		},
	}

	contractEngine, err := contractapi.NewChaincode(assetContract)
	if err != nil {
		fmt.Printf("Error creating Example01 chaincode: %s", err.Error())
		fmt.Println()
		return
	}
	contractEngine.Info = metadata.InfoMetadata{
		Title: "SmartContract Set Example 01",
		Version: "1.0.0",
	}

	err = contractEngine.Start()
	if err != nil {
		fmt.Printf("Error starting Example01 chaincode: %s", err.Error())
		fmt.Println()
	}
}
