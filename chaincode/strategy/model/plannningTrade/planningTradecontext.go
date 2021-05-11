package planningTrade

import "github.com/hyperledger/fabric-contract-api-go/contractapi"

type TransactionContextInterface interface {
	contractapi.TransactionContextInterface
	GetStrategyList() PlanningTradeListInterface
}

type TransactionContext struct {
	contractapi.TransactionContext
	planningTradeList *PlanningTradeList
}

func (tc *TransactionContext) GetStrategyList() PlanningTradeListInterface {
	if tc.planningTradeList == nil {
		tc.planningTradeList = newPlanningTradeList(tc)
	}

	return tc.planningTradeList
}
