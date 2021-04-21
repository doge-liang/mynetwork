package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (s *SmartContract) Subscribe(ctx contractapi.TransactionContextInterface, id string) error {
	key := GetStrategyKey(id)
	strategyJSON, err := ctx.GetStub().GetState(key)
	if err != nil {
		return err
	}
	if strategyJSON == nil {
		return fmt.Errorf("Strategy %s does not exist", key)
	}
	var strategy Strategy
	err = json.Unmarshal(strategyJSON, &strategy)
	if err != nil {
		return err
	}
	clientID, err := ctx.GetClientIdentity().GetID()
	if err != nil {
		return err
	}
	if clientID == "" {
		return fmt.Errorf("The ID is empty")
	}
	if in(clientID, strategy.Subscribers) {
		return fmt.Errorf("You have been subscribed.")
	}
	strategy.Subscribers = append(strategy.Subscribers, clientID)
	strategyJSON, err = json.Marshal(strategy)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(key, strategyJSON)
}
