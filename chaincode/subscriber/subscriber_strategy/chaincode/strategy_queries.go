package chaincode

import (
	"encoding/json"
	"fmt"
	"sort"

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

// 根据 Strategy 的 ID 读取某条策略
// 如果是公有的直接 GetState(key) 即可
// 如果是私有的则需要使用 PrivateStrategy 结构来返回数据，把哈希值填私有的数据部分里面
func (s *SmartContract) ReadStrategy(ctx contractapi.TransactionContextInterface, id string) (*Strategy, error) {
	key := GetStrategyKey(id)
	strategyJSON, err := ctx.GetStub().GetState(key)
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

	clientID, err := ctx.GetClientIdentity().GetID()
	if err != nil {
		return nil, err
	}

	// 私有策略
	if strategy.State == "private" {
		// 未订阅，无法访问该链码，应通过 ReadPrivateStrategy() 查询
		if !in(clientID, strategy.Subscribers) {
			return nil, fmt.Errorf("Your have not subscribe this strategy yet.")
		}
		// TODO 已订阅，检查发布者公布的数据和私有数据的 hash 是否一致，一致的才公布
		pubilcTradesHash, _ := s.GetTradesHash(ctx, PUBLIC_COLLECTION, id)
		publicPositionsHash, _ := s.GetTradesHash(ctx, PUBLIC_COLLECTION, id)
		privateTradesHash, _ := s.GetPositionsHash(ctx, PRIVATE_COLLECTION, id)
		privatePositionsHash, _ := s.GetPositionsHash(ctx, PUBLIC_COLLECTION, id)
		if pubilcTradesHash == privateTradesHash && publicPositionsHash == privatePositionsHash {
			tradesJSON, err := ctx.GetStub().GetPrivateData(PUBLIC_COLLECTION, GetTradesKey(id))
			if err != nil {
				return nil, err
			}
			positionsJSON, err := ctx.GetStub().GetPrivateData(PUBLIC_COLLECTION, GetPositionsKey(id))
			if err != nil {
				return nil, err
			}
			if tradesJSON != nil {
				var trades Trades
				err = json.Unmarshal(tradesJSON, &trades)
				strategy.Trades = trades.Trades
			}
			if positionsJSON != nil {
				var positions Positions
				err = json.Unmarshal(positionsJSON, &positions)
				strategy.Positions = positions.Positions

			}

			return &strategy, err
		}
	}

	return &strategy, nil
}

func (s *SmartContract) GetTradesHash(ctx contractapi.TransactionContextInterface, collection string, id string) (string, error) {
	tradesKey := GetTradesKey(id)
	tradesHash, err := ctx.GetStub().GetPrivateDataHash(collection, tradesKey)
	if err != nil {
		return "", err
	}
	return string(tradesHash), nil
}

func (s *SmartContract) GetPositionsHash(ctx contractapi.TransactionContextInterface, collection string, id string) (string, error) {
	positionsKey := GetTradesKey(id)
	positionsHash, err := ctx.GetStub().GetPrivateDataHash(collection, positionsKey)
	if err != nil {
		return "", err
	}
	return string(positionsHash), nil
}

func (s *SmartContract) IsSubscirbed(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	strategy, err := s.ReadStrategy(ctx, id)
	if err != nil {
		return false, err
	}

	clientID, err := ctx.GetClientIdentity().GetID()
	if err != nil {
		return false, err
	}

	return in(clientID, strategy.Subscribers), nil
}

func in(target string, str_array []string) bool {
	sort.Strings(str_array)
	index := sort.SearchStrings(str_array, target)
	//index的取值：[0,len(str_array)]
	if index < len(str_array) && str_array[index] == target { //需要注意此处的判断，先判断 &&左侧的条件，如果不满足则结束此处判断，不会再进行右侧的判断
		return true
	}
	return false
}
