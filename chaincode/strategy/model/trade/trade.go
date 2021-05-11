package trade

import (
	"encoding/json"
	"fmt"
	ledgerapi "mynetwork/chaincode/strategy/ledger-api"
	"time"
)

type Trade struct {
	ID         string    `json:"ID"`         // 交易 ID
	StockID    string    `json:"stockID"`    // 交易股票
	Amount     float64   `json:"amount"`     // 交易份额（买卖用正负来表示）
	Commission float64   `json:"commission"` // 交易佣金
	DateTime   time.Time `json:"dateTime"`   // 交易时间戳
	Price      float64   `json:"price"`      // 成交价
	StrategyID string    `json:"strategyID"` // 关联策略 ID
}

// type Trades struct {
// 	StrategyID string  `json:"strategyID"`
// 	Trades     []Trade `json:"trades"` // 交易记录
// }

func GetTradesKey(StrategyID string, id string) string {
	// keyParts := ledgerapi.SplitKey(strategyKey)
	// return ledgerapi.MakeKey(consts.STRATEGY, consts.TRADES, keyParts[len(keyParts)-1])
	return ledgerapi.MakeKey(StrategyID, id)
}

// GetSplitKey returns values which should be used to form key
func (t *Trade) GetSplitKey() []string {
	// return ledgerapi.SplitKey(t.ID)
	return []string{t.StrategyID, t.ID}
}

// Serialize formats the commercial paper as JSON bytes
func (t *Trade) Serialize() ([]byte, error) {
	return json.Marshal(t)
}

func DeserializeTrade(bytes []byte, trade *Trade) error {
	err := json.Unmarshal(bytes, trade)

	if err != nil {
		return fmt.Errorf("Error deserializing trade. %s", err.Error())
	}

	return nil
}
