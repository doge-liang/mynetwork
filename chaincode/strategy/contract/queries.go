package contract

import (
	"mynetwork/chaincode/strategy/constants"
	. "mynetwork/chaincode/strategy/model"
	. "mynetwork/chaincode/strategy/model/list"
	. "mynetwork/chaincode/strategy/utils"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract definition
type SmartContract struct {
	contractapi.Contract
}

type Subscription struct {
	StrategyID  string   `json:"strategyID"`  // 策略 ID
	Subscribers []string `json:"subscribers"` // 订阅者列表
}

// type PublicStrategy struct {
// 	ID   string `json:"ID"`   // 策略 ID
// 	Name string `json:"name"` // 策略名
// 	// Provider     string     `json:"provider"`     // 发布者
// 	MaxDrawdown  float64    `json:"maxDrawdown"`  // 最大回撤
// 	AnnualReturn float64    `json:"annualReturn"` // 年化收益率
// 	State        string     `json:"state"`        // 是否公开
// 	Trades       []Trade    `json:"trades"`       // 交易记录
// 	Positions    []Position `json:"positions"`    // 持仓记录
// }

// 获取所有策略
func (s *SmartContract) GetAllStrategies(ctx TransactionContextInterface) ([]*Strategy, error) {

	strats, err := ctx.GetStrategyList().GetAllStrategies()
	if err != nil {
		return nil, err
	}

	return strats, nil
}

// 通过策略 ID 读取策略交易记录
func (s *SmartContract) GetTradesByStrategyID(ctx TransactionContextInterface, strategyID string) ([]*Trade, error) {
	// MSPID, err := ctx.GetClientIdentity().GetMSPID()
	// if err != nil {
	// 	return nil, err
	// }

	trades, err := ctx.GetTradeList().GetTradesByStrategyID(strategyID)
	if err != nil {
		return nil, err
	}

	return trades, nil
	// return nil, fmt.Errorf("unknown MSPID: %s", orgID)
}

// 根据策略 ID 获取持仓信息
func (s *SmartContract) GetPositionsByStrategyID(ctx TransactionContextInterface, strategyID string) (PositionsOutput, error) {
	strat, err := ctx.GetStrategyList().GetStrategy(strategyID)
	if err != nil {
		return PositionsOutput{}, err
	}

	// 私有策略
	if strat.IsPrivate() {
		isSubed, err := s.IsSubscirbed(ctx, strategyID)
		if err != nil {
			return PositionsOutput{}, err
		}
		// 未订阅
		if !isSubed {
			phlist, err := ctx.GetPrivatePositionList(constants.PRIVATE_COLLECTION).GetPositionsHashByStrategyID(strategyID)
			if err != nil {
				return PositionsOutput{}, err
			}

			var output PositionsOutput
			output.PositionsHash = phlist

			return output, nil
		}

		// 已订阅
		ptlist, err := ctx.GetPrivatePositionList(constants.PUBLIC_COLLECTION).GetPositionsByStrategyIDInPrivate(strategyID)
		if err != nil {
			return PositionsOutput{}, err
		}

		var output PositionsOutput
		output.Positions = ptlist

		return output, nil
	}
	plist, err := ctx.GetPositionList().GetPositionsByStrategyID(strategyID)
	if err != nil {
		return PositionsOutput{}, err
	}

	var output PositionsOutput
	output.Positions = plist

	return output, nil
}

func (s *SmartContract) GetPlanningTradesByStrategyID(ctx TransactionContextInterface, strategyID string) (PlanningTradesOutput, error) {
	strat, err := ctx.GetStrategyList().GetStrategy(strategyID)
	if err != nil {
		return PlanningTradesOutput{}, err
	}

	// 私有策略
	if strat.IsPrivate() {
		isSubed, err := s.IsSubscirbed(ctx, strategyID)
		if err != nil {
			return PlanningTradesOutput{}, err
		}

		// 未订阅
		if !isSubed {
			pthlist, err := ctx.GetPrivatePlanningTradeList(constants.PRIVATE_COLLECTION).GetPlanningTradesHashByStrategyID(strategyID)
			if err != nil {
				return PlanningTradesOutput{}, err
			}

			var output PlanningTradesOutput
			output.PlanningTradesHash = pthlist

			return output, nil
		}
		// 已订阅
		ptlist, err := ctx.GetPrivatePlanningTradeList(constants.PUBLIC_COLLECTION).GetPlanningTradesByStrategyIDInPrivate(strategyID)
		if err != nil {
			return PlanningTradesOutput{}, err
		}

		var output PlanningTradesOutput
		output.PlanningTrades = ptlist

		return output, nil

	}

	ptlist, err := ctx.GetPlanningTradeList().GetPlanningTradesByStrategyID(strategyID)
	if err != nil {
		return PlanningTradesOutput{}, err
	}

	var output PlanningTradesOutput
	output.PlanningTrades = ptlist

	return output, nil
}

// 用 ID 检查策略是否存在
// func (s *SmartContract) StrategyExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
// 	strategyJSON, err := ctx.GetStub().GetState(id)
// 	if err != nil {
// 		return false, fmt.Errorf("failed to read from world state: %v", err)
// 	}

// 	return strategyJSON != nil, nil
// }

// 读哈希
// func (s *SmartContract) ReadPrivateStrategy(ctx contractapi.TransactionContextInterface, id string) (*PrivateStrategy, error) {
// 	key := GetStrategyKey(id)
// 	strategyJSON, err := ctx.GetStub().GetState(key)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var strategy Strategy
// 	err = json.Unmarshal(strategyJSON, &strategy)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// TradesHash, err := s.GetTradesHash(ctx, PRIVATE_COLLECTION, id)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }
// 	PositionsHash, err := s.GetPositionsHash(ctx, PRIVATE_COLLECTION, id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	PlanningTradesHash, err := s.GetPlanningTradesHash(ctx, PRIVATE_COLLECTION, id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	privateStrategy := PrivateStrategy{
// 		Strategy:       &strategy,
// 		Trades:         []*Trade{},
// 		Positions:      PositionsHash,
// 		PlanningTrades: PlanningTradesHash,
// 	}
// 	return &privateStrategy, nil
// }

// 通过策略的 id 得到交易数据
// func (s *SmartContract) ReadTrades(ctx contractapi.TransactionContextInterface, id string) ([]Trade, error) {
// 	tradesJSON, err := ctx.GetStub().GetPrivateData(PRIVATE_COLLECTION, GetTradesKey(id))
// 	if err != nil {
// 		return nil, err
// 	}
// 	var trades Trades
// 	err = json.Unmarshal(tradesJSON, &trades)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return trades.Trades, nil
// }

// 通过策略的 id 获得仓位数据
// func (s *SmartContract) ReadPositions(ctx contractapi.TransactionContextInterface, id string) ([]Position, error) {
// 	positionsJSON, err := ctx.GetStub().GetPrivateData(PRIVATE_COLLECTION, GetPositionsKey(id))
// 	if err != nil {
// 		return nil, err
// 	}
// 	var positions Positions
// 	err = json.Unmarshal(positionsJSON, &positions)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return positions.Positions, nil
// }

// 订阅者的合约

// 根据 Strategy 的 ID 读取某条策略
// 如果是公有的直接 GetState(key) 即可
// 如果是私有的则需要使用 PrivateStrategy 结构来返回数据，把哈希值填私有的数据部分里面
// func (s *SmartContract) sub_ReadStrategy(ctx contractapi.TransactionContextInterface, strategy *Strategy) (*Strategy, error) {

// 	clientID, err := ctx.GetClientIdentity().GetID()
// 	if err != nil {
// 		return nil, err
// 	}

// 	// 私有策略
// 	// 未订阅，无法访问该链码，应通过 ReadPrivateStrategy() 查询
// 	if !utils.in(clientID, strategy.Subscribers) {
// 		return nil, fmt.Errorf("Your have not subscribe this strategy yet.")
// 	}
// 	// 已订阅，检查发布者公布的数据和私有数据的 hash 是否一致，一致的才公布
// 	// pubilcTradesHash, _ := s.GetTradesHash(ctx, PUBLIC_COLLECTION, strategy.ID)
// 	pubilcPlanningTradesHash, _ := s.GetPlanningTradesHash(ctx, PUBLIC_COLLECTION, strategy.ID)
// 	publicPositionsHash, _ := s.GetPositionsHash(ctx, PUBLIC_COLLECTION, strategy.ID)
// 	// privateTradesHash, _ := s.GetTradesHash(ctx, PRIVATE_COLLECTION, strategy.ID)
// 	privatePlanningTradesHash, _ := s.GetPlanningTradesHash(ctx, PRIVATE_COLLECTION, strategy.ID)
// 	privatePositionsHash, _ := s.GetPositionsHash(ctx, PRIVATE_COLLECTION, strategy.ID)
// 	if publicPositionsHash == privatePositionsHash &&
// 		pubilcPlanningTradesHash == privatePlanningTradesHash {
// 		tradesJSON, err := ctx.GetStub().GetPrivateData(PUBLIC_COLLECTION, GetTradesKey(strategy.ID))
// 		if err != nil {
// 			return nil, err
// 		}
// 		planningTradesJSON, err := ctx.GetStub().GetPrivateData(PUBLIC_COLLECTION, GetTradesKey(strategy.ID))
// 		if err != nil {
// 			return nil, err
// 		}
// 		positionsJSON, err := ctx.GetStub().GetPrivateData(PUBLIC_COLLECTION, GetPositionsKey(strategy.ID))
// 		if err != nil {
// 			return nil, err
// 		}
// 		if tradesJSON != nil {
// 			var trades Trades
// 			err = json.Unmarshal(tradesJSON, &trades)
// 			strategy.Trades = trades.Trades
// 			if err != nil {
// 				return nil, err
// 			}
// 		}
// 		if planningTradesJSON != nil {
// 			var trades Trades
// 			err = json.Unmarshal(tradesJSON, &trades)
// 			strategy.Trades = trades.Trades
// 			if err != nil {
// 				return nil, err
// 			}
// 		}
// 		if positionsJSON != nil {
// 			var positions Positions
// 			err = json.Unmarshal(positionsJSON, &positions)
// 			strategy.Positions = positions.Positions
// 			if err != nil {
// 				return nil, err
// 			}
// 		}

// 		return strategy, nil
// 	}

// 	return nil, fmt.Errorf("Strategy verify fail")
// }

// func (s *SmartContract) pro_ReadStrategy(ctx contractapi.TransactionContextInterface, strategy *Strategy) (*Strategy, error) {
// 	clientID, err := ctx.GetClientIdentity().GetID()
// 	if err != nil {
// 		return nil, err
// 	}
// 	// 是否是该策略的发布者
// 	if strategy.Provider == clientID {
// 		tradesKey := GetTradesKey(strategy.ID)
// 		trades, _ := s.ReadTrades(ctx, tradesKey)
// 		strategy.Trades = trades
// 		positionsKey := GetPositionsKey(strategy.ID)
// 		positions, _ := s.ReadPositions(ctx, positionsKey)
// 		strategy.Positions = positions
// 		return strategy, nil
// 	}

// 	return nil, fmt.Errorf("You are not permitted to access.")
// }

// func (s *SmartContract) GetTradesHash(ctx contractapi.TransactionContextInterface, collection string, id string) (string, error) {
// 	tradesKey := GetTradesKey(id)
// 	tradesHash, err := ctx.GetStub().GetPrivateDataHash(collection, tradesKey)
// 	if err != nil {
// 		return "", err
// 	}
// 	return fmt.Sprintf("%x", tradesHash), nil
// }

// func (s *SmartContract) GetPositionsHash(ctx contractapi.TransactionContextInterface, collection string, id string) (string, error) {
// 	positionsKey := GetPositionsKey(id)
// 	positionsHash, err := ctx.GetStub().GetPrivateDataHash(collection, positionsKey)
// 	if err != nil {
// 		return "", err
// 	}
// 	return fmt.Sprintf("%x", positionsHash), nil
// }

// func (s *SmartContract) GetPlanningTradesHash(ctx contractapi.TransactionContextInterface, collection string, id string) (string, error) {
// 	planningTradesKey := GetPlanningTradesKey(id)
// 	planningTradesHash, err := ctx.GetStub().GetPrivateDataHash(collection, planningTradesKey)
// 	if err != nil {
// 		return "", err
// 	}
// 	return fmt.Sprintf("%x", planningTradesHash), nil
// }

func (s *SmartContract) IsSubscirbed(ctx TransactionContextInterface, strategyID string) (bool, error) {
	strat, err := ctx.GetStrategyList().GetStrategy(strategyID)
	if err != nil {
		return false, err
	}

	clientID, err := ctx.GetClientIdentity().GetID()
	if err != nil {
		return false, err
	}

	return In(clientID, strat.Subscribers), nil
}
