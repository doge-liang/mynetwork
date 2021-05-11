package trade

import (
	"mynetwork/chaincode/strategy/consts"
	ledgerapi "mynetwork/chaincode/strategy/ledger-api"
)

type TradeListInterface interface {
	AddTrade(*Trade) error
	GetTrade(string, string) (*Trade, error)
	UpdateTrade(*Trade) error
}

type TradeList struct {
	stateList ledgerapi.StateListInterface
}

func (tl *TradeList) AddTrade(trade *Trade) error {
	return tl.stateList.AddState(trade)
}

func (tl *TradeList) GetTrade(strategyID string, id string) (*Trade, error) {
	t := new(Trade)

	err := tl.stateList.GetState(GetTradesKey(strategyID, id), t)

	if err != nil {
		return nil, err
	}

	return t, nil
}

func (tl *TradeList) UpdateTrade(t *Trade) error {
	return tl.stateList.UpdateState(t)
}

func newTradeList(ctx TransactionContextInterface) *TradeList {
	stateList := new(ledgerapi.StateList)
	stateList.Ctx = ctx
	stateList.Name = "org.mynetwork." + consts.TRADES + "list"
	stateList.Deserialize = func(bytes []byte, state ledgerapi.StateInterface) error {
		return DeserializeTrade(bytes, state.(*Trade))
	}

	list := new(TradeList)
	list.stateList = stateList

	return list
}
