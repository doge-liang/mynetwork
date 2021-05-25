package list

import (
	"log"
	"mynetwork/chaincode/strategy/constants"
	"mynetwork/chaincode/strategy/ledgerapi"
	. "mynetwork/chaincode/strategy/model"
)

// type TradeListInterface interface {
// 	AddTrade(*Trade) error
// 	// GetTrade(string, string) (*Trade, error)
// 	GetTradesByStrategyID(string) ([]*Trade, error)
// 	// UpdateTrade(*Trade) error
// }

type TradeList struct {
	stateList ledgerapi.StateListInterface
}

func (tl *TradeList) AddTrade(trade *Trade) error {
	return tl.stateList.AddState(trade)
}

// func (tl *TradeList) GetTrade(strategyID string, id string) (*Trade, error) {
// 	t := new(Trade)

// 	err := tl.stateList.GetState(GetTradesKey(strategyID, id), t)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return t, nil
// }

// 根据策略ID 获取交易页
func (tl *TradeList) GetTradesByStrategyIDPage(strategyID string, bookmark string, pagesize int32) ([]*Trade, string, error) {
	iter, bookmark, err := tl.stateList.GetStateByPartialCompositeKeyWithPagination([]string{strategyID}, pagesize, bookmark)
	if err != nil {
		return nil, "", err
	}

	var ts []*Trade
	i := 0
	for iter.HasNext() {
		i++
		log.Print(i)
		response, err := iter.Next()
		if err != nil {
			return nil, "", err
		}

		var t Trade
		err = DeserializeTrade(response.Value, &t)
		if err != nil {
			return nil, "", err
		}
		ts = append(ts, &t)
	}

	return ts, bookmark, err
}

// 根据策略ID 获取交易列表
func (tl *TradeList) GetTradesByStrategyID(strategyID string) ([]*Trade, error) {
	iter, err := tl.stateList.GetStateByPartialCompositeKey([]string{strategyID})
	if err != nil {
		return nil, err
	}

	var ts []*Trade
	i := 0
	for iter.HasNext() {
		i++
		log.Print(i)
		response, err := iter.Next()
		if err != nil {
			return nil, err
		}

		var t Trade
		err = DeserializeTrade(response.Value, &t)
		if err != nil {
			return nil, err
		}
		ts = append(ts, &t)
	}

	return ts, err
}

func (tl *TradeList) DeleteTrades(strategyID string) error {
	ts, err := tl.GetTradesByStrategyID(strategyID)
	if err != nil {
		return err
	}

	for _, t := range ts {
		err := tl.stateList.DelState(GetTradesKey(strategyID, t.ID))
		if err != nil {
			return err
		}
	}

	return nil
}

func (tl *TradeList) DelTrade(t *Trade) error {
	return tl.stateList.DelState(GetTradesKey(t.StrategyID, t.ID))
}

func newTradeList(ctx TransactionContextInterface) *TradeList {
	stateList := new(ledgerapi.StateList)
	stateList.Ctx = ctx
	stateList.Name = "org.mynetwork." + constants.TRADES + "list"
	stateList.Deserialize = func(bytes []byte, state ledgerapi.StateInterface) error {
		return DeserializeTrade(bytes, state.(*Trade))
	}

	list := new(TradeList)
	list.stateList = stateList

	return list
}
