package storage

import (
	currencytypes "github.com/ProtoconNet/mitum-currency/v3/types"
	"github.com/ProtoconNet/mitum2/base"
	"github.com/ProtoconNet/mitum2/util/encoder"
	"github.com/ProtoconNet/mitum2/util/hint"
)

func (it *UpdateDatasItem) unpack(enc encoder.Encoder, ht hint.Hint,
	cAdr, dataKey, dataValue, cid string,
) error {
	it.BaseHinter = hint.NewBaseHinter(ht)

	switch a, err := base.DecodeAddress(cAdr, enc); {
	case err != nil:
		return err
	default:
		it.contract = a
	}

	it.dataKey = dataKey
	it.dataValue = dataValue
	it.currency = currencytypes.CurrencyID(cid)

	return nil
}
