package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// GetAllStrategies returns all strategies found in world state
func (s *SmartContract) GetAllStrategies(ctx contractapi.TransactionContextInterface) ([]*Strategy, error) {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all strategies in the chaincode namespace.
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var strategies []*Strategy
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var strategy Strategy
		err = json.Unmarshal(queryResponse.Value, &strategy)
		if err != nil {
			return nil, err
		}
		strategies = append(strategies, &strategy)
	}

	return strategies, nil
}

// 用 ID 检查策略是否存在
func (s *SmartContract) StrategyExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	strategyJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return strategyJSON != nil, nil
}

// ReadStrategy returns the strategy stored in the world state with given id.
func (s *SmartContract) ReadStrategy(ctx contractapi.TransactionContextInterface, id string) (*Strategy, error) {
	strategyJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if strategyJSON == nil {
		return nil, fmt.Errorf("the strategy %s does not exist", id)
	}

	var strategy Strategy
	err = json.Unmarshal(strategyJSON, &strategy)
	if err != nil {
		return nil, err
	}

	return &strategy, nil
}

func (s *SmartContract) ReadTrades(ctx contractapi.TransactionContextInterface, id string) (*Trades, error) {
	tradesJSON, err := ctx.GetStub().GetPrivateData("ProviderMSPPrivateCollection", id+"_trades")
	if err != nil {
		return nil, err
	}
	var trades Trades
	err = json.Unmarshal(tradesJSON, &trades)
	if err != nil {
		return nil, err
	}
	return &trades, nil
}
func (s *SmartContract) ReadPositions(ctx contractapi.TransactionContextInterface, id string) (*Positions, error) {
	positionsJSON, err := ctx.GetStub().GetPrivateData("ProviderMSPPrivateCollection", id+"_positions")
	if err != nil {
		return nil, err
	}
	var positions Positions
	err = json.Unmarshal(positionsJSON, &positions)
	if err != nil {
		return nil, err
	}
	return &positions, nil
}
