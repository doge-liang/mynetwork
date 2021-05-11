package ledgerapi

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type StateListInterface interface {
	AddState(StateInterface) error
	AddStateIn(string, StateInterface) error

	GetState(string, StateInterface) error
	GetStateIn(string, string, StateInterface) error

	GetStateByPartialCompositeKey([]string, []StateInterface) error

	UpdateState(StateInterface) error
	UpdateStateIn(string, StateInterface) error

	DeleteState(string) error
	DeleteStateIn(string, string) error
}

type StateList struct {
	Ctx         contractapi.TransactionContextInterface
	Name        string
	Deserialize func([]byte, StateInterface) error
}

// 添加数据到公开账本
func (sl *StateList) AddState(state StateInterface) error {
	key, _ := sl.Ctx.GetStub().CreateCompositeKey(sl.Name, state.GetSplitKey())
	data, err := state.Serialize()

	if err != nil {
		return err
	}

	return sl.Ctx.GetStub().PutState(key, data)
}

// 添加数据到私有数据集
func (sl *StateList) AddStateIn(collection string, state StateInterface) error {
	key, _ := sl.Ctx.GetStub().CreateCompositeKey(sl.Name, state.GetSplitKey())
	data, err := state.Serialize()

	if err != nil {
		return err
	}

	return sl.Ctx.GetStub().PutPrivateData(collection, key, data)
}

// 公开账本获取数据
func (sl *StateList) GetState(key string, state StateInterface) error {
	ledgerKey, _ := sl.Ctx.GetStub().CreateCompositeKey(sl.Name, SplitKey(key))
	data, err := sl.Ctx.GetStub().GetState(ledgerKey)

	if err != nil {
		return err
	} else if data == nil {
		return fmt.Errorf("No state found for %s", key)
	}

	return sl.Deserialize(data, state)
}

// 根据部分 ID 获取信息
func (sl *StateList) GetStateByPartialCompositeKey(keys []string, states []StateInterface) error {
	resultsIterator, err := sl.Ctx.GetStub().GetStateByPartialCompositeKey(sl.Name, keys)
	if err != nil {
		return err
	}
	defer resultsIterator.Close()

	// var states []*StateInterface
	for resultsIterator.HasNext() {
		resultJSON, err := resultsIterator.Next()

		if err != nil {
			return err
		}

		var state StateInterface
		sl.Deserialize(resultJSON.Value, state)

		states = append(states, state)
	}

	return nil
}

// 从隐私数据集中获取数据
func (sl *StateList) GetStateIn(key string, collection string, state StateInterface) error {
	ledgerKey, _ := sl.Ctx.GetStub().CreateCompositeKey(sl.Name, SplitKey(key))
	data, err := sl.Ctx.GetStub().GetPrivateData(collection, ledgerKey)

	if err != nil {
		return err
	} else if data == nil {
		return fmt.Errorf("No state found for %s in %s collection", key, collection)
	}

	return sl.Deserialize(data, state)
}

// 更新账本状态
func (sl *StateList) UpdateState(state StateInterface) error {
	return sl.AddState(state)
}

// 更新私有数据集中的账本状态
func (sl *StateList) UpdateStateIn(collection string, state StateInterface) error {
	return sl.AddStateIn(collection, state)
}

// 从公共账本删除数据
func (sl *StateList) DeleteState(key string) error {
	return sl.Ctx.GetStub().DelState(key)
}

// 从私有数据集中删除数据
func (sl *StateList) DeleteStateIn(collection string, key string) error {
	return sl.Ctx.GetStub().DelPrivateData(collection, key)
}
