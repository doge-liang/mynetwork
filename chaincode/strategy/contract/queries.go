package contract

import (
	"log"
	"mynetwork/chaincode/strategy/constants"
	. "mynetwork/chaincode/strategy/model"
	. "mynetwork/chaincode/strategy/model/list"
	. "mynetwork/chaincode/strategy/utils"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract definition
type SmartContract struct {
	contractapi.Contract
}

type Subscription struct {
	StrategyID  string   `json:"strategyID"`  // 策略 ID
	Subscribers []string `json:"subscribers"` // 订阅者列表
}

// type PublicStrategy struct {
// 	ID   string `json:"id"`   // 策略 ID
// 	Name string `json:"name"` // 策略名
// 	// Provider     string     `json:"provider"`     // 发布者
// 	MaxDrawdown  float64    `json:"maxDrawdown"`  // 最大回撤
// 	AnnualReturn float64    `json:"annualReturn"` // 年化收益率
// 	State        string     `json:"state"`        // 是否公开
// 	Trades       []Trade    `json:"trades"`       // 交易记录
// 	Positions    []Position `json:"positions"`    // 持仓记录
// }

// 获取所有策略
func (s *SmartContract) GetAllStrategies(ctx TransactionContextInterface) ([]*StrategyOutput, error) {

	strats, err := ctx.GetStrategyList().GetAllStrategies()
	if err != nil {
		return nil, err
	}

	var output []*StrategyOutput
	for _, strat := range strats {
		log.Print(strat)
		stratOutput := new(StrategyOutput)
		isSub, err := s.IsSubscirbed(ctx, strat.ID)
		if err != nil {
			return nil, err
		}
		log.Print(strat.ID)
		stratOutput.ID = strat.ID
		stratOutput.Name = strat.Name
		stratOutput.State = strat.State
		stratOutput.SharpeRatio = strat.SharpeRatio
		stratOutput.MaxDrawdown = strat.MaxDrawdown
		stratOutput.AnnualReturn = strat.AnnualReturn
		stratOutput.Subscribers = []string{}
		stratOutput.IsSub = isSub
		output = append(output, stratOutput)
	}

	return output, nil
}

// func (s *SmartContract) GetAllTradesByStrategyID(ctx TransactionContextInterface, strategyID string) error {
// 	_, err := ctx.GetTradeList().GetTradesByStrategyID(strategyID)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// 通过策略 ID 读取策略交易记录页面
func (s *SmartContract) GetTradesPageByStrategyID(ctx TransactionContextInterface, strategyID string, bookmark string, pageSize int32) (*TradesOutput, error) {
	output := TradesOutput{
		Trades:   []*Trade{},
		Bookmark: bookmark,
	}

	trades, bookmark, err := ctx.GetTradeList().GetTradesByStrategyIDPage(strategyID, bookmark, pageSize)
	if err != nil {
		return &output, err
	} else if trades == nil {
		return &output, nil
	}

	output.Trades = trades
	output.Bookmark = bookmark

	return &output, nil
	// return nil, fmt.Errorf("unknown MSPID: %s", orgID)
}

// 根据策略 ID 获取持仓信息
func (s *SmartContract) GetPositionsByStrategyID(ctx TransactionContextInterface, strategyID string) (*PositionsOutput, error) {
	output := PositionsOutput{
		Positions:     []*Position{},
		PositionsHash: []*PositionHash{},
	}

	// return &output, nil

	strat, err := ctx.GetStrategyList().GetStrategy(strategyID)
	if err != nil {
		return &output, err
	}

	// 私有策略
	if strat.IsPrivate() {
		isSubed, err := s.IsSubscirbed(ctx, strategyID)
		if err != nil {
			return &output, err
		}
		isProvided, err := s.IsProvided(ctx, strategyID)
		if err != nil {
			return nil, err
		}

		if isProvided {
			log.Print("是发布者")
			plist, err := ctx.GetPrivatePositionList(constants.PRIVATE_COLLECTION).GetPrivatePositionsByStrategyID(strategyID)
			if err != nil {
				return &output, err
			} else if plist == nil {
				return &output, nil
			}

			output = PositionsOutput{
				Positions:     plist,
				PositionsHash: []*PositionHash{},
			}

			return &output, nil
		}
		// 未订阅且不是发布者
		if !isSubed {
			log.Print("未订阅")
			phlist, err := ctx.GetPrivatePositionList(constants.PUBLIC_COLLECTION).GetPositionsHashByStrategyID(strategyID)
			if err != nil {
				return &output, err
			}

			output = PositionsOutput{
				Positions:     []*Position{},
				PositionsHash: phlist,
			}

			return &output, nil
		}

		// 已订阅
		log.Print("已订阅")
		ptlist, err := ctx.GetPrivatePositionList(constants.PUBLIC_COLLECTION).GetPrivatePositionsByStrategyID(strategyID)
		if err != nil {
			return &output, err
		}

		output = PositionsOutput{
			Positions:     ptlist,
			PositionsHash: []*PositionHash{},
		}

		return &output, nil
	}
	// 公开策略
	plist, err := ctx.GetPositionList().GetPositionsByStrategyID(strategyID)
	if err != nil {
		return &output, err
	}

	output = PositionsOutput{
		Positions:     plist,
		PositionsHash: []*PositionHash{},
	}

	return &output, nil
}

// 根据策略 ID 获取交易信号
func (s *SmartContract) GetPlanningTradesByStrategyID(ctx TransactionContextInterface, strategyID string) (*PlanningTradesOutput, error) {
	output := PlanningTradesOutput{
		PlanningTrades:     []*PlanningTrade{},
		PlanningTradesHash: []*PlanningTradeHash{},
	}

	strat, err := ctx.GetStrategyList().GetStrategy(strategyID)
	if err != nil {
		return &output, err
	}

	// 私有策略
	if strat.IsPrivate() {
		isSubed, err := s.IsSubscirbed(ctx, strategyID)
		if err != nil {
			return &output, err
		}

		isProvided, err := s.IsProvided(ctx, strategyID)
		if err != nil {
			return nil, err
		}

		if isProvided {
			log.Print("是发布者")
			ptlist, err := ctx.GetPrivatePlanningTradeList(constants.PRIVATE_COLLECTION).GetPrivatePlanningTradesByStrategyID(strategyID)
			if err != nil {
				return &output, err
			} else if ptlist == nil {
				return &output, nil
			}
			log.Print(ptlist)

			output = PlanningTradesOutput{
				PlanningTrades:     ptlist,
				PlanningTradesHash: []*PlanningTradeHash{},
			}

			return &output, nil
		}

		// 未订阅
		if !isSubed {
			log.Print("未订阅")
			pthlist, err := ctx.GetPrivatePlanningTradeList(constants.PUBLIC_COLLECTION).GetPlanningTradesHashByStrategyID(strategyID)
			if err != nil {
				return &output, err
			} else if pthlist == nil {
				return &output, nil
			}

			output = PlanningTradesOutput{
				PlanningTrades:     []*PlanningTrade{},
				PlanningTradesHash: pthlist,
			}

			return &output, nil
		}
		// 已订阅
		log.Print("已订阅")
		ptlist, err := ctx.GetPrivatePlanningTradeList(constants.PUBLIC_COLLECTION).GetPrivatePlanningTradesByStrategyID(strategyID)
		if err != nil {
			return &output, err
		} else if ptlist == nil {
			return &output, nil
		}

		output = PlanningTradesOutput{
			PlanningTrades:     ptlist,
			PlanningTradesHash: []*PlanningTradeHash{},
		}

		return &output, nil

	}
	// 公开策略
	ptlist, err := ctx.GetPlanningTradeList().GetPlanningTradesByStrategyID(strategyID)
	if err != nil {
		return &output, err
	} else if ptlist == nil {
		return &output, nil
	}

	output = PlanningTradesOutput{
		PlanningTrades:     ptlist,
		PlanningTradesHash: []*PlanningTradeHash{},
	}

	return &output, nil
}

func (s *SmartContract) IsSubscirbed(ctx TransactionContextInterface, strategyID string) (bool, error) {
	strat, err := ctx.GetStrategyList().GetStrategy(strategyID)
	if err != nil {
		return false, err
	}

	clientID, err := ctx.GetClientIdentity().GetID()
	if err != nil {
		return false, err
	}

	return In(clientID, strat.Subscribers), nil
}

func (s *SmartContract) IsProvided(ctx TransactionContextInterface, strategyID string) (bool, error) {
	strat, err := ctx.GetStrategyList().GetStrategy(strategyID)
	if err != nil {
		return false, err
	}

	clientID, err := ctx.GetClientIdentity().GetID()
	if err != nil {
		return false, err
	}

	log.Print(clientID)

	return clientID == strat.Provider, nil
}
