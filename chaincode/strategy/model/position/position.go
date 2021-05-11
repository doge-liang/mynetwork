package position

import (
	"encoding/json"
	"fmt"
	ledgerapi "mynetwork/chaincode/strategy/ledger-api"
)

type Position struct {
	ID         string  `json:"ID"`         // 仓位 ID
	StockID    string  `json:"stockID"`    // 股票代码
	Value      float64 `json:"value"`      // 仓位
	StrategyID string  `json:"strategyID"` // 关联策略 ID
	// Price      float64 `json:"Price"`      // 现有股价
}

// type Positions struct {
// 	Positions []*Position `json:"positions"` // 持仓记录
// }

func GetPositionsKey(strategyID string, id string) string {
	// keyParts := ledgerapi.SplitKey(strategyKey)
	// return ledgerapi.MakeKey(consts.STRATEGY, consts.POSITIONS, keyParts[len(keyParts)-1])
	return ledgerapi.MakeKey(strategyID, id)
}

// GetSplitKey returns values which should be used to form key
func (p *Position) GetSplitKey() []string {
	// return ledgerapi.SplitKey(pos.ID)
	return []string{p.StrategyID, p.ID}
}

// Serialize formats the commercial paper as JSON bytes
func (p *Position) Serialize() ([]byte, error) {
	return json.Marshal(p)
}

func DeserializePosition(bytes []byte, position *Position) error {
	err := json.Unmarshal(bytes, position)

	if err != nil {
		return fmt.Errorf("Error deserializing trade. %s", err.Error())
	}

	return nil
}
