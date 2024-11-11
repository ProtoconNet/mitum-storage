package storage

import (
	"github.com/ProtoconNet/mitum-currency/v3/common"
	"github.com/ProtoconNet/mitum2/base"
	"github.com/ProtoconNet/mitum2/util/encoder"
	"github.com/pkg/errors"
)

func (fact *UpdateDatasFact) unpack(enc encoder.Encoder, sa string, bit []byte) error {
	switch a, err := base.DecodeAddress(sa, enc); {
	case err != nil:
		return err
	default:
		fact.sender = a
	}

	hit, err := enc.DecodeSlice(bit)
	if err != nil {
		return err
	}

	items := make([]UpdateDatasItem, len(hit))
	for i := range hit {
		j, ok := hit[i].(UpdateDatasItem)
		if !ok {
			return common.ErrTypeMismatch.Wrap(errors.Errorf("expected %T, not %T", UpdateDatasItem{}, hit[i]))
		}

		items[i] = j
	}
	fact.items = items

	return nil
}
