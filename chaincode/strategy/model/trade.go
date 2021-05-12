package chaincode

import (
	"encoding/json"
	"fmt"
	ledgerapi "mynetwork/chaincode/strategy/ledgerapi"
	"time"
)

type Trade struct {
	ID         string    `json:"ID"`         // 交易 ID
	StrategyID string    `json:"strategyID"` // 关联策略 ID
	StockID    string    `json:"stockID"`    // 交易股票
	Amount     float64   `json:"amount"`     // 交易份额（买卖用正负来表示）
	Commission float64   `json:"commission"` // 交易佣金
	DateTime   time.Time `json:"dateTime"`   // 交易时间戳
	Price      float64   `json:"price"`      // 成交价
}

// 继承自 StateInterface
func (t *Trade) GetSplitKey() []string {
	return []string{t.StrategyID, t.ID}
}

// 继承自 StateInterface
func (t *Trade) Serialize() ([]byte, error) {
	return json.Marshal(t)
}

// 策略的交易记录保持 Provider 私有，用户发起 subscribe 后，再由 管理员操作资产转移
type Trades struct {
	StrategyID string  `json:"strategyID"`
	Trades     []Trade `json:"trades"` // 交易记录
}

func GetTradesKey(strategyKey string, id string) string {
	return ledgerapi.MakeKey(strategyKey, id)
}

func DeserializeTrade(bytes []byte, trade *Trade) error {
	err := json.Unmarshal(bytes, trade)

	if err != nil {
		return fmt.Errorf("Error deserializing trade. %s", err.Error())
	}

	return nil
}
