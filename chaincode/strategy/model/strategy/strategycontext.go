package strategy

import "github.com/hyperledger/fabric-contract-api-go/contractapi"

type TransactionContextInterface interface {
	contractapi.TransactionContextInterface
	GetStrategyList() StrategyListInterface
}

type TransactionContext struct {
	contractapi.TransactionContext
	strategyList *StrategyList
}

func (tc *TransactionContext) GetStrategyList() StrategyListInterface {
	if tc.strategyList == nil {
		tc.strategyList = newStrategyList(tc)
	}

	return tc.strategyList
}
