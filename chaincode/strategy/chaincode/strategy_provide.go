package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// 保存策略
func (s *SmartContract) SaveStrategy(ctx contractapi.TransactionContextInterface, strategy *Strategy) error {
	if strategy.State == "private" {
		strategy.ID = GetStrategyKey(strategy.ID)
		positions := Positions{
			StrategyID: strategy.ID,
			Positions:  strategy.Positions,
		}
		trades := Trades{
			StrategyID:     strategy.ID,
			Trades:         strategy.Trades,
			PlanningTrades: strategy.PlanningTrades,
		}

		positionsJSON, err := json.Marshal(positions)
		if err != nil {
			return err
		}
		tradesJSON, err := json.Marshal(trades)
		if err != nil {
			return err
		}
		positionsKey := GetPositionsKey(strategy.ID)
		err = ctx.GetStub().PutPrivateData(PRIVATE_COLLECTION, positionsKey, positionsJSON)
		if err != nil {
			return err
		}
		err = ctx.GetStub().PutPrivateData(PUBLIC_COLLECTION, positionsKey, positionsJSON)

		tradesKey := GetTradesKey(strategy.ID)
		err = ctx.GetStub().PutPrivateData(PRIVATE_COLLECTION, tradesKey, tradesJSON)
		if err != nil {
			return err
		}
		err = ctx.GetStub().PutPrivateData(PUBLIC_COLLECTION, tradesKey, tradesJSON)
		if err != nil {
			return err
		}

		// 清空公共部分
		strategy.Trades = []Trade{}
		strategy.PlanningTrades = []PlanningTrade{}
		strategy.Positions = []Position{}

		// err = ctx.GetStub().PutState(STRATEGY_COUNT, []byte(strconv.Itoa(strategyCount+1)))

		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	clientID, err := ctx.GetClientIdentity().GetID()
	if err != nil {
		return err
	}
	strategy.Provider = clientID
	strategyJSON, err := json.Marshal(strategy)
	if err != nil {
		return err
	}
	key := GetStrategyKey(strategy.ID)
	strategy.ID = key
	err = ctx.GetStub().PutState(key, strategyJSON)
	// err = ctx.GetStub().PutState(STRATEGY_COUNT, []byte(strconv.Itoa(strategyCount+1)))

	if err != nil {
		return fmt.Errorf("failed to put to world state. %v", err)
	}
	return nil

}

// 更新策略
func (s *SmartContract) UpdateStrategy(ctx contractapi.TransactionContextInterface, strategy *Strategy) error {
	key := GetStrategyKey(strategy.ID)
	exist, err := s.StrategyExists(ctx, key)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("the strategy %s does not exist", strategy.ID)
	}

	return s.SaveStrategy(ctx, strategy)
}

// 将策略状态改为公共
func (s *SmartContract) SetStrategyPublic(ctx contractapi.TransactionContextInterface, id string) error {
	trades, err := s.ReadTrades(ctx, id)
	if err != nil {
		return err
	}
	positions, err := s.ReadPositions(ctx, id)
	if err != nil {
		return err
	}

	// 移除私有数据
	strategy, err := s.ReadStrategy(ctx, id)
	if err != nil {
		return err
	}
	// 添加私有数据并修改状态
	strategy.Trades = trades
	strategy.Positions = positions
	strategy.State = "public"
	return s.SaveStrategy(ctx, strategy)
}

// 将策略状态改为私有
func (s *SmartContract) SetStrategyPrivate(ctx contractapi.TransactionContextInterface, id string) error {
	strategy, err := s.ReadStrategy(ctx, id)
	if err != nil {
		return err
	}
	strategy.State = "private"
	return s.SaveStrategy(ctx, strategy)
}

func (s *SmartContract) DeleteTrades(ctx contractapi.TransactionContextInterface, id string) error {
	key := GetTradesKey(id)
	return ctx.GetStub().DelPrivateData(PRIVATE_COLLECTION, key)
}

func (s *SmartContract) DeletePositions(ctx contractapi.TransactionContextInterface, id string) error {
	key := GetPositionsKey(id)
	return ctx.GetStub().DelPrivateData(PRIVATE_COLLECTION, key)
}

func (s *SmartContract) DeleteStrategy(ctx contractapi.TransactionContextInterface, id string) error {
	key := GetStrategyKey(id)
	strategy, err := s.ReadStrategy(ctx, id)
	if err != nil {
		return err
	}
	if strategy == nil {
		return fmt.Errorf("the strategy %s does not exist", key)
	}
	// 从公共的 state 中删除
	err = ctx.GetStub().DelState(key)
	if strategy.State == "private" {
		s.DeleteTrades(ctx, id)
		s.DeletePositions(ctx, id)
	}
	return err
}
