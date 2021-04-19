package chaincode

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract definition
type SmartContract struct {
	contractapi.Contract
}

type Trade struct {
	// ID         string    `json:"ID"`         // 交易 ID
	StockID    string    `json:"stockID"`    // 交易股票
	Amount     float64   `json:"amount"`     // 交易份额（买卖用正负来表示）
	Commission float64   `json:"commission"` // 交易佣金
	DateTime   time.Time `json:"dateTime"`   // 交易时间
	Price      float64   `json:"price"`      // 成交价
	// StrategyID string    `json:"strategyID"` // 关联策略 ID
}

type Position struct {
	// ID      string  `json:"ID"`      // 仓位 ID
	StockID string  `json:"stockID"` // 股票代码
	Price   float64 `json:"Price"`   // 现有股价
	Amount  float64 `json:"amount"`  // 仓位
	// StrategyID string  `json:"strategyID"` // 关联策略 ID
}

// 策略公开的数据
type Strategy struct {
	ID           string     `json:"ID"`           // 策略 ID
	Name         string     `json:"name"`         // 策略名
	Provider     string     `json:"provider"`     // 发布者
	MaxDrawdown  float64    `json:"maxDrawdown"`  // 最大回撤
	AnnualReturn float64    `json:"annualReturn"` // 年化收益率
	Subscribers  []string   `json:"subscribers"`  // 订阅者证书列表
	State        string     `json:"state"`        // 是否公开
	Trades       []Trade    `json:"trades"`       // 交易记录
	Positions    []Position `json:"positions"`    // 持仓记录
}

// 策略的交易记录保持 Provider 私有，用户发起 subscribe 后，再由 管理员操作资产转移
type Trades struct {
	StrategyID string  `json:"strategyID"`
	Trades     []Trade `json:"trades"` // 交易记录
}

// 策略的现有仓位保持私有
type Positions struct {
	StrategyID string     `json:"strategyID"`
	Positions  []Position `json:"positions"` // 持仓记录
}

type Subscription struct {
	StrategyID  string   `json:"strategyID"`  // 策略 ID
	Subscribers []string `json:"subscribers"` // 订阅者列表
}

const TRADES_SUFFIX = "_trades"
const POSITIONS_SUFFIX = "_positions"
const PRIVATE_COLLECTION = "ProviderMSPPrivateCollection"
const STRATEGY_COUNT = "STRATEG_COUNT"
const STRATEGY_PERFIX = "startegy"

// InitLedger adds a base set of strategies to the ledger
func (s *SmartContract) Init(ctx contractapi.TransactionContextInterface) error {
	ctx.GetStub().PutState(STRATEGY_COUNT, []byte(strconv.Itoa(0)))
	clientID, err := ctx.GetClientIdentity().GetID()
	if err != nil {
		return err
	}
	strategies := []Strategy{
		{
			// ID:           "strategy1",
			Name:         "测试策略名",
			Provider:     clientID,
			MaxDrawdown:  0.1,
			AnnualReturn: 0.2,
			State:        "public",
			Subscribers:  []string{},
			Trades: []Trade{
				{
					// ID:      "position1",
					StockID: "stock1",
					Price:   100.0,
					Amount:  10.0,
				},
			},
			Positions: []Position{
				{
					// ID:      "position1",
					StockID: "stock1",
					Price:   100.0,
					Amount:  10.0,
				},
			},
		},
		{
			// ID:           "strategy2",
			Name:         "测试策略名",
			Provider:     clientID,
			MaxDrawdown:  0.1,
			AnnualReturn: 0.2,
			State:        "public",
			Subscribers:  []string{},
			Trades: []Trade{
				{
					// ID:         "trade2",
					StockID:    "stock1",
					Amount:     10.0,
					Commission: 10.0,
					DateTime:   time.Now(),
					Price:      100.0,
				},
			},
			Positions: []Position{
				{
					// ID:      "position2",
					StockID: "stock1",
					Price:   100.0,
					Amount:  10.0,
				},
			},
		},
		{
			// ID:           "strategy3",
			Name:         "测试策略名",
			Provider:     clientID,
			MaxDrawdown:  0.1,
			AnnualReturn: 0.2,
			State:        "private",
			Subscribers:  []string{},
			Trades: []Trade{
				{
					// ID:         "trade3",
					StockID:    "stock1",
					Amount:     10.0,
					Commission: 10.0,
					DateTime:   time.Now(),
					Price:      100.0,
				},
			},
			Positions: []Position{
				{
					// ID:      "position3",
					StockID: "stock1",
					Price:   100.0,
					Amount:  10.0,
				},
			},
		},
		{
			// ID:           "strategy4",
			Name:         "测试策略名",
			Provider:     clientID,
			MaxDrawdown:  0.1,
			AnnualReturn: 0.2,
			State:        "public",
			Subscribers:  []string{},
			Trades:       []Trade{},
			Positions:    []Position{},
		},
	}

	for _, strategy := range strategies {

		if strategy.State == "private" {
			err = s.SaveStrategyPrivate(ctx, &strategy)
			if err != nil {
				return err
			}
		}
		err = s.SaveStrategy(ctx, &strategy)

		if err != nil {
			return err
		}
	}

	return nil
}

// 保存公有策略
func (s *SmartContract) SaveStrategy(ctx contractapi.TransactionContextInterface, strategy *Strategy) error {
	strategyCountBytes, _ := ctx.GetStub().GetState(STRATEGY_COUNT)
	strategyCount, _ := strconv.Atoi(string(strategyCountBytes))
	strategy.ID = STRATEGY_PERFIX + string(strategyCountBytes)
	strategyJSON, err := json.Marshal(strategy)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState(strategy.ID, strategyJSON)
	err = ctx.GetStub().PutState(STRATEGY_COUNT, []byte(strconv.Itoa(strategyCount+1)))

	if err != nil {
		return fmt.Errorf("failed to put to world state. %v", err)
	}
	return nil
}

// 保存私有策略
func (s *SmartContract) SaveStrategyPrivate(ctx contractapi.TransactionContextInterface, strategy *Strategy) error {
	strategyCountBytes, _ := ctx.GetStub().GetState(STRATEGY_COUNT)
	strategyCount, _ := strconv.Atoi(string(strategyCountBytes))
	strategy.ID = STRATEGY_PERFIX + string(strategyCountBytes)
	positions := Positions{
		StrategyID: strategy.ID,
		Positions:  strategy.Positions,
	}
	trades := Trades{
		StrategyID: strategy.ID,
		Trades:     strategy.Trades,
	}

	positionsJSON, err := json.Marshal(positions)
	if err != nil {
		return err
	}
	tradesJSON, err := json.Marshal(trades)
	if err != nil {
		return err
	}

	err = ctx.GetStub().PutPrivateData(PRIVATE_COLLECTION, strategy.ID+POSITIONS_SUFFIX, positionsJSON)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutPrivateData(PRIVATE_COLLECTION, strategy.ID+TRADES_SUFFIX, tradesJSON)
	if err != nil {
		return err
	}

	// 清空公共部分
	strategy.Trades = []Trade{}
	strategy.Positions = []Position{}

	err = s.SaveStrategy(ctx, strategy)
	err = ctx.GetStub().PutState(STRATEGY_COUNT, []byte(strconv.Itoa(strategyCount+1)))

	if err != nil {
		return fmt.Errorf("failed to put to world state. %v", err)
	}

	return nil
}

// 更新策略
func (s *SmartContract) UpdateStrategy(ctx contractapi.TransactionContextInterface, strategy *Strategy) error {
	exist, err := s.StrategyExists(ctx, strategy.ID)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("the strategy %s does not exist", strategy.ID)
	}

	strategyJSON, err := json.Marshal(strategy)

	return ctx.GetStub().PutState(strategy.ID, strategyJSON)
}

// 将策略状态改为公共
func (s *SmartContract) SetStrategyPublic(ctx contractapi.TransactionContextInterface, id string) error {
	trades, err := s.ReadTrades(ctx, id)
	if err != nil {
		return err
	}
	positions, err := s.ReadPositions(ctx, id)
	if err != nil {
		return err
	}

	// 移除私有数据

	strategy, err := s.ReadStrategy(ctx, id)
	if err != nil {
		return err
	}
	// 添加私有数据并修改状态
	strategy.Trades = trades.Trades
	strategy.Positions = positions.Positions
	strategy.State = "public"
	return s.SaveStrategy(ctx, strategy)
}

// 将策略状态改为私有
func (s *SmartContract) SetStrategyPrivate(ctx contractapi.TransactionContextInterface, id string) error {
	strategy, err := s.ReadStrategy(ctx, id)
	if err != nil {
		return err
	}
	strategy.State = "private"
	return s.SaveStrategyPrivate(ctx, strategy)
}

func (s *SmartContract) DeleteTrades(ctx contractapi.TransactionContextInterface, id string) error {
	return ctx.GetStub().DelPrivateData(PRIVATE_COLLECTION, id)
}
