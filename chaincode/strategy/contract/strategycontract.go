package contract

import (
	"fmt"
	"mynetwork/chaincode/strategy/consts"
	"mynetwork/chaincode/strategy/model/strategy"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// 保存策略
func (s *SmartContract) SaveStrategy(ctx strategy.TransactionContextInterface, strat *strategy.Strategy) error {
	mspId, _ := ctx.GetClientIdentity().GetMSPID()
	// 只有 Provider 才能写策略
	if mspId != "Provider" {
		return fmt.Errorf("you are not in provider org.")
	}

	clientID, err := ctx.GetClientIdentity().GetID()
	if err != nil {
		return fmt.Errorf("failed to get identity. %v", err)
	}
	strat.Provider = clientID
	err = ctx.GetStrategyList().AddStrategy(strat)

	if err != nil {
		return fmt.Errorf("failed to put to world state. %v", err)
	}
	return nil

}

// 更新策略
func (s *SmartContract) UpdateStrategy(ctx strategy.TransactionContextInterface, strat *strategy.Strategy) error {

	mspId, _ := ctx.GetClientIdentity().GetMSPID()
	// 只有 Provider 才能写策略
	if mspId != "Provider" {
		return fmt.Errorf("you are not in provider org.")
	}

	clientID, err := ctx.GetClientIdentity().GetID()

	if err != nil {
		return fmt.Errorf("failed to get identity. %v", err)
	}

	oldstrat, err := ctx.GetStrategyList().GetStrategy(strat.ID)

	if err != nil {
		return err
	}

	if oldstrat.Provider != clientID || oldstrat.Provider != strat.Provider {
		return fmt.Errorf("you have no permitted to update a strategy provided by other.")
	}

	err = ctx.GetStrategyList().UpdateStrategy(strat)

	if err != nil {
		return fmt.Errorf("failed to update to world state. %v", err)
	}

	return nil
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
	return ctx.GetStub().DelPrivateData(consts.PRIVATE_COLLECTION, key)
}

func (s *SmartContract) DeletePositions(ctx contractapi.TransactionContextInterface, id string) error {
	key := GetPositionsKey(id)
	return ctx.GetStub().DelPrivateData(consts.PRIVATE_COLLECTION, key)
}

func (s *SmartContract) DeleteStrategy(ctx strategy.TransactionContextInterface, id string) error {
	key := strategy.GetStrategyKey(id)
	strat, err := ctx.GetStrategyList().GetStrategy(key)

	if err != nil {
		return err
	}

	if strat == nil {
		return fmt.Errorf("the strategy %s does not exist", key)
	}
	// 从公共的 state 中删除
	if strat.IsPrivate() {
		s.DeleteTradeByStrategyID(ctx, id)
		s.DeletePositionByStrategyID(ctx, id)
	}
	err = ctx.GetStub().DelState(key)
	return err
}
