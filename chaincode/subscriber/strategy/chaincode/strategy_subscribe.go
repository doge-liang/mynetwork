package chaincode

import (
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract definition
type SmartContract struct {
	contractapi.Contract
}

// 策略的交易记录保持 Provider 私有，用户发起 subscribe 后，再由
type Trade struct {
	ID         string    `json:"ID"`         // 交易 ID
	StrategyID string    `json:"strategyID"` // 关联策略 ID
	StockID    string    `json:"stockID"`    // 交易股票
	Amount     float64   `json:"amount"`     // 交易份额（买卖用正负来表示）
	Commission float64   `json:"commission"` // 交易佣金
	DateTime   time.Time `json:"dateTime"`   // 交易时间
	Price      float64   `json:"price"`      // 成交价
}

// 策略的现有仓位保持私有
type Position struct {
	ID         string  `json:"ID"`         // 仓位 ID
	StockID    string  `json:"stockID"`    // 股票代码
	StrategyID string  `json:"strategyID"` // 关联策略 ID
	Price      float64 `json:"Price"`      // 现有股价
	Amount     float64 `json:"amount"`     // 仓位
}

// 策略公开的数据
type Strategy struct {
	ID           string  `json:"ID"`           // 策略 ID
	Name         string  `json:"name"`         // 策略名
	Provider     string  `json:"provider"`     // 发布者
	MaxDrawdown  float64 `json:"maxDrawdown"`  // 最大回撤
	AnnualReturn float64 `json:"annualReturn"` // 年化收益率
}

type Subscription struct {
	StrategyID  string   `json:"strategyID"`  // 策略 ID
	Subscribers []string `json:"subscribers"` // 订阅者列表
}

