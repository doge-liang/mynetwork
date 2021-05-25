package contract

import (
	"fmt"
	"log"
	. "mynetwork/chaincode/strategy/model/list"
	. "mynetwork/chaincode/strategy/utils"
	"sort"
)

// 订阅
func (s *SmartContract) Subscribe(ctx TransactionContextInterface, strategyID string) error {
	strat, err := ctx.GetStrategyList().GetStrategy(strategyID)
	if err != nil {
		return err
	}

	clientID, err := ctx.GetClientIdentity().GetID()
	if err != nil {
		return err
	}

	v, f, err := ctx.GetClientIdentity().GetAttributeValue("strategy.role")
	if err != nil {
		return err
	}
	if f {
		log.Printf("Receive the request with strategy.role : %v", v)
	} else {
		log.Print("Did not receieve the attr")
	}
	if In(clientID, strat.Subscribers) {
		return fmt.Errorf("You have been subscribed.")
	}
	strat.Subscribers = append(strat.Subscribers, clientID)

	return ctx.GetStrategyList().UpdateStrategy(strat)
}

// 取消订阅
func (s *SmartContract) UnSubscribe(ctx TransactionContextInterface, strategyID string) error {
	strat, err := ctx.GetStrategyList().GetStrategy(strategyID)
	if err != nil {
		return err
	}
	clientID, err := ctx.GetClientIdentity().GetID()
	if err != nil {
		return err
	}

	if !In(clientID, strat.Subscribers) {
		return fmt.Errorf("You have not been subscribed.")
	}
	pos := sort.SearchStrings(strat.Subscribers, clientID)
	strat.Subscribers = append(strat.Subscribers[:pos], strat.Subscribers[(pos+1):]...)
	return ctx.GetStrategyList().UpdateStrategy(strat)
}
