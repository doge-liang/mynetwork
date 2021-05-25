package chaincode

import (
	"encoding/json"
	"fmt"
	ledgerapi "mynetwork/chaincode/strategy/ledgerapi"
)

type Position struct {
	ID         string  `json:"id"`         // 仓位 ID
	StrategyID string  `json:"strategyID"` // 关联策略 ID
	StockID    string  `json:"stockID"`    // 股票代码
	Value      float64 `json:"value"`      // 仓位
	// Price      float64 `json:"price"`      // 现有股价
}

// 继承自 StateInterface
func (p *Position) GetSplitKey() []string {
	return []string{p.StrategyID, p.ID}
}

// 继承自 StateInterface
func (p *Position) Serialize() ([]byte, error) {
	return json.Marshal(p)
}

func GetPositionsKey(strategyKey string, id string) string {
	return ledgerapi.MakeKey(strategyKey, id)
}

func DeserializePosition(bytes []byte, position *Position) error {
	err := json.Unmarshal(bytes, position)

	if err != nil {
		return fmt.Errorf("Error deserializing position. %s", err.Error())
	}

	return nil
}

// 私有策略需要维护一个公共账本的 key 值列表供查询哈希值进行验证
type PositionPublic struct {
	ID         string `json:"id"`
	StrategyID string `json:"strategyID"` // 关联策略 ID
}

// 继承自 StateInterface
func (p *PositionPublic) GetSplitKey() []string {
	return []string{p.StrategyID, p.ID}
}

// 继承自 StateInterface
func (p *PositionPublic) Serialize() ([]byte, error) {
	return json.Marshal(p)
}

func DeserializePositionPublic(bytes []byte, pp *PositionPublic) error {
	err := json.Unmarshal(bytes, pp)

	if err != nil {
		return fmt.Errorf("Error deserializing positionPublic. %s", err.Error())
	}

	return nil
}

type PositionHash struct {
	ID       string `json:"id"`
	Hashcode string `json:"hashcode"`
}

type PositionsOutput struct {
	Positions     []*Position     `json:"positions"`
	PositionsHash []*PositionHash `json:"positionsHash"`
}
