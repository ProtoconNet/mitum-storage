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
	DeleteDataFactHint = hint.MustNewHint("mitum-storage-delete-data-operation-fact-v0.0.1")
	DeleteDataHint     = hint.MustNewHint("mitum-storage-delete-data-operation-v0.0.1")
)

type DeleteDataFact struct {
	mitumbase.BaseFact
	sender   mitumbase.Address
	contract mitumbase.Address
	dataKey  string
	currency currencytypes.CurrencyID
}

func NewDeleteDataFact(
	token []byte, sender, contract mitumbase.Address,
	key string, currency currencytypes.CurrencyID) DeleteDataFact {
	bf := mitumbase.NewBaseFact(DeleteDataFactHint, token)
	fact := DeleteDataFact{
		BaseFact: bf,
		sender:   sender,
		contract: contract,
		dataKey:  key,
		currency: currency,
	}

	fact.SetHash(fact.GenerateHash())
	return fact
}

func (fact DeleteDataFact) IsValid(b []byte) error {
	if len(fact.dataKey) < 1 || len(fact.dataKey) > types.MaxKeyLen {
		return common.ErrFactInvalid.Wrap(
			common.ErrValOOR.Wrap(
				errors.Errorf("invalid data key length %v < 1 or %v > %v", len(fact.dataKey), len(fact.dataKey), types.MaxKeyLen)))
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

func (fact DeleteDataFact) Hash() util.Hash {
	return fact.BaseFact.Hash()
}

func (fact DeleteDataFact) GenerateHash() util.Hash {
	return valuehash.NewSHA256(fact.Bytes())
}

func (fact DeleteDataFact) Bytes() []byte {
	return util.ConcatBytesSlice(
		fact.Token(),
		fact.sender.Bytes(),
		fact.contract.Bytes(),
		[]byte(fact.dataKey),
		fact.currency.Bytes(),
	)
}

func (fact DeleteDataFact) Token() mitumbase.Token {
	return fact.BaseFact.Token()
}

func (fact DeleteDataFact) Sender() mitumbase.Address {
	return fact.sender
}

func (fact DeleteDataFact) Contract() mitumbase.Address {
	return fact.contract
}

func (fact DeleteDataFact) DataKey() string {
	return fact.dataKey
}

func (fact DeleteDataFact) Currency() currencytypes.CurrencyID {
	return fact.currency
}

type DeleteData struct {
	common.BaseOperation
}

func NewDeleteData(fact DeleteDataFact) (DeleteData, error) {
	return DeleteData{BaseOperation: common.NewBaseOperation(DeleteDataHint, fact)}, nil
}
