package storage

import (
	"github.com/ProtoconNet/mitum-currency/v3/common"
	currencytypes "github.com/ProtoconNet/mitum-currency/v3/types"
	"github.com/ProtoconNet/mitum-storage/types"
	mitumbase "github.com/ProtoconNet/mitum2/base"
	"github.com/ProtoconNet/mitum2/util"
	"github.com/ProtoconNet/mitum2/util/hint"
	"github.com/ProtoconNet/mitum2/util/valuehash"
	"github.com/pkg/errors"
)

var (
	CreateDataFactHint = hint.MustNewHint("mitum-storage-create-data-operation-fact-v0.0.1")
	CreateDataHint     = hint.MustNewHint("mitum-storage-create-data-operation-v0.0.1")
)

type CreateDataFact struct {
	mitumbase.BaseFact
	sender    mitumbase.Address
	contract  mitumbase.Address
	dataKey   string
	dataValue string
	currency  currencytypes.CurrencyID
}

func NewCreateDataFact(
	token []byte, sender, contract mitumbase.Address,
	key, value string, currency currencytypes.CurrencyID) CreateDataFact {
	bf := mitumbase.NewBaseFact(CreateDataFactHint, token)
	fact := CreateDataFact{
		BaseFact:  bf,
		sender:    sender,
		contract:  contract,
		dataKey:   key,
		dataValue: value,
		currency:  currency,
	}

	fact.SetHash(fact.GenerateHash())
	return fact
}

func (fact CreateDataFact) IsValid(b []byte) error {
	if len(fact.dataKey) < 1 || len(fact.dataKey) > types.MaxKeyLen {
		return common.ErrFactInvalid.Wrap(
			common.ErrValOOR.Wrap(
				errors.Errorf("invalid data key length %v < 1 or %v > %v", len(fact.dataKey), len(fact.dataKey), types.MaxKeyLen)))
	}

	if !currencytypes.ReValidSpcecialCh.Match([]byte(fact.dataKey)) {
		return common.ErrFactInvalid.Wrap(common.ErrValueInvalid.Wrap(errors.Errorf("date key %s, must match regex `^[^\\s:/?#\\[\\]$@]*$`", fact.dataKey)))
	}

	if len(fact.dataValue) < 1 || len(fact.dataValue) > types.MaxDataLen {
		return common.ErrFactInvalid.Wrap(
			common.ErrValOOR.Wrap(
				errors.Errorf("invalid data value length %v < 1 or %v > %v", len(fact.dataValue), len(fact.dataValue), types.MaxDataLen)))
	}

	if fact.sender.Equal(fact.contract) {
		return common.ErrFactInvalid.Wrap(
			common.ErrSelfTarget.Wrap(errors.Errorf("sender %v is same with contract account", fact.sender)))
	}

	if err := util.CheckIsValiders(nil, false,
		fact.BaseHinter,
		fact.sender,
		fact.contract,
		fact.currency,
	); err != nil {
		return common.ErrFactInvalid.Wrap(err)
	}

	if err := common.IsValidOperationFact(fact, b); err != nil {
		return common.ErrFactInvalid.Wrap(err)
	}

	return nil
}

func (fact CreateDataFact) Hash() util.Hash {
	return fact.BaseFact.Hash()
}

func (fact CreateDataFact) GenerateHash() util.Hash {
	return valuehash.NewSHA256(fact.Bytes())
}

func (fact CreateDataFact) Bytes() []byte {
	return util.ConcatBytesSlice(
		fact.Token(),
		fact.sender.Bytes(),
		fact.contract.Bytes(),
		[]byte(fact.dataKey),
		[]byte(fact.dataValue),
		fact.currency.Bytes(),
	)
}

func (fact CreateDataFact) Token() mitumbase.Token {
	return fact.BaseFact.Token()
}

func (fact CreateDataFact) Sender() mitumbase.Address {
	return fact.sender
}

func (fact CreateDataFact) Contract() mitumbase.Address {
	return fact.contract
}

func (fact CreateDataFact) DataKey() string {
	return fact.dataKey
}

func (fact CreateDataFact) DataValue() string {
	return fact.dataValue
}

func (fact CreateDataFact) Currency() currencytypes.CurrencyID {
	return fact.currency
}

func (fact CreateDataFact) Addresses() ([]mitumbase.Address, error) {
	as := []mitumbase.Address{fact.sender}

	return as, nil
}

type CreateData struct {
	common.BaseOperation
}

func NewCreateData(fact CreateDataFact) (CreateData, error) {
	return CreateData{BaseOperation: common.NewBaseOperation(CreateDataHint, fact)}, nil
}
