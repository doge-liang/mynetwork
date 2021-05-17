package ledgerapi

import (
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type StateListInterface interface {
	AddState(StateInterface) error
	GetState(string, StateInterface) error
	UpdateState(StateInterface) error
	DelState(string) error

	GetStateByPartialCompositeKey([]string) (shim.StateQueryIteratorInterface, error)
	GetStateByPartialCompositeKeyWithPagination([]string, int32, string) (shim.StateQueryIteratorInterface, string, error)
}

// StateList useful for managing putting data in and out
// of the ledger. Implementation of StateListInterface
type StateList struct {
	Ctx         contractapi.TransactionContextInterface
	Name        string
	Deserialize func([]byte, StateInterface) error
}

// AddState puts state into world state
func (sl *StateList) AddState(state StateInterface) error {
	key, _ := sl.Ctx.GetStub().CreateCompositeKey(sl.Name, state.GetSplitKey())
	data, err := state.Serialize()

	if err != nil {
		return err
	}

	return sl.Ctx.GetStub().PutState(key, data)
}

// GetState returns state from world state. Unmarshalls the JSON
// into passed state. Key is the split key value used in Add/Update
// joined using a colon
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

func (sl *StateList) GetStateByPartialCompositeKey(keys []string) (shim.StateQueryIteratorInterface, error) {
	resultsIterator, err := sl.Ctx.GetStub().GetStateByPartialCompositeKey(sl.Name, keys)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	return resultsIterator, nil
}

// UpdateState puts state into world state. Same as AddState but
// separate as semantically different
func (sl *StateList) UpdateState(state StateInterface) error {
	return sl.AddState(state)
}

func (sl *StateList) DelState(key string) error {
	ledgerKey, _ := sl.Ctx.GetStub().CreateCompositeKey(sl.Name, SplitKey(key))
	err := sl.Ctx.GetStub().DelState(ledgerKey)
	if err != nil {
		return err
	}
	return nil
}

func (sl *StateList) GetStateByPartialCompositeKeyWithPagination(keys []string, pageSize int32, bookmark string) (shim.StateQueryIteratorInterface, string, error) {
	resultsIterator, metadata, err := sl.Ctx.GetStub().GetStateByPartialCompositeKeyWithPagination(sl.Name, keys, pageSize, bookmark)
	if err != nil {
		return nil, "", err
	}
	return resultsIterator, metadata.Bookmark, err
}
