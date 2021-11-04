package models

// database model
type DBBaseModel struct {
	BaseKey  string
	BaseInfo string
}

// DTO(Data Transfer Object) Model
type DTOBaseModel struct {
	BaseKey   string `json:"baseKey"`
	BaseValue string `json:"baseValue"`
}

// History Info
type DTOHistoryModel struct {
	TxId      string `json:"txId"`
	Value     string `json:"dataInfo"`
	Timestamp string `json:"txTime"`
	IsDelete  bool   `json:"isDelete"`
}

func DTOBase2Db(dtoModel DTOBaseModel) DBBaseModel {
	return DBBaseModel{BaseKey: dtoModel.BaseKey, BaseInfo: dtoModel.BaseValue}
}

func Db2DTOBase(model DBBaseModel) DTOBaseModel {
	return DTOBaseModel{BaseKey: model.BaseKey, BaseValue: model.BaseInfo}
}
