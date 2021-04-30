package chaincode

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract definition
type SmartContract struct {
	contractapi.Contract
}

type Trade struct {
	ID         string    `json:"ID"`         // 交易 ID
	StockID    string    `json:"stockID"`    // 交易股票
	Amount     float64   `json:"amount"`     // 交易份额（买卖用正负来表示）
	Commission float64   `json:"commission"` // 交易佣金
	DateTime   time.Time `json:"dateTime"`   // 交易时间戳
	Price      float64   `json:"price"`      // 成交价
	// StrategyID string    `json:"strategyID"` // 关联策略 ID
}

type PlanningTrade struct {
	StockID string  `json:"stockID"` // 交易股票
	Amount  float64 `json:"amount"`  // 交易份额（买卖用正负来表示）
}

type Position struct {
	// ID      string  `json:"ID"`      // 仓位 ID
	StockID string  `json:"stockID"` // 股票代码
	Value   float64 `json:"value"`   // 仓位
	// Price   float64 `json:"Price"`   // 现有股价
	// StrategyID string  `json:"strategyID"` // 关联策略 ID
}

// 策略公开的数据
type Strategy struct {
	ID             string          `json:"ID"`             // 策略 ID
	Name           string          `json:"name"`           // 策略名
	Provider       string          `json:"provider"`       // 发布者
	MaxDrawdown    float64         `json:"maxDrawdown"`    // 最大回撤
	AnnualReturn   float64         `json:"annualReturn"`   // 年化收益率
	SharpeRatio    float64         `json:"sharpeRatio"`    // 夏普率
	Subscribers    []string        `json:"subscribers"`    // 订阅者证书列表
	State          string          `json:"state"`          // 是否公开
	Trades         []Trade         `json:"trades"`         // 交易记录
	PlanningTrades []PlanningTrade `json:"planningTrades"` // 计划交易
	Positions      []Position      `json:"positions"`      // 持仓记录
}

// 策略的交易记录保持 Provider 私有，用户发起 subscribe 后，再由 管理员操作资产转移
type Trades struct {
	StrategyID     string          `json:"strategyID"`
	Trades         []Trade         `json:"trades"`         // 交易记录
	PlanningTrades []PlanningTrade `json:"planningTrades"` // 计划交易
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

type PrivateStrategy struct {
	ID   string `json:"ID"`   // 策略 ID
	Name string `json:"name"` // 策略名
	// Provider     string  `json:"provider"`     // 发布者
	MaxDrawdown  float64 `json:"maxDrawdown"`  // 最大回撤
	AnnualReturn float64 `json:"annualReturn"` // 年化收益率
	State        string  `json:"state"`        // 是否公开
	Trades       string  `json:"trades"`       // 交易记录Hash
	Positions    string  `json:"positions"`    // 持仓记录Hash
}

type PublicStrategy struct {
	ID   string `json:"ID"`   // 策略 ID
	Name string `json:"name"` // 策略名
	// Provider     string     `json:"provider"`     // 发布者
	MaxDrawdown  float64    `json:"maxDrawdown"`  // 最大回撤
	AnnualReturn float64    `json:"annualReturn"` // 年化收益率
	State        string     `json:"state"`        // 是否公开
	Trades       []Trade    `json:"trades"`       // 交易记录
	Positions    []Position `json:"positions"`    // 持仓记录
}

const TRADES = "trades"
const POSITIONS = "positions"
const STRATEGY = "strategy"
const PRIVATE_COLLECTION = "ProviderMSPPrivateCollection"
const PUBLIC_COLLECTION = "strategyPublicCollection"

// 公共合约

// 获取所有策略
func (s *SmartContract) GetAllStrategies(ctx contractapi.TransactionContextInterface) ([]*Strategy, error) {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all strategies in the chaincode namespace.
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var strategies []*Strategy
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var strategy Strategy
		err = json.Unmarshal(queryResponse.Value, &strategy)
		if err != nil {
			return nil, err
		}
		strategies = append(strategies, &strategy)
	}

	return strategies, nil
}

// 读取策略
func (s *SmartContract) ReadStrategy(ctx contractapi.TransactionContextInterface, id string) (*Strategy, error) {
	orgID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return nil, err
	}
	if orgID == "SubscriberMSP" {
		return s.sub_ReadStrategy(ctx, id)
	}
	if orgID == "ProviderMSP" {
		return s.pro_ReadStrategy(ctx, id)
	}

	return nil, fmt.Errorf("unknown MSPID: %s", orgID)
}

// 用 ID 检查策略是否存在
func (s *SmartContract) StrategyExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	strategyJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return strategyJSON != nil, nil
}

func (s *SmartContract) ReadPrivateStrategy(ctx contractapi.TransactionContextInterface, id string) (*PrivateStrategy, error) {
	key := GetStrategyKey(id)
	strategyJSON, err := ctx.GetStub().GetState(key)
	if err != nil {
		return nil, err
	}
	var strategy Strategy
	err = json.Unmarshal(strategyJSON, &strategy)
	if err != nil {
		return nil, err
	}
	TradesHash, err := s.GetTradesHash(ctx, PRIVATE_COLLECTION, id)
	if err != nil {
		return nil, err
	}
	PositionsHash, err := s.GetPositionsHash(ctx, PRIVATE_COLLECTION, id)
	if err != nil {
		return nil, err
	}
	privateStrategy := PrivateStrategy{
		ID:           strategy.ID,
		Name:         strategy.Name,
		MaxDrawdown:  strategy.MaxDrawdown,
		AnnualReturn: strategy.AnnualReturn,
		State:        strategy.State,
		Trades:       TradesHash,
		Positions:    PositionsHash,
	}
	return &privateStrategy, nil
}

// 通过策略的 id 得到交易数据
func (s *SmartContract) ReadTrades(ctx contractapi.TransactionContextInterface, id string) ([]Trade, error) {
	tradesJSON, err := ctx.GetStub().GetPrivateData(PRIVATE_COLLECTION, GetTradesKey(id))
	if err != nil {
		return nil, err
	}
	var trades Trades
	err = json.Unmarshal(tradesJSON, &trades)
	if err != nil {
		return nil, err
	}
	return trades.Trades, nil
}

// 通过策略的 id 获得仓位数据
func (s *SmartContract) ReadPositions(ctx contractapi.TransactionContextInterface, id string) ([]Position, error) {
	positionsJSON, err := ctx.GetStub().GetPrivateData(PRIVATE_COLLECTION, GetPositionsKey(id))
	if err != nil {
		return nil, err
	}
	var positions Positions
	err = json.Unmarshal(positionsJSON, &positions)
	if err != nil {
		return nil, err
	}

	return positions.Positions, nil
}

// 订阅者的合约

// 根据 Strategy 的 ID 读取某条策略
// 如果是公有的直接 GetState(key) 即可
// 如果是私有的则需要使用 PrivateStrategy 结构来返回数据，把哈希值填私有的数据部分里面
func (s *SmartContract) sub_ReadStrategy(ctx contractapi.TransactionContextInterface, id string) (*Strategy, error) {
	key := GetStrategyKey(id)
	strategyJSON, err := ctx.GetStub().GetState(key)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if strategyJSON == nil {
		return nil, fmt.Errorf("the strategy %s does not exist", id)
	}

	var strategy Strategy
	err = json.Unmarshal(strategyJSON, &strategy)
	if err != nil {
		return nil, err
	}

	clientID, err := ctx.GetClientIdentity().GetID()
	if err != nil {
		return nil, err
	}

	// 私有策略
	if strategy.State == "private" {
		// 未订阅，无法访问该链码，应通过 ReadPrivateStrategy() 查询
		if !in(clientID, strategy.Subscribers) {
			return nil, fmt.Errorf("Your have not subscribe this strategy yet.")
		}
		// TODO 已订阅，检查发布者公布的数据和私有数据的 hash 是否一致，一致的才公布
		pubilcTradesHash, _ := s.GetTradesHash(ctx, PUBLIC_COLLECTION, id)
		publicPositionsHash, _ := s.GetTradesHash(ctx, PUBLIC_COLLECTION, id)
		privateTradesHash, _ := s.GetPositionsHash(ctx, PRIVATE_COLLECTION, id)
		privatePositionsHash, _ := s.GetPositionsHash(ctx, PUBLIC_COLLECTION, id)
		if pubilcTradesHash == privateTradesHash && publicPositionsHash == privatePositionsHash {
			tradesJSON, err := ctx.GetStub().GetPrivateData(PUBLIC_COLLECTION, GetTradesKey(id))
			if err != nil {
				return nil, err
			}
			positionsJSON, err := ctx.GetStub().GetPrivateData(PUBLIC_COLLECTION, GetPositionsKey(id))
			if err != nil {
				return nil, err
			}
			if tradesJSON != nil {
				var trades Trades
				err = json.Unmarshal(tradesJSON, &trades)
				strategy.Trades = trades.Trades
			}
			if positionsJSON != nil {
				var positions Positions
				err = json.Unmarshal(positionsJSON, &positions)
				strategy.Positions = positions.Positions

			}

			return &strategy, err
		}
	}

	return &strategy, nil
}

func (s *SmartContract) pro_ReadStrategy(ctx contractapi.TransactionContextInterface, id string) (*Strategy, error) {
	strategyJSON, err := ctx.GetStub().GetState(GetStrategyKey(id))
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if strategyJSON == nil {
		return nil, fmt.Errorf("the strategy %s does not exist", id)
	}

	var strategy Strategy
	err = json.Unmarshal(strategyJSON, &strategy)
	if err != nil {
		return nil, err
	}
	// 对私有策略的特殊处理
	if strategy.State == "private" {
		key := MakeKey(STRATEGY, TRADES, id)
		trades, _ := s.ReadTrades(ctx, key)
		strategy.Trades = trades
		key = MakeKey(STRATEGY, POSITIONS, id)
		positions, _ := s.ReadPositions(ctx, key)
		strategy.Positions = positions
		return &strategy, nil
	}

	return &strategy, nil
}

func (s *SmartContract) GetTradesHash(ctx contractapi.TransactionContextInterface, collection string, id string) (string, error) {
	tradesKey := GetTradesKey(id)
	tradesHash, err := ctx.GetStub().GetPrivateDataHash(collection, tradesKey)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", tradesHash), nil
}

func (s *SmartContract) GetPositionsHash(ctx contractapi.TransactionContextInterface, collection string, id string) (string, error) {
	positionsKey := GetTradesKey(id)
	positionsHash, err := ctx.GetStub().GetPrivateDataHash(collection, positionsKey)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", positionsHash), nil
}

func (s *SmartContract) IsSubscirbed(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	strategy, err := s.ReadStrategy(ctx, id)
	if err != nil {
		return false, err
	}

	clientID, err := ctx.GetClientIdentity().GetID()
	if err != nil {
		return false, err
	}

	return in(clientID, strategy.Subscribers), nil
}
