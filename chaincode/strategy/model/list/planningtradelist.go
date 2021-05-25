package list

import (
	"log"
	"mynetwork/chaincode/strategy/constants"
	"mynetwork/chaincode/strategy/ledgerapi"

	. "mynetwork/chaincode/strategy/model"
)

// type PlanningTradeListInterface interface {
// 	// AddPlanningTrade(*PlanningTrade) error
// 	GetPlanningTradesByStrategyID(string) ([]*PlanningTrade, error)
// 	UpdatePlanningTrades([]*PlanningTrade, string) error
// 	GetPlanningTradesHashByStrategyID(string) ([]*PlanningTradeHash, error)
// }

type PlanningTradeList struct {
	stateList        ledgerapi.StateListInterface
	privateStateList ledgerapi.PrivateStateListInterface
}

// func (ptl *PlanningTradeList) AddPlanningTrade(planningTrade *PlanningTrade) error {
// 	return ptl.stateList.AddState(planningTrade)
// }

// 根据策略ID 获取信号
func (ptl *PlanningTradeList) GetPlanningTradesByStrategyID(strategyID string) ([]*PlanningTrade, error) {

	iter, err := ptl.stateList.GetStateByPartialCompositeKey([]string{strategyID})
	if err != nil {
		return nil, err
	}
	var pts []*PlanningTrade
	for iter.HasNext() {
		response, err := iter.Next()
		if err != nil {
			return nil, err
		}

		var pt PlanningTrade
		err = DeserializePlanningTrade(response.Value, &pt)
		if err != nil {
			return nil, err
		}
		pts = append(pts, &pt)
	}

	return pts, err
}

// 根据策略ID 从私有数据集中获取信号
func (ptl *PlanningTradeList) GetPrivatePlanningTradesByStrategyID(strategyID string) ([]*PlanningTrade, error) {
	iter, err := ptl.privateStateList.GetPrivateDataByPartialCompositeKey([]string{strategyID})
	log.Print(strategyID)
	if err != nil {
		return nil, err
	}
	var pts []*PlanningTrade
	for iter.HasNext() {
		response, err := iter.Next()
		if err != nil {
			return nil, err
		}

		var pt PlanningTrade
		err = DeserializePlanningTrade(response.Value, &pt)
		if err != nil {
			return nil, err
		}
		pts = append(pts, &pt)
	}

	return pts, err
}

// 按照策略ID删除信号信息
func (ptl *PlanningTradeList) DelPlanningTradesByStrategyID(strategyID string) error {
	pts, err := ptl.GetPlanningTradesByStrategyID(strategyID)

	if err != nil {
		return err
	}

	for _, p := range pts {
		err := ptl.stateList.DelState(GetPlanningTradesKey(GetStrategyKey(strategyID), p.ID))
		if err != nil {
			return err
		}

	}

	return nil
}

// 按照策略ID删除信号信息
func (ptl *PlanningTradeList) DelPrivatePlanningTrades(pts []*PlanningTrade) error {

	for _, p := range pts {
		err := ptl.privateStateList.DelStateIn(GetPlanningTradesKey(GetStrategyKey(p.StrategyID), p.ID))

		if err != nil {
			return err
		}
	}

	return nil
}

// 更新公开的信号
func (ptl *PlanningTradeList) UpdatePlanningTrades(pts []*PlanningTrade, strategyID string) error {
	// if len(pts) == 0 {
	// 	return fmt.Errorf("Receive an empty planntingtrades slice")
	// }
	err := ptl.DelPlanningTradesByStrategyID(strategyID)
	if err != nil {
		return err
	}
	for _, pt := range pts {
		err := ptl.stateList.UpdateState(pt)
		if err != nil {
			return err
		}
	}
	return nil
}

// 更新私有数据集中的信号
func (ptl *PlanningTradeList) AddPrivatePlanningTrades(pts []*PlanningTrade) error {
	for _, pt := range pts {
		err := ptl.privateStateList.UpdateStateIn(pt)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ptl *PlanningTradeList) GetPlanningTradesHashByStrategyID(strategyID string) ([]*PlanningTradeHash, error) {
	// 从公共账本查询 ID 列表
	iter, err := ptl.privateStateList.GetPrivateDataByPartialCompositeKey([]string{strategyID})
	if err != nil {
		return nil, err
	}

	var ptps []*PlanningTradePublic
	for iter.HasNext() {
		response, err := iter.Next()
		if err != nil {
			return nil, err
		}

		var ptp PlanningTradePublic
		err = DeserializePlanningTradePublic(response.Value, &ptp)
		if err != nil {
			return nil, err
		}
		// log.Print(ptp.ID + ":" + ptp.StrategyID)
		ptps = append(ptps, &ptp)
	}

	var pths []*PlanningTradeHash
	for _, ptp := range ptps {
		var pth PlanningTradeHash
		pth.ID = ptp.ID
		err := ptl.privateStateList.GetStateHash(GetPlanningTradesKey(ptp.StrategyID, ptp.ID), &pth.Hashcode)
		if err != nil {
			return nil, err
		}

		pths = append(pths, &pth)
	}

	return pths, nil
}

func newPlanningTradeList(ctx TransactionContextInterface) *PlanningTradeList {
	stateList := new(ledgerapi.StateList)
	stateList.Ctx = ctx

	stateList.Name = "org.mynetwork." + constants.PLANNINGTRADES + "list"
	stateList.Deserialize = func(bytes []byte, state ledgerapi.StateInterface) error {
		return DeserializePlanningTrade(bytes, state.(*PlanningTrade))
	}

	list := new(PlanningTradeList)
	list.stateList = stateList

	return list
}

func newPrivatePlanningTradeList(ctx TransactionContextInterface, collection string) *PlanningTradeList {
	privateStateList := new(ledgerapi.PrivateStateList)
	privateStateList.Ctx = ctx

	privateStateList.Name = "org.mynetwork." + constants.PLANNINGTRADES + "list"
	privateStateList.Collection = collection
	privateStateList.Deserialize = func(bytes []byte, state ledgerapi.StateInterface) error {
		return DeserializePlanningTrade(bytes, state.(*PlanningTrade))
	}

	list := newPlanningTradeList(ctx)
	list.privateStateList = privateStateList

	return list
}
