package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract definition
type SmartContract struct {
	contractapi.Contract
}

type Trade struct {
	ID         string    `json:"ID"`         // 交易id
	StockID    string    `json:"stockID"`    // 交易股票
	Amount     float64   `json:"amount"`     // 交易份额（买卖用正负来表示）
	Commission float64   `json:"commission"` // 交易佣金
	DateTime   time.Time `json:"dateTime"`   // 交易时间
	Price      float64   `json:"price"`      // 成交价
}

type Position struct {
	ID     string  `json:"ID"`     // 股票代码
	Price  float64 `json:"Price"`  // 现有股价
	Amount float64 `json:"amount"` // 仓位
}

type Strategy struct {
	ID           string     `json:"ID"`           // 策略 ID
	Name         string     `json:"name"`         // 策略名
	Provider     string     `json:"provider"`     // 发布者
	Subscribers  []string   `json:"subscribers"`  // 订阅者列表
	MaxDrawdown  float64    `json:"maxDrawdown"`  // 最大回撤
	AnnualReturn float64    `json:"annualReturn"` // 年化收益率
	Trades       []Trade    `json:"trades"`       // 交易记录
	Positions    []Position `json:"positions"`    // 持仓
}

// InitLedger adds a base set of strategies to the ledger
func (s *SmartContract) Init(ctx contractapi.TransactionContextInterface) error {
	strategies := []Strategy{
		{
			ID:           "1",
			Name:         "测试策略名",
			Provider:     "string",
			MaxDrawdown:  0.1,
			AnnualReturn: 0.2,
			Trades: []Trade{
				{
					ID:         "测试股票",
					Amount:     10.0,
					Commission: 10.0,
					DateTime:   time.Now(),
					Price:      100.0,
				},
			},
			Positions: []Position{
				{
					ID:     "测试股票",
					Price:  100.0,
					Amount: 10.0,
				},
			},
		},
	}
	for _, strategy := range strategies {
		strategyJSON, err := json.Marshal(strategy)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(strategy.ID, strategyJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}
	return nil
}

// CreateStrategy issues a new strategy to the world state with given details.
func (s *SmartContract) CreateStrategy(
	ctx contractapi.TransactionContextInterface,
	id string,
	name string,
	provider string,
	maxdrawdown float64,
	annualreturn float64,
) error {

	exists, err := s.StrategyExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the strategy %s already exists", id)
	}
	strategy := Strategy{
		ID:           id,
		Name:         name,
		Provider:     provider,
		MaxDrawdown:  maxdrawdown,
		AnnualReturn: annualreturn,
	}

	strategyJSON, err := json.Marshal(strategy)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(id, strategyJSON)
}

// GetAllStrategies returns all strategies found in world state
func (s *SmartContract) GetAllStrategies(ctx contractapi.TransactionContextInterface) ([]*Strategy, error) {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all strategies in the chaincode namespace.
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var strategies []*Strategy
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var strategy Strategy
		err = json.Unmarshal(queryResponse.Value, &strategy)
		if err != nil {
			return nil, err
		}
		strategies = append(strategies, &strategy)
	}

	return strategies, nil
}

func (s *SmartContract) subscribe(ctx contractapi.TransactionContextInterface, id string) {

}

// UpdateStrategy updates an existing strategy in the world state with provided parameters.
func (s *SmartContract) UpdateStrategy(
	ctx contractapi.TransactionContextInterface,
	id string,
	name string,
	provider string,
	maxdrawdown float64,
	annualreturn float64,
	trades []Trade) error {
	exists, err := s.StrategyExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the strategy %s does not exist", id)
	}

	// overwriting original strategy with new strategy
	strategy := Strategy{
		ID:           id,
		Name:         name,
		Provider:     provider,
		MaxDrawdown:  maxdrawdown,
		AnnualReturn: annualreturn,
		Trades:       trades,
	}
	strategyJSON, err := json.Marshal(strategy)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, strategyJSON)
}

// DeleteStrategy deletes an given strategy from the world state.
func (s *SmartContract) DeleteStrategy(ctx contractapi.TransactionContextInterface, id string) error {
	exists, err := s.StrategyExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the strategy %s does not exist", id)
	}

	return ctx.GetStub().DelState(id)
}

// 用 ID 检查策略是否存在
func (s *SmartContract) StrategyExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	strategyJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return strategyJSON != nil, nil
}

// ReadStrategy returns the strategy stored in the world state with given id.
func (s *SmartContract) ReadStrategy(ctx contractapi.TransactionContextInterface, id string) (*Strategy, error) {
	strategyJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if strategyJSON == nil {
		return nil, fmt.Errorf("the strategy %s does not exist", id)
	}

	var strategy Strategy
	err = json.Unmarshal(strategyJSON, &strategy)
	if err != nil {
		return nil, err
	}

	return &strategy, nil
}

func main() {
	strategyChaincode, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		log.Panicf("Error creating strategy-subscribe chaincode: %v", err)
	}

	if err := strategyChaincode.Start(); err != nil {
		log.Panicf("Error starting strategy-subscribe chaincode: %v", err)
	}
}
