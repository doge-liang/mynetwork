package planningTrade

import (
	"mynetwork/chaincode/strategy/consts"
	ledgerapi "mynetwork/chaincode/strategy/ledger-api"
)

type PlanningTradeListInterface interface {
	AddPlanningTrade(*PlanningTrade) error
	GetPlanningTrade(string) (*PlanningTrade, error)
	UpdatePlanningTrade(*PlanningTrade) error
}

type PlanningTradeList struct {
	stateList ledgerapi.StateListInterface
}

func (ptl *PlanningTradeList) AddPlanningTrade(planningTrade *PlanningTrade) error {
	return ptl.stateList.AddState(planningTrade)
}

func (ptl *PlanningTradeList) GetPlanningTrade(id string) (*PlanningTrade, error) {
	pt := new(PlanningTrade)

	err := ptl.stateList.GetState(id, pt)

	if err != nil {
		return nil, err
	}

	return pt, err
}

func (ptl *PlanningTradeList) UpdatePlanningTrade(planningTrade *PlanningTrade) error {
	return ptl.stateList.UpdateState(planningTrade)
}

func newPlanningTradeList(ctx TransactionContextInterface) *PlanningTradeList {
	stateList := new(ledgerapi.StateList)
	stateList.Ctx = ctx
	
	stateList.Name = "org.mynetwork." + consts.PLANNINGTRADES + "list"
	stateList.Deserialize = func(bytes []byte, state ledgerapi.StateInterface) error {
		return DeserializePlanningTrade(bytes, state.(*PlanningTrade))
	}

	list := new(PlanningTradeList)
	list.stateList = stateList

	return list
}
