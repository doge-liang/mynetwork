package list

import (
	"mynetwork/chaincode/strategy/constants"
	"mynetwork/chaincode/strategy/ledgerapi"
	. "mynetwork/chaincode/strategy/model"
)

// type PositionListInterface interface {
// 	// AddPositions([]*Position) error
// 	GetPositionsByStrategyID(string) ([]*Position, error)
// 	GetPositionsHashByStrategyID(string) ([]*PositionHash, error)
// 	UpdatePositions([]*Position, string) error
// 	// DeleteAllPositions(string) error
// }

type PositionList struct {
	stateList        ledgerapi.StateListInterface
	privateStateList ledgerapi.PrivateStateListInterface
}

// 通过策略ID 获取所有持仓信息
func (pl *PositionList) GetPositionsByStrategyID(strategyID string) ([]*Position, error) {

	iter, err := pl.stateList.GetStateByPartialCompositeKey([]string{strategyID})
	if err != nil {
		return nil, err
	}

	var ps []*Position
	for iter.HasNext() {
		response, err := iter.Next()
		if err != nil {
			return nil, err
		}

		var p Position
		err = DeserializePosition(response.Value, &p)
		if err != nil {
			return nil, err
		}

		ps = append(ps, &p)
	}

	return ps, nil
}

// 通过策略 ID 从私有数据集中获取所有持仓信息
func (pl *PositionList) GetPrivatePositionsByStrategyID(strategyID string) ([]*Position, error) {

	iter, err := pl.privateStateList.GetPrivateDataByPartialCompositeKey([]string{strategyID})
	if err != nil {
		return nil, err
	}

	var ps []*Position
	for iter.HasNext() {
		response, err := iter.Next()
		if err != nil {
			return nil, err
		}

		var p Position
		err = DeserializePosition(response.Value, &p)
		if err != nil {
			return nil, err
		}

		ps = append(ps, &p)
	}

	return ps, nil
}

// 按照策略ID删除仓位信息
func (pl *PositionList) DelPositionsByStrategyID(strategyID string) error {
	positionsList, err := pl.GetPositionsByStrategyID(strategyID)

	if err != nil {
		return err
	}

	for _, p := range positionsList {
		err := pl.stateList.DelState(GetPositionsKey(strategyID, p.ID))

		if err != nil {
			return err
		}
	}

	return nil
}

// 按照策略ID删除信号信息
func (pl *PositionList) DelPrivatePositions(ps []*Position) error {

	for _, p := range ps {
		err := pl.privateStateList.DelStateIn(GetPositionsKey(GetStrategyKey(p.StrategyID), p.ID))

		if err != nil {
			return err
		}
	}

	return nil
}

// 更新策略的仓位信息
func (pl *PositionList) UpdatedPositions(ps []*Position, strategyID string) error {
	err := pl.DelPositionsByStrategyID(strategyID)
	if err != nil {
		return err
	}

	for _, p := range ps {
		err := pl.stateList.UpdateState(p)

		if err != nil {
			return err
		}
	}

	return nil
}

// 更新私有策略的仓位信息
func (pl *PositionList) AddPrivatePositions(ps []*Position) error {
	for _, p := range ps {
		err := pl.privateStateList.UpdateStateIn(p)

		if err != nil {
			return err
		}
	}

	return nil
}

func (pl *PositionList) GetPositionsHashByStrategyID(strategyID string) ([]*PositionHash, error) {
	// 从公共账本查询 ID 列表
	iter, err := pl.privateStateList.GetPrivateDataByPartialCompositeKey([]string{strategyID})
	if err != nil {
		return nil, err
	}

	var pps []*PositionPublic
	for iter.HasNext() {
		response, err := iter.Next()
		if err != nil {
			return nil, err
		}

		var pp PositionPublic
		err = DeserializePositionPublic(response.Value, &pp)
		if err != nil {
			return nil, err
		}

		pps = append(pps, &pp)
	}

	var phs []*PositionHash
	for _, pp := range pps {
		var ph PositionHash
		ph.ID = pp.ID
		err := pl.privateStateList.GetStateHash(GetPositionsKey(pp.StrategyID, pp.ID), &ph.Hashcode)
		if err != nil {
			return nil, err
		}

		phs = append(phs, &ph)
	}

	return phs, nil
}

func newPositionList(ctx TransactionContextInterface) *PositionList {
	stateList := new(ledgerapi.StateList)
	stateList.Ctx = ctx
	stateList.Name = "org.mynetwork." + constants.POSITIONS + "list"
	stateList.Deserialize = func(bytes []byte, state ledgerapi.StateInterface) error {
		return DeserializePosition(bytes, state.(*Position))
	}

	list := new(PositionList)
	list.stateList = stateList

	return list
}

func newPrivatePositionList(ctx TransactionContextInterface, collection string) *PositionList {

	privateStateList := new(ledgerapi.PrivateStateList)
	privateStateList.Ctx = ctx
	privateStateList.Name = "org.mynetwork." + constants.POSITIONS + "list"
	privateStateList.Collection = collection
	privateStateList.Deserialize = func(bytes []byte, state ledgerapi.StateInterface) error {
		return DeserializePosition(bytes, state.(*Position))
	}

	list := newPositionList(ctx)
	list.privateStateList = privateStateList

	return list
}
