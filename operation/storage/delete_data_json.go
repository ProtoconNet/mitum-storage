package storage

import (
	"github.com/ProtoconNet/mitum-currency/v3/common"
	"github.com/ProtoconNet/mitum-currency/v3/types"
	mitumbase "github.com/ProtoconNet/mitum2/base"
	"github.com/ProtoconNet/mitum2/util"
	"github.com/ProtoconNet/mitum2/util/encoder"
)

type DeleteDataFactJSONMarshaler struct {
	mitumbase.BaseFactJSONMarshaler
	Sender   mitumbase.Address `json:"sender"`
	Contract mitumbase.Address `json:"contract"`
	DataKey  string            `json:"dataKey"`
	Currency types.CurrencyID  `json:"currency"`
}

func (fact DeleteDataFact) MarshalJSON() ([]byte, error) {
	return util.MarshalJSON(DeleteDataFactJSONMarshaler{
		BaseFactJSONMarshaler: fact.BaseFact.JSONMarshaler(),
		Sender:                fact.sender,
		Contract:              fact.contract,
		DataKey:               fact.dataKey,
		Currency:              fact.currency,
	})
}

type DeleteDataFactJSONUnmarshaler struct {
	mitumbase.BaseFactJSONUnmarshaler
	Sender   string `json:"sender"`
	Contract string `json:"contract"`
	DataKey  string `json:"dataKey"`
	Currency string `json:"currency"`
}

func (fact *DeleteDataFact) DecodeJSON(b []byte, enc encoder.Encoder) error {
	var u DeleteDataFactJSONUnmarshaler
	if err := enc.Unmarshal(b, &u); err != nil {
		return common.DecorateError(err, common.ErrDecodeJson, *fact)
	}

	fact.BaseFact.SetJSONUnmarshaler(u.BaseFactJSONUnmarshaler)

	if err := fact.unpack(enc, u.Sender, u.Contract, u.DataKey, u.Currency); err != nil {
		return common.DecorateError(err, common.ErrDecodeJson, *fact)
	}

	return nil
}

func (op DeleteData) MarshalJSON() ([]byte, error) {
	return util.MarshalJSON(
		op.BaseOperation.JSONMarshaler(),
	)
}

func (op *DeleteData) DecodeJSON(b []byte, enc encoder.Encoder) error {
	var ubo common.BaseOperation
	if err := ubo.DecodeJSON(b, enc); err != nil {
		return common.DecorateError(err, common.ErrDecodeJson, *op)
	}

	op.BaseOperation = ubo

	return nil
}
