package storage

import (
	"github.com/ProtoconNet/mitum-currency/v3/common"
	"github.com/ProtoconNet/mitum-currency/v3/operation/test"
	"github.com/ProtoconNet/mitum-currency/v3/state/extension"
	currencytypes "github.com/ProtoconNet/mitum-currency/v3/types"
	statestorage "github.com/ProtoconNet/mitum-storage/state"
	"github.com/ProtoconNet/mitum-storage/types"
	"github.com/ProtoconNet/mitum2/base"
	"github.com/ProtoconNet/mitum2/util"
)

type TestCreateDataProcessor struct {
	*test.BaseTestOperationProcessorNoItem[CreateData]
}

func NewTestCreateDataProcessor(tp *test.TestProcessor) TestCreateDataProcessor {
	t := test.NewBaseTestOperationProcessorNoItem[CreateData](tp)
	return TestCreateDataProcessor{BaseTestOperationProcessorNoItem: &t}
}

func (t *TestCreateDataProcessor) Create() *TestCreateDataProcessor {
	t.Opr, _ = NewCreateDataProcessor()(
		base.GenesisHeight,
		t.GetStateFunc,
		nil, nil,
	)
	return t
}

func (t *TestCreateDataProcessor) SetCurrency(
	cid string, am int64, addr base.Address, target []currencytypes.CurrencyID, instate bool,
) *TestCreateDataProcessor {
	t.BaseTestOperationProcessorNoItem.SetCurrency(cid, am, addr, target, instate)

	return t
}

func (t *TestCreateDataProcessor) SetAmount(
	am int64, cid currencytypes.CurrencyID, target []currencytypes.Amount,
) *TestCreateDataProcessor {
	t.BaseTestOperationProcessorNoItem.SetAmount(am, cid, target)

	return t
}

func (t *TestCreateDataProcessor) SetContractAccount(
	owner base.Address, priv string, amount int64, cid currencytypes.CurrencyID, target []test.Account, inState bool,
) *TestCreateDataProcessor {
	t.BaseTestOperationProcessorNoItem.SetContractAccount(owner, priv, amount, cid, target, inState)

	return t
}

func (t *TestCreateDataProcessor) SetAccount(
	priv string, amount int64, cid currencytypes.CurrencyID, target []test.Account, inState bool,
) *TestCreateDataProcessor {
	t.BaseTestOperationProcessorNoItem.SetAccount(priv, amount, cid, target, inState)

	return t
}

func (t *TestCreateDataProcessor) LoadOperation(fileName string,
) *TestCreateDataProcessor {
	t.BaseTestOperationProcessorNoItem.LoadOperation(fileName)

	return t
}

func (t *TestCreateDataProcessor) Print(fileName string,
) *TestCreateDataProcessor {
	t.BaseTestOperationProcessorNoItem.Print(fileName)

	return t
}

func (t *TestCreateDataProcessor) RunPreProcess() *TestCreateDataProcessor {
	t.BaseTestOperationProcessorNoItem.RunPreProcess()

	return t
}

func (t *TestCreateDataProcessor) RunProcess() *TestCreateDataProcessor {
	t.BaseTestOperationProcessorNoItem.RunProcess()

	return t
}

func (t *TestCreateDataProcessor) SetService(
	contract base.Address,
	pid string,
) *TestCreateDataProcessor {
	design := types.NewDesign(pid)

	st := common.NewBaseState(base.Height(1), statestorage.DesignStateKey(contract), statestorage.NewDesignStateValue(design), nil, []util.Hash{})
	t.SetState(st, true)

	cst, found, _ := t.MockGetter.Get(extension.StateKeyContractAccount(contract))
	if !found {
		panic("contract account not set")
	}
	status, err := extension.StateContractAccountValue(cst)
	if err != nil {
		panic(err)
	}

	nstatus := status.SetIsActive(true)
	cState := common.NewBaseState(base.Height(1), extension.StateKeyContractAccount(contract), extension.NewContractAccountStateValue(nstatus), nil, []util.Hash{})
	t.SetState(cState, true)

	return t
}

func (t *TestCreateDataProcessor) MakeOperation(
	sender base.Address,
	privatekey base.Privatekey,
	contract base.Address,
	key, value string,
	currency currencytypes.CurrencyID,
) *TestCreateDataProcessor {
	op, _ := NewCreateData(
		NewCreateDataFact(
			[]byte("token"),
			sender,
			contract,
			key,
			value,
			currency,
		))
	_ = op.Sign(privatekey, t.NetworkID)
	t.Op = op

	return t
}
