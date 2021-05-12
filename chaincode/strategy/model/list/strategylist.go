package list

import (
	"encoding/json"
	"mynetwork/chaincode/strategy/constants"
	"mynetwork/chaincode/strategy/ledgerapi"
	. "mynetwork/chaincode/strategy/model"
)

// type StrategyListInterface interface {
// 	AddStrategy(*Strategy) error
// 	GetStrategy(string) (*Strategy, error)
// 	UpdateStrategy(*Strategy) error
// 	GetAllStrategies() ([]*Strategy, error)
// }

type StrategyList struct {
	stateList ledgerapi.StateListInterface
}

// 添加策略到公开账本
func (sl *StrategyList) AddStrategy(strategy *Strategy) error {
	return sl.stateList.AddState(strategy)
}

// 从公开账本查询策略
func (sl *StrategyList) GetStrategy(id string) (*Strategy, error) {
	strat := new(Strategy)

	err := sl.stateList.GetState(GetStrategyKey(id), strat)

	return strat, err
}

// 查询所有策略
func (sl *StrategyList) GetAllStrategies() ([]*Strategy, error) {
	iter, err := sl.stateList.GetStateByPartialCompositeKey([]string{})
	if err != nil {
		return nil, err
	}

	var strats []*Strategy
	for iter.HasNext() {
		response, err := iter.Next()
		if err != nil {
			return nil, err
		}

		var strat Strategy
		err = json.Unmarshal(response.Value, &strat)
		if err != nil {
			return nil, err
		}

		strats = append(strats, &strat)
	}

	return strats, nil
}

// 从公开账本更新策略
func (sl *StrategyList) UpdateStrategy(strat *Strategy) error {
	return sl.stateList.UpdateState(strat)
}

func newStrategyList(ctx TransactionContextInterface) *StrategyList {
	stateList := new(ledgerapi.StateList)
	stateList.Ctx = ctx
	stateList.Name = "org.mynetwork." + constants.STRATEGY + "list"
	stateList.Deserialize = func(bytes []byte, state ledgerapi.StateInterface) error {
		return DeserializeStrategy(bytes, state.(*Strategy))
	}

	list := new(StrategyList)
	list.stateList = stateList

	return list
}

// 从公开账本中删除策略
func (sl *StrategyList) DeleteStrategy(StrategyID string) error {
	return sl.stateList.DelState(GetStrategyKey(StrategyID))
}
