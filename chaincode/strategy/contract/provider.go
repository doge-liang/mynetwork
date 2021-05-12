package contract

import (
	"fmt"
	. "mynetwork/chaincode/strategy/model"
	. "mynetwork/chaincode/strategy/model/list"
)

// 发布策略
func (s *SmartContract) Distribute(ctx TransactionContextInterface, strategy *Strategy) error {

	MSPID, _ := ctx.GetClientIdentity().GetMSPID()
	if MSPID != "ProviderMSP" {
		return fmt.Errorf("You are not in Provider Org.")
	}

	clientID, err := ctx.GetClientIdentity().GetID()
	if err != nil {
		return err
	}

	strategy.Provider = clientID

	err = ctx.GetStrategyList().AddStrategy(strategy)
	if err != nil {
		return err
	}

	return nil
}

// 删除策略
func (s *SmartContract) DeleteStrategy(ctx TransactionContextInterface, strategyID string) error {

	MSPID, _ := ctx.GetClientIdentity().GetMSPID()

	if MSPID != "ProviderMSP" {
		return fmt.Errorf("You are not in Provider Org.")
	}

	strat, err := ctx.GetStrategyList().GetStrategy(strategyID)
	if err != nil {
		return err
	}

	if strat == nil {
		return fmt.Errorf("the strategy %s does not exist", GetStrategyKey(strategyID))
	}

	// 从公共的 state 中删除
	err = ctx.GetStrategyList().DeleteStrategy(strategyID)
	if strat.IsPrivate() {
		err := ctx.GetTradeList().DeleteTrades(strategyID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *SmartContract) Update(ctx TransactionContextInterface, strategyID string, ts []*Trade, pts []*PlanningTrade, ps []*Position) error {
	MSPID, _ := ctx.GetClientIdentity().GetMSPID()
	if MSPID != "ProviderMSP" {
		return fmt.Errorf("You are not in Provider Org.")
	}

	for _, t := range ts {
		err := ctx.GetTradeList().AddTrade(t)
		if err != nil {
			return err
		}
	}

	strat, err := ctx.GetStrategyList().GetStrategy(strategyID)
	if err != nil {
		return err
	}

	if strat.IsPrivate() {
		return nil
	}

	err = ctx.GetPlanningTradeList().UpdatePlanningTrades(pts, strategyID)
	if err != nil {
		return err
	}

	err = ctx.GetPositionList().UpdatePositions(ps, strategyID)
	if err != nil {
		return err
	}

	return nil
}

// 更新策略
// func (s *SmartContract) UpdateStrategy(ctx contractapi.TransactionContextInterfaceInterfaceInterface, strategy *Strategy) error {

// 	mspId, _ := ctx.GetClientIdentity().GetMSPID()

// 	if mspId != "Provider" {
// 		return fmt.Errorf("You are not in Provider Org.")
// 	}

// 	key := GetStrategyKey(strategy.ID)
// 	exist, err := s.StrategyExists(ctx, key)
// 	if err != nil {
// 		return err
// 	}
// 	if !exist {
// 		return fmt.Errorf("the strategy %s does not exist", strategy.ID)
// 	}

// 	return s.SaveStrategy(ctx, strategy)
// }

// 将策略状态改为公共
// func (s *SmartContract) SetStrategyPublic(ctx contractapi.TransactionContextInterfaceInterfaceInterface, id string) error {

// 	mspId, _ := ctx.GetClientIdentity().GetMSPID()

// 	if mspId != "Provider" {
// 		return fmt.Errorf("You are not in Provider Org.")
// 	}

// 	trades, err := s.ReadTrades(ctx, id)
// 	if err != nil {
// 		return err
// 	}
// 	positions, err := s.ReadPositions(ctx, id)
// 	if err != nil {
// 		return err
// 	}

// 	// 移除私有数据
// 	strategy, err := s.ReadStrategy(ctx, id)
// 	if err != nil {
// 		return err
// 	}
// 	// 添加私有数据并修改状态
// 	strategy.Trades = trades
// 	strategy.Positions = positions
// 	strategy.State = "public"
// 	return s.SaveStrategy(ctx, strategy)
// }

// 将策略状态改为私有
// func (s *SmartContract) SetStrategyPrivate(ctx contractapi.TransactionContextInterfaceInterfaceInterface, id string) error {

// 	mspId, _ := ctx.GetClientIdentity().GetMSPID()

// 	if mspId != "Provider" {
// 		return fmt.Errorf("You are not in Provider Org.")
// 	}

// 	strategy, err := s.ReadStrategy(ctx, id)
// 	if err != nil {
// 		return err
// 	}
// 	strategy.State = "private"
// 	return s.SaveStrategy(ctx, strategy)
// }

// func (s *SmartContract) DeleteTrades(ctx contractapi.TransactionContextInterfaceInterfaceInterface, id string) error {

// 	mspId, _ := ctx.GetClientIdentity().GetMSPID()

// 	if mspId != "Provider" {
// 		return fmt.Errorf("You are not in Provider Org.")
// 	}

// 	key := GetTradesKey(id)
// 	return ctx.GetStub().DelPrivateData(PRIVATE_COLLECTION, key)
// }

// func (s *SmartContract) DeletePositions(ctx contractapi.TransactionContextInterfaceInterfaceInterface, id string) error {

// 	mspId, _ := ctx.GetClientIdentity().GetMSPID()

// 	if mspId != "Provider" {
// 		return fmt.Errorf("You are not in Provider Org.")
// 	}

// 	key := GetPositionsKey(id)
// 	return ctx.GetStub().DelPrivateData(PRIVATE_COLLECTION, key)
// }
