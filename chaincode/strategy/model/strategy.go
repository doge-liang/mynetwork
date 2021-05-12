package chaincode

import (
	"encoding/json"
	"fmt"
	ledgerapi "mynetwork/chaincode/strategy/ledgerapi"
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

// 策略公开的数据
type Strategy struct {
	ID           string   `json:"ID"`           // 策略 ID
	Name         string   `json:"name"`         // 策略名
	Provider     string   `json:"provider"`     // 发布者
	MaxDrawdown  float64  `json:"maxDrawdown"`  // 最大回撤
	AnnualReturn float64  `json:"annualReturn"` // 年化收益率
	SharpeRatio  float64  `json:"sharpeRatio"`  // 夏普率
	Subscribers  []string `json:"subscribers"`  // 订阅者证书列表
	State        State    `json:"state"`        // 是否公开
	// Trades         []Trade         `json:"trades"`         // 交易记录
	// PlanningTrades []PlanningTrade `json:"planningTrades"` // 计划交易
	// Positions      []Position      `json:"positions"`      // 持仓记录
}

func (strat *Strategy) IsPrivate() bool {
	return strat.State == PRIVATE
}

func (strat *Strategy) IsPublic() bool {
	return strat.State == PUBLIC
}

func (strat *Strategy) SetPrivate() {
	strat.State = PRIVATE
}

func (strat *Strategy) SetPublic() {
	strat.State = PUBLIC
}

// 继承自 StateInterface
func (strat *Strategy) GetSplitKey() []string {
	return []string{strat.ID}
}

// 继承自 StateInterface
func (strat *Strategy) Serialize() ([]byte, error) {
	return json.Marshal(strat)
}

type PrivateStrategy struct {
	*Strategy
	Trades         []*Trade `json:"trades"`         // 交易记录Hash
	Positions      string   `json:"positions"`      // 持仓记录Hash
	PlanningTrades string   `json:"planningTrades"` // 计划交易Hash
	// Provider     string  `json:"provider"`     // 发布者
}

func GetStrategyKey(strategyID string) string {
	return ledgerapi.MakeKey(strategyID)
}

func DeserializeStrategy(bytes []byte, strategy *Strategy) error {
	err := json.Unmarshal(bytes, strategy)

	if err != nil {
		return fmt.Errorf("Error deserializing strategy. %s", err.Error())
	}

	return nil
}
