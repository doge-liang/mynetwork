package strategy

import (
	"mynetwork/chaincode/strategy/consts"
	ledgerapi "mynetwork/chaincode/strategy/ledger-api"
)

type StrategyListInterface interface {
	AddStrategy(*Strategy) error
	GetStrategy(string, string) (*Strategy, error)
	UpdateStrategy(*Strategy) error
}

type StrategyList struct {
	stateList ledgerapi.StateListInterface
}

// 添加策略到公开账本
func (stratlist *StrategyList) AddStrategy(strategy *Strategy) error {
	return stratlist.stateList.AddState(strategy)
}

// 添加策略到私有数据集
func (stratList *StrategyList) AddStrategyIn(collection string, strategy *Strategy) error {
	return stratList.AddStrategyIn(collection, strategy)
}

// 从公开账本查询策略
func (stratlist *StrategyList) GetStrategy(provider string, id string) (*Strategy, error) {
	strat := new(Strategy)

	err := stratlist.stateList.GetState(GetStrategyKey(provider, id), strat)

	return strat, err
}

// 从隐私数据集查询策略
func (stratList *StrategyList) GetStrategyIn(provider string, id string, collection string) (*Strategy, error) {
	strat := new(Strategy)

	err := stratList.stateList.GetStateIn(GetStrategyKey(provider, id), collection, strat)

	return strat, err
}

// 从公开账本更新策略
func (stratlist *StrategyList) UpdateStrategy(strategy *Strategy) error {
	return stratlist.stateList.UpdateState(strategy)
}

// 从隐私数据集更新策略
func (stratlist *StrategyList) UpdateStrategyIn(collection string, strategy *Strategy) error {
	return stratlist.stateList.UpdateStateIn(collection, strategy)
}

func newStrategyList(ctx TransactionContextInterface) *StrategyList {
	stateList := new(ledgerapi.StateList)
	stateList.Ctx = ctx
	stateList.Name = "org.mynetwork." + consts.STRATEGY + "list"
	stateList.Deserialize = func(bytes []byte, state ledgerapi.StateInterface) error {
		return DeserializeStrategy(bytes, state.(*Strategy))
	}

	list := new(StrategyList)
	list.stateList = stateList

	return list
}
