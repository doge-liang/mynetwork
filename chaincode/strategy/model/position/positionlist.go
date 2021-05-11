package position

import (
	"mynetwork/chaincode/strategy/consts"
	ledgerapi "mynetwork/chaincode/strategy/ledger-api"
)

type PositionListInterface interface {
	// AddPositions([]*Position) error
	GetPositionsByStrategyID(string) ([]*Position, error)
	UpdatePositions([]*Position) error
	DeleteAllPositions(string) error
}

type PositionList struct {
	stateList ledgerapi.StateListInterface
}

// func (pl *PositionList) AddPosition(p *Position) error {
// 	return pl.stateList.AddState(p)
// }

// 通过策略ID 获取所有持仓信息
func (pl *PositionList) GetPositionsByStrategyID(strategyID string) ([]*Position, error) {
	ps := *new([]*Position)

	// p := new(Position)
	var states []ledgerapi.StateInterface
	err := pl.stateList.GetStateByPartialCompositeKey([]string{strategyID}, states)

	if err != nil {
		return nil, err
	}

	for _, state := range states {
		ps = append(ps, state.(*Position))
	}

	return ps, nil
}

func (pl *PositionList) DeleteAllPositions(strategyID string) error {
	positionsList, err := pl.GetPositionsByStrategyID(strategyID)

	if err != nil {
		return err
	}

	for _, p := range positionsList {
		err := pl.stateList.DeleteState(GetPositionsKey(strategyID, p.ID))

		if err != nil {
			return err
		}
	}

	return nil
}

// 更新所有策略
func (pl *PositionList) UpdatePositions(ps []*Position) error {
	strategyID := ps[0].StrategyID

	err := pl.DeleteAllPositions(strategyID)

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

func newPositionList(ctx TransactionContextInterface) *PositionList {
	stateList := new(ledgerapi.StateList)
	stateList.Ctx = ctx
	stateList.Name = "org.mynetwork." + consts.POSITIONS + "list"
	stateList.Deserialize = func(bytes []byte, state ledgerapi.StateInterface) error {
		return DeserializePosition(bytes, state.(*Position))
	}

	list := new(PositionList)
	list.stateList = stateList

	return list
}
