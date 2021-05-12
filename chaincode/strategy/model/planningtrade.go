package chaincode

import (
	"encoding/json"
	"fmt"
	ledgerapi "mynetwork/chaincode/strategy/ledgerapi"
)

type PlanningTrade struct {
	ID         string  `json:"ID"`         // 计划交易ID
	StrategyID string  `json:"strategyID"` // 关联策略ID
	StockID    string  `json:"stockID"`    // 交易股票
	Amount     float64 `json:"amount"`     // 交易份额（买卖用正负来表示）
}

// 继承自 StateInterface
func (pt *PlanningTrade) GetSplitKey() []string {
	return []string{pt.StrategyID, pt.ID}
}

// 继承自 StateInterface
func (pt *PlanningTrade) Serialize() ([]byte, error) {
	return json.Marshal(pt)
}

func GetPlanningTradesKey(strategyKey string, id string) string {
	return ledgerapi.MakeKey(strategyKey, id)
}

func DeserializePlanningTrade(bytes []byte, planningTrade *PlanningTrade) error {
	err := json.Unmarshal(bytes, planningTrade)

	if err != nil {
		return fmt.Errorf("Error deserializing planning trade. %s", err.Error())
	}

	return nil
}

type PlanningTradePublic struct {
	ID         string `json:"ID"` // 计划交易ID
	StrategyID string `json:"strategyID"`
}

// 继承自 StateInterface
func (ptp *PlanningTradePublic) GetSplitKey() []string {
	return []string{ptp.StrategyID, ptp.ID}
}

// 继承自 StateInterface
func (ptp *PlanningTradePublic) Serialize() ([]byte, error) {
	return json.Marshal(ptp)
}

func DeserializePlanningTradePublic(bytes []byte, ptp *PlanningTradePublic) error {
	err := json.Unmarshal(bytes, ptp)

	if err != nil {
		return fmt.Errorf("Error deserializing planningTradePublic. %s", err.Error())
	}

	return nil
}

type PlanningTradeHash struct {
	ID       string `json:"ID"`
	Hashcode string `json:"hashcode"`
}

type PlanningTradesOutput struct {
	PlanningTrades     []*PlanningTrade     `json:"planningTrades"`
	PlanningTradesHash []*PlanningTradeHash `json:"planningTradesHash"`
}
