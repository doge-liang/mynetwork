package chaincode

import (
	"encoding/json"
	"fmt"
	"strings"
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

const TRADES = "trades"
const POSITIONS = "positions"
const STRATEGY = "strategy"
const PRIVATE_COLLECTION = "ProviderMSPPrivateCollection"
const PUBLIC_COLLECTION = "strategyPublicCollection"

func (s *SmartContract) subscribe(ctx contractapi.TransactionContextInterface, id string) error {
	strategy, err := s.ReadStrategy(ctx, id)
	if err != nil {
		return err
	}
	if strategy == nil {
		return fmt.Errorf("Strategy %s does not exist", id)
	}
	clientID, err := ctx.GetClientIdentity().GetID()
	if err != nil {
		return err
	}
	key := GetStrategyKey(id)

	strategy.Subscribers = append(strategy.Subscribers, clientID)
	strategyJSON, err := json.Marshal(strategy)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(key, strategyJSON)
}

func MakeKey(keyParts ...string) string {
	return strings.Join(keyParts, "_")
}

func SplitKey(key string) []string {
	return strings.Split(key, "_")
}

func GetTradesKey(strategyKey string) string {
	keyParts := SplitKey(strategyKey)
	return MakeKey(STRATEGY, TRADES, keyParts[len(keyParts)-1])
}

func GetPositionsKey(strategyKey string) string {
	keyParts := SplitKey(strategyKey)
	return MakeKey(STRATEGY, POSITIONS, keyParts[len(keyParts)-1])
}

func GetStrategyKey(strategyKey string) string {
	keyParts := SplitKey(strategyKey)
	return MakeKey(STRATEGY, keyParts[len(keyParts)-1])
}
