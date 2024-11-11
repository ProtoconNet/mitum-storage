package cmds

type StorageCommand struct {
	CreateData    CreateDataCommand    `cmd:"" name:"create-data" help:"create new storage data"`
	CreateDatas   CreateDatasCommand   `cmd:"" name:"create-datas" help:"create new storage datas"`
	UpdateData    UpdateDataCommand    `cmd:"" name:"update-data" help:"update storage data"`
	UpdateDatas   UpdateDatasCommand   `cmd:"" name:"update-datas" help:"update storage datas"`
	DeleteData    DeleteDataCommand    `cmd:"" name:"delete-data" help:"delete storage data"`
	RegisterModel RegisterModelCommand `cmd:"" name:"register-model" help:"register storage model"`
}
