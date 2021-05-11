package position

import "github.com/hyperledger/fabric-contract-api-go/contractapi"

type TransactionContextInterface interface {
	contractapi.TransactionContextInterface
	GetPositionList() PositionListInterface
}

type TransactionContext struct {
	contractapi.TransactionContext
	positionsList *PositionList
}

func (tc *TransactionContext) GetPositionList() PositionListInterface {
	if tc.positionsList == nil {
		tc.positionsList = newPositionList(tc)
	}

	return tc.positionsList
}
