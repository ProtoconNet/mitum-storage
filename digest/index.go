package digest

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var indexPrefix = "mitum_digest_"

var accountIndexModels = []mongo.IndexModel{
	{
		Keys: bson.D{bson.E{Key: "address", Value: 1}, bson.E{Key: "height", Value: -1}},
		Options: options.Index().
			SetName("mitum_digest_account"),
	},
	{
		Keys: bson.D{bson.E{Key: "height", Value: -1}},
		Options: options.Index().
			SetName("mitum_digest_account_height"),
	},
	{
		Keys: bson.D{bson.E{Key: "pubs", Value: 1}, bson.E{Key: "height", Value: 1}, bson.E{Key: "address", Value: 1}},
		Options: options.Index().
			SetName("mitum_digest_account_publiskeys"),
	},
}

var balanceIndexModels = []mongo.IndexModel{
	{
		Keys: bson.D{bson.E{Key: "address", Value: 1}, bson.E{Key: "height", Value: -1}},
		Options: options.Index().
			SetName("mitum_digest_balance"),
	},
	{
		Keys: bson.D{
			bson.E{Key: "address", Value: 1},
			bson.E{Key: "currency", Value: 1},
			bson.E{Key: "height", Value: -1},
		},
		Options: options.Index().
			SetName("mitum_digest_balance_currency"),
	},
	{
		Keys: bson.D{bson.E{Key: "height", Value: -1}},
		Options: options.Index().
			SetName("mitum_digest_balance_height"),
	},
}

var operationIndexModels = []mongo.IndexModel{
	{
		Keys: bson.D{bson.E{Key: "addresses", Value: 1}, bson.E{Key: "height", Value: 1}, bson.E{Key: "index", Value: 1}},
		Options: options.Index().
			SetName("mitum_digest_account_operation"),
	},
	{
		Keys: bson.D{bson.E{Key: "height", Value: 1}, bson.E{Key: "index", Value: 1}},
		Options: options.Index().
			SetName("mitum_digest_operation"),
	},
	{
		Keys: bson.D{bson.E{Key: "height", Value: -1}},
		Options: options.Index().
			SetName("mitum_digest_operation_height"),
	},
}

var storageDataServiceIndexModels = []mongo.IndexModel{
	{
		Keys: bson.D{
			bson.E{Key: "contract", Value: 1},
			bson.E{Key: "height", Value: -1}},
		Options: options.Index().
			SetName(indexPrefix + "storage_data_service_contract_height"),
	},
}

var storageDataIndexModels = []mongo.IndexModel{
	{
		Keys: bson.D{
			bson.E{Key: "contract", Value: 1},
			bson.E{Key: "data_key", Value: 1},
			bson.E{Key: "height", Value: -1}},
		Options: options.Index().
			SetName(indexPrefix + "storage_data_contract_datakey_height"),
	},
}

var defaultIndexes = map[string] /* collection */ []mongo.IndexModel{
	defaultColNameAccount:     accountIndexModels,
	defaultColNameBalance:     balanceIndexModels,
	defaultColNameOperation:   operationIndexModels,
	defaultColNameStorage:     storageDataServiceIndexModels,
	defaultColNameStorageData: storageDataIndexModels,
}
