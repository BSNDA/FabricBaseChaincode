package chaincode

import (
	"bsnBaseChaincode/models"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-contract-api-go/metadata"
	"time"
)

func NewChaincode() *BSNBaseChaincode {

	demo := new(BSNBaseChaincode)

	demo.Info = metadata.InfoMetadata{
		Version:     "0.0.1",
		Title:       "base",
		Description: "bsn base chaincode",
	}
	demo.Name = "base"
	return demo
}

type BSNBaseChaincode struct {
	contractapi.Contract
}

func (c *BSNBaseChaincode) Set(ctx contractapi.TransactionContextInterface, data string) (string, error) {
	SetLogger("save data to start......")
	defer SetLogger("save data to end......")

	//verify data 1、key value not empty
	var dtoModel models.DTOBaseModel
	err := json.Unmarshal([]byte(data), &dtoModel)
	if err != nil {
		return "", errors.New(fmt.Sprintf("failed to convert data:%s", err.Error()))
	}
	err = DataCheck(dtoModel.BaseKey)
	if err != nil {
		return "", errors.New(fmt.Sprintf("data formatting error:%s", err.Error()))
	}

	mainKey := constructKey(dtoModel.BaseKey)
	SetLogger("key query", mainKey)

	result, err := ctx.GetStub().GetState(mainKey)
	if err != nil {
		return "", errors.New(fmt.Sprintf("failed to get primary key:%s", err.Error()))
	}
	if len(result) > 0 {
		SetLogger(fmt.Sprintf("the【%s】info already exists", mainKey))
		return "", errors.New(fmt.Sprintf("key value already exists"))
	}

	SetLogger(fmt.Sprintf("start to add data【%s】", mainKey))

	//save data to database
	dbBaseModel := models.DTOBase2Db(dtoModel)
	reqJsonValue, err := json.Marshal(&dbBaseModel)
	if err != nil {
		return "", errors.New(fmt.Sprintf("failed to convert data:%s", err.Error()))
	}

	err = ctx.GetStub().PutState(mainKey, reqJsonValue)
	if err != nil {
		return "", errors.New(fmt.Sprintf("failed to add data:%s", err.Error()))
	}

	SetLogger("finish saving data", mainKey)

	ctx.GetStub().SetEvent("event_set", []byte(data))

	return "SUCCESS", nil

}

func (c *BSNBaseChaincode) Update(ctx contractapi.TransactionContextInterface, data string) (string, error) {
	SetLogger("start to update data......")
	defer SetLogger("finish updating data......")

	//verify data 1、key value not empty
	var dtoModel models.DTOBaseModel
	err := json.Unmarshal([]byte(data), &dtoModel)
	if err != nil {
		return "", errors.New(fmt.Sprintf("failed to convert data:%s", err.Error()))
	}
	err = DataCheck(dtoModel.BaseKey)
	if err != nil {
		return "", errors.New(fmt.Sprintf("data formatting error:%s", err.Error()))
	}

	mainKey := constructKey(dtoModel.BaseKey)

	//verify data 2、key value exists
	result, err := ctx.GetStub().GetState(mainKey)
	SetLogger(fmt.Sprintf("search for data with the key value of【%s】 ", mainKey))

	if err != nil {
		return "", errors.New(fmt.Sprintf("failed to get primary key:%s", err.Error()))
	}
	if len(result) == 0 {
		return "", errors.New(fmt.Sprintf("the key value not exists"))
	}

	SetLogger(fmt.Sprintf("succeed to find data with the key value of【%s】", mainKey))

	//change data in database
	dbBaseModel := models.DTOBase2Db(dtoModel)
	reqJsonValue, err := json.Marshal(&dbBaseModel)
	if err != nil {
		return "", errors.New(fmt.Sprintf("failed to convert data:%s", err.Error()))
	}

	err = ctx.GetStub().PutState(mainKey, reqJsonValue)
	if err != nil {
		return "", errors.New(fmt.Sprintf("failed to update data:%s", err.Error()))
	}
	SetLogger(fmt.Sprintf("key value: 【%s】data updated", mainKey))

	SetLogger("end of data modification")

	ctx.GetStub().SetEvent("event_update", []byte(data))

	return "SUCCESS", nil
}

func (c *BSNBaseChaincode) Delete(ctx contractapi.TransactionContextInterface, key string) (string, error) {
	SetLogger("start to delet data......")
	defer SetLogger("finish deleting data......")

	//verify data 1、key value not empty

	err := DataCheck(key)
	if err != nil {
		return "", errors.New(fmt.Sprintf("data formatting error:%s", err.Error()))
	}

	mainKey := constructKey(key)

	SetLogger(fmt.Sprintf("search for data with the key value of【%s】", mainKey))

	//verify data 2、key value exists
	result, err := ctx.GetStub().GetState(mainKey)
	if err != nil {
		return "", errors.New(fmt.Sprintf("failed to get primary key:%s", err.Error()))
	}
	if len(result) == 0 {
		return "", errors.New(fmt.Sprintf("key value not exist"))
	}

	//delete data in database
	err = ctx.GetStub().DelState(mainKey)
	if err != nil {
		return "", errors.New(fmt.Sprintf("failed to delete data:%s", err.Error()))
	}
	SetLogger(fmt.Sprintf("key value: 【%s】 data deleted", mainKey))

	SetLogger("finish deleting data")

	ctx.GetStub().SetEvent("event_delete", []byte(key))

	return "SUCCESS", nil
}

// data query
func (c *BSNBaseChaincode) Get(ctx contractapi.TransactionContextInterface, key string) (string, error) {
	SetLogger("start to get data......")
	defer SetLogger("finish getting data......")

	//verify data 1、key value not empty
	err := DataCheck(key)
	if err != nil {
		return "", errors.New(fmt.Sprintf("data formatting error:%s", err.Error()))
	}

	mainKey := constructKey(key)
	SetLogger(fmt.Sprintf("search for【%s】data", mainKey))

	//verify data 2、key value exists
	result, err := ctx.GetStub().GetState(mainKey)
	if err != nil {
		return "", errors.New(fmt.Sprintf("failed to get primary key:%s", err.Error()))
	}
	if len(result) == 0 {
		SetLogger(fmt.Sprintf("data with key value of not found ", mainKey))
		return "", errors.New("data not found")
	}

	var dbBaseModel models.DBBaseModel
	err = json.Unmarshal(result, &dbBaseModel)
	if err != nil {
		return "", errors.New(fmt.Sprintf("data formatting error:%s", err.Error()))
	}

	//convert the data in the database to return data
	dtoModel := models.Db2DTOBase(dbBaseModel)
	SetLogger(fmt.Sprintf("query of data with the key value of【%s】completed", mainKey))

	SetLogger("finish getting data")
	return dtoModel.BaseValue, nil
}

// query of data history
func (c *BSNBaseChaincode) GetHistory(ctx contractapi.TransactionContextInterface, key string) (string, error) {
	SetLogger("start to get data history......")
	defer SetLogger("finish getting data history......")

	//verify key value not empty
	err := DataCheck(key)
	if err != nil {
		return "", errors.New(fmt.Sprintf("data formatting error:%s", err.Error()))
	}

	mainKey := constructKey(key)
	SetLogger(fmt.Sprintf("search for history info with the key value of", mainKey))

	resultsIterator, err := ctx.GetStub().GetHistoryForKey(mainKey)
	if err != nil {
		SetLogger(fmt.Sprintf("failed to get history info:%s", err.Error()))
		return "", errors.New(fmt.Sprintf("failed to get history info:%s", err.Error()))
	}

	res := []models.DTOHistoryModel{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			SetLogger(fmt.Sprintf("failed to get history info:%s", err.Error()))
			return "", errors.New(fmt.Sprintf("failed to get history info:%s", err.Error()))
		}
		txTimestamp := queryResponse.GetTimestamp()
		txTime := ""

		if txTimestamp != nil {
			txTime = time.Unix(txTimestamp.Seconds, 0).Format("2006-01-02 15:04:05")
		}

		temp := models.DTOHistoryModel{
			TxId:      queryResponse.TxId,
			IsDelete:  queryResponse.IsDelete,
			Value:     string(queryResponse.Value),
			Timestamp: txTime,
		}

		res = append(res, temp)

	}
	resultByte, _ := json.Marshal(res)

	SetLogger("query of history data ended......")

	return string(resultByte), nil
}

// GetEvaluateTransactions returns a list of function names that should be tagged in the
// metadata as "evaluate" to indicate to a user of the chaincode that they should query
// rather than invoke these functions
func (d *BSNBaseChaincode) GetEvaluateTransactions() []string {
	return []string{"Get", "GetHistory"}
}
