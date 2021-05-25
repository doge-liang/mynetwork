package ledgerapi

import (
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type PrivateStateListInterface interface {
	SetCollection(string)

	AddStateIn(StateInterface) error
	GetStateIn(string, StateInterface) error
	GetStateHash(string, *string) error
	UpdateStateIn(StateInterface) error
	DelStateIn(string) error

	GetPrivateDataByPartialCompositeKey([]string) (shim.StateQueryIteratorInterface, error)
}

type PrivateStateList struct {
	Ctx         contractapi.TransactionContextInterface
	Name        string
	Deserialize func([]byte, StateInterface) error
	Collection  string
}

func (psl *PrivateStateList) SetCollection(collection string) {
	psl.Collection = collection
}

func (psl *PrivateStateList) AddStateIn(state StateInterface) error {
	key, _ := psl.Ctx.GetStub().CreateCompositeKey(MakeKey(psl.Name, psl.Collection), state.GetSplitKey())
	data, err := state.Serialize()

	if err != nil {
		return err
	}

	return psl.Ctx.GetStub().PutPrivateData(psl.Collection, key, data)
}

func (psl *PrivateStateList) GetStateIn(key string, state StateInterface) error {
	ledgerKey, _ := psl.Ctx.GetStub().CreateCompositeKey(MakeKey(psl.Name, psl.Collection), SplitKey(key))
	data, err := psl.Ctx.GetStub().GetPrivateData(psl.Collection, ledgerKey)

	if err != nil {
		return err
	} else if data == nil {
		return fmt.Errorf("No state found for %s", key)
	}

	return psl.Deserialize(data, state)
}

func (psl *PrivateStateList) GetPrivateDataByPartialCompositeKey(keys []string) (shim.StateQueryIteratorInterface, error) {
	resultsIterator, err := psl.Ctx.GetStub().GetPrivateDataByPartialCompositeKey(psl.Collection, MakeKey(psl.Name, psl.Collection), keys)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	return resultsIterator, nil
}

func (psl *PrivateStateList) UpdateStateIn(state StateInterface) error {
	return psl.AddStateIn(state)
}

func (psl *PrivateStateList) DelStateIn(key string) error {
	ledgerKey, _ := psl.Ctx.GetStub().CreateCompositeKey(MakeKey(psl.Name, psl.Collection), SplitKey(key))
	err := psl.Ctx.GetStub().DelPrivateData(psl.Collection, ledgerKey)
	if err != nil {
		return err
	}
	return nil
}

func (psl *PrivateStateList) GetStateHash(key string, hashcode *string) error {
	ledgerKey, _ := psl.Ctx.GetStub().CreateCompositeKey(MakeKey(psl.Name, psl.Collection), SplitKey(key))
	data, err := psl.Ctx.GetStub().GetPrivateDataHash(psl.Collection, ledgerKey)

	if err != nil {
		return err
	} else if data == nil {
		return fmt.Errorf("No state found for %s", key)
	}

	*hashcode = fmt.Sprintf("%x", data)

	return nil
}
