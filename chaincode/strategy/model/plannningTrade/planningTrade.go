package planningTrade

import (
	"encoding/json"
	"fmt"
	ledgerapi "mynetwork/chaincode/strategy/ledger-api"
)

type PlanningTrade struct {
	ID         string  `json:"ID"`
	StockID    string  `json:"stockID"`    // 交易股票
	Amount     float64 `json:"amount"`     // 交易份额（买卖用正负来表示）
	StrategyID string  `json:"strategyID"` // 关联策略ID
}

// type PlanningTrades struct {
// 	StrategyID     string          `json:"strategyID"`
// 	PlanningTrades []PlanningTrade `json:"planningTrades"` // 计划交易
// }

func GetPlanningTradesKey(strategyKey string, id string) string {
	// keyParts := ledgerapi.SplitKey(strategyKey)
	// return ledgerapi.MakeKey(consts.STRATEGY, consts.PLANNINGTRADES, keyParts[len(keyParts)-1])
	return ledgerapi.MakeKey(strategyKey, id)
}

// GetSplitKey returns values which should be used to form key
func (pt *PlanningTrade) GetSplitKey() []string {
	return []string{pt.StrategyID, pt.ID}
}

// Serialize formats the commercial paper as JSON bytes
func (pt *PlanningTrade) Serialize() ([]byte, error) {
	return json.Marshal(pt)
}

func DeserializePlanningTrade(bytes []byte, planningTrade *PlanningTrade) error {
	err := json.Unmarshal(bytes, planningTrade)

	if err != nil {
		return fmt.Errorf("Error deserializing planning trade. %s", err.Error())
	}

	return nil
}
