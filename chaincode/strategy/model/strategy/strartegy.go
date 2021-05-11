package strategy

import (
	"encoding/json"
	"fmt"
	ledgerapi "mynetwork/chaincode/strategy/ledger-api"
	"mynetwork/chaincode/strategy/model/trade"
	"time"
)

type State uint

const (
	PUBLIC State = iota
	PRIVATE
)

func (state State) String() string {
	names := []string{"PUBLIC", "PRIVATE"}

	if state != PUBLIC && state != PRIVATE {
		return "UNKNOWN"
	}

	return names[state]
}

func GetStrategyKey(provider string, strategyID string) string {
	// keyParts := ledgerapi.SplitKey(strategyKey)
	return ledgerapi.MakeKey(provider, strategyID)
}

// 策略数据实体
type Strategy struct {
	ID           string    `json:"ID"`           // 策略 ID
	Name         string    `json:"name"`         // 策略名
	Provider     string    `json:"provider"`     // 发布者
	MaxDrawdown  float64   `json:"maxDrawdown"`  // 最大回撤
	AnnualReturn float64   `json:"annualReturn"` // 年化收益率
	SharpeRatio  float64   `json:"sharpeRatio"`  // 夏普率
	State        State     `json:"state"`        // 是否公开
	Subscribers  []string  `json:"subscribers"`  // 订阅者证书列表
	CreateAt     time.Time `json:"createAt"`     // 创建时间
	// Trades         []trade.Trade                 `json:"trades"`         // 交易记录
	// PlanningTrades []planningTrade.PlanningTrade `json:"planningTrades"` // 计划交易
	// Positions      []position.Position           `json:"positions"`      // 持仓记录
}

// 订阅者查看隐私策略时返回的数据
type PrivateStrategy struct {
	*Strategy
	Trades         []trade.Trade `json:"trades"`         // 交易记录Hash
	Positions      []string      `json:"positions"`      // 持仓记录Hash
	PlanningTrades []string      `json:"planningTrades"` // 计划交易Hash
}

func (strat *Strategy) IsPrivate() bool {
	return strat.State.String() == "PRIVATE"
}

func (strat *Strategy) IsPublic() bool {
	return strat.State.String() == "PUBLIC"
}

// GetSplitKey returns values which should be used to form key
func (strat *Strategy) GetSplitKey() []string {
	// return ledgerapi.SplitKey(strat.ID)
	return []string{strat.Provider, strat.ID}
}

// Serialize formats the commercial paper as JSON bytes
func (strat *Strategy) Serialize() ([]byte, error) {
	return json.Marshal(strat)
}

func DeserializeStrategy(bytes []byte, strategy *Strategy) error {
	err := json.Unmarshal(bytes, strategy)

	if err != nil {
		return fmt.Errorf("Error deserializing strategy. %s", err.Error())
	}

	return nil
}
