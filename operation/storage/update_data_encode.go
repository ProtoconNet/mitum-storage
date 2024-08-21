package storage

import (
	"github.com/ProtoconNet/mitum-currency/v3/types"
	mitumbase "github.com/ProtoconNet/mitum2/base"
	"github.com/ProtoconNet/mitum2/util/encoder"
)

func (fact *UpdateDataFact) unpack(
	enc encoder.Encoder,
	sa, ta string,
	dataKey, dataValue, cid string,
) error {
	switch sender, err := mitumbase.DecodeAddress(sa, enc); {
	case err != nil:
		return err
	default:
		fact.sender = sender
	}

	switch contract, err := mitumbase.DecodeAddress(ta, enc); {
	case err != nil:
		return err
	default:
		fact.contract = contract
	}

	fact.dataKey = dataKey
	fact.dataValue = dataValue
	fact.currency = types.CurrencyID(cid)

	return nil
}
