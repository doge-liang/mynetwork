package list

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type TransactionContextInterface interface {
	contractapi.TransactionContextInterface

	GetStrategyList() *StrategyList

	GetTradeList() *TradeList

	GetPlanningTradeList() *PlanningTradeList
	GetPrivatePlanningTradeList(string) *PlanningTradeList

	GetPositionList() *PositionList
	GetPrivatePositionList(string) *PositionList
}

type TransactionContext struct {
	contractapi.TransactionContext

	strategyList      *StrategyList
	tradeList         *TradeList
	positionList      *PositionList
	PlanningTradeList *PlanningTradeList
}

func (tc *TransactionContext) GetStrategyList() *StrategyList {
	if tc.strategyList == nil {
		tc.strategyList = newStrategyList(tc)
	}

	return tc.strategyList
}

func (tc *TransactionContext) GetTradeList() *TradeList {
	if tc.tradeList == nil {
		tc.tradeList = newTradeList(tc)
	}

	return tc.tradeList
}

func (tc *TransactionContext) GetPlanningTradeList() *PlanningTradeList {
	if tc.PlanningTradeList == nil {
		tc.PlanningTradeList = newPlanningTradeList(tc)
	}

	return tc.PlanningTradeList
}

func (tc *TransactionContext) GetPrivatePlanningTradeList(collection string) *PlanningTradeList {
	if tc.positionList == nil {
		tc.positionList = newPrivatePositionList(tc, collection)
	}

	return tc.PlanningTradeList
}

func (tc *TransactionContext) GetPositionList() *PositionList {
	if tc.positionList == nil {
		tc.positionList = newPositionList(tc)
	}

	return tc.positionList
}

func (tc *TransactionContext) GetPrivatePositionList(collection string) *PositionList {
	if tc.positionList == nil {
		tc.positionList = newPrivatePositionList(tc, collection)
	}

	return tc.positionList
}
