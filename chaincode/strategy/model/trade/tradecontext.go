package trade

import "github.com/hyperledger/fabric-contract-api-go/contractapi"

type TransactionContextInterface interface {
	contractapi.TransactionContextInterface
	GetTradeList() TradeListInterface
}

type TransactionContext struct {
	contractapi.TransactionContext
	tradeList *TradeList
}

func (tc *TransactionContext) GetTradeList() TradeListInterface {
	if tc.tradeList == nil {
		tc.tradeList = newTradeList(tc)
	}

	return tc.tradeList
}
