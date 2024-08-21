package digest

import (
	"github.com/ProtoconNet/mitum-storage/state"
	mitumbase "github.com/ProtoconNet/mitum2/base"
	"go.mongodb.org/mongo-driver/mongo"
)

func (bs *BlockSession) prepareStorage() error {
	if len(bs.sts) < 1 {
		return nil
	}

	var storageModels []mongo.WriteModel
	var storageDataModels []mongo.WriteModel
	for i := range bs.sts {
		st := bs.sts[i]
		switch {
		case state.IsDesignStateKey(st.Key()):
			j, err := bs.handleStorageDesignState(st)
			if err != nil {
				return err
			}
			storageModels = append(storageModels, j...)
		case state.IsDataStateKey(st.Key()):
			j, err := bs.handleStorageDataState(st)
			if err != nil {
				return err
			}
			storageDataModels = append(storageDataModels, j...)
		default:
			continue
		}
	}

	bs.storageModels = storageModels
	bs.storageDataModels = storageDataModels

	return nil
}

func (bs *BlockSession) handleStorageDesignState(st mitumbase.State) ([]mongo.WriteModel, error) {
	if serviceDesignDoc, err := NewDesignDoc(st, bs.st.Encoder()); err != nil {
		return nil, err
	} else {
		return []mongo.WriteModel{
			mongo.NewInsertOneModel().SetDocument(serviceDesignDoc),
		}, nil
	}
}

func (bs *BlockSession) handleStorageDataState(st mitumbase.State) ([]mongo.WriteModel, error) {
	if StorageDataDoc, err := NewStorageDataDoc(st, bs.block.Manifest().ProposedAt(), bs.st.Encoder()); err != nil {
		return nil, err
	} else {
		return []mongo.WriteModel{
			mongo.NewInsertOneModel().SetDocument(StorageDataDoc),
		}, nil
	}
}
