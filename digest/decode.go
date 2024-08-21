package digest

import (
	bsonenc "github.com/ProtoconNet/mitum-currency/v3/digest/util/bson"
	"github.com/ProtoconNet/mitum-storage/types"
	"github.com/ProtoconNet/mitum2/util"
	"github.com/ProtoconNet/mitum2/util/encoder"
	"github.com/ProtoconNet/mitum2/util/hint"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

func LoadStorageDataFromDoc(b []byte, encs *encoder.Encoders) (bson.Raw /* id */, interface{} /* data */, int64 /* height */, string /* operations */, string /* timestamp */, bool /* deleted */, error) {

	var bd StorageDataDocBSONUnMarshaler
	if err := bsonenc.Unmarshal(b, &bd); err != nil {
		return nil, nil, 0, "", "", false, err
	}

	ht, err := hint.ParseHint(bd.E)
	if err != nil {
		return nil, nil, 0, "", "", false, err
	}

	enc, found := encs.Find(ht)
	if !found {
		return nil, nil, 0, "", "", false, util.ErrNotFound.Errorf("Encoder not found for %q", ht)
	}

	if !bd.H {
		return nil, nil, 0, "", "", false, nil
	}

	doc, ok := bd.D.DocumentOK()
	if !ok {
		return nil, nil, 0, "", "", false, errors.Errorf("Hinted should be mongodb Document")
	}

	var data interface{}
	if i, err := enc.Decode([]byte(doc)); err != nil {
		return nil, nil, 0, "", "", false, err
	} else {
		data = i
	}

	return bd.I, data, bd.HT, bd.O, bd.T, bd.DL, nil
}

func LoadStorageData(decoder func(interface{}) error, encs *encoder.Encoders) (*types.Data, int64, string, string, bool, error) {
	var b bson.Raw

	if err := decoder(&b); err != nil {
		return nil, 0, "", "", false, err
	}

	if _, hinter, height, operation, timestamp, deleted, err := LoadStorageDataFromDoc(b, encs); err != nil {
		return nil, 0, "", "", false, err
	} else if m, ok := hinter.(types.Data); !ok {
		return nil, 0, "", "", false, errors.Errorf("Not types.Data: %T", hinter)
	} else {
		return &m, height, operation, timestamp, deleted, nil
	}
}
