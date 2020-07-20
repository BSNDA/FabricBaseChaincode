package bsnchaincode

import (
	"encoding/json"
	"fmt"
	"github.com/BSNDA/FabricBaseChaincode/chaincode/go/bsnBaseCC/models"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"time"
)

// define global constants
const (
	// define the prefix of key value
	key_prefix = "base_key_"
)

// set the rules for generating key values
func constructKey(baseKey string) string {
	return key_prefix + baseKey
}

// save the data
func Set(stubInterface shim.ChaincodeStubInterface, strings []string) peer.Response {
	SetLogger("save data to start......")
	defer SetLogger("save data to end......")
	if len(strings) != 1 {
		return shim.Error("parameters error")
	}
	//verify data 1、key value not empty
	var dtoModel models.DTOBaseModel
	err := json.Unmarshal([]byte(strings[0]), &dtoModel)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to convert data:%s", err.Error()))
	}
	err = DataCheck(dtoModel.BaseKey)
	if err != nil {
		return shim.Error(fmt.Sprintf("data formatting error:%s", err.Error()))
	}

	mainKey := constructKey(dtoModel.BaseKey)
	SetLogger("key query", mainKey)

	//verify data 2、data not found
	result, err := stubInterface.GetState(mainKey)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to get primary key:%s", err.Error()))
	}
	if len(result) > 0 {
		SetLogger(fmt.Sprintf("the【%s】info already exists", mainKey))
		return shim.Error(fmt.Sprintf("key value already exists"))
	}

	SetLogger(fmt.Sprintf("start to add data【%s】", mainKey))

	//save data to database
	dbBaseModel := models.DTOBase2Db(dtoModel)
	reqJsonValue, err := json.Marshal(&dbBaseModel)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to convert data:%s", err.Error()))
	}

	err = stubInterface.PutState(mainKey, reqJsonValue)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to add data:%s", err.Error()))
	}

	SetLogger("finish saving data", mainKey)

	return shim.Success([]byte("SUCCESS"))
}

// update data
func Update(stubInterface shim.ChaincodeStubInterface, strings []string) peer.Response {
	SetLogger("start to update data......")
	defer SetLogger("finish updating data......")
	if len(strings) != 1 {
		return shim.Error("parameters error")
	}

	//verify data 1、key value not empty
	var dtoModel models.DTOBaseModel
	err := json.Unmarshal([]byte(strings[0]), &dtoModel)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to convert data:%s", err.Error()))
	}
	err = DataCheck(dtoModel.BaseKey)
	if err != nil {
		return shim.Error(fmt.Sprintf("data formatting error:%s", err.Error()))
	}

	mainKey := constructKey(dtoModel.BaseKey)

	//verify data 2、key value exists
	result, err := stubInterface.GetState(mainKey)
	SetLogger(fmt.Sprintf("search for data with the key value of【%s】 ", mainKey))

	if err != nil {
		return shim.Error(fmt.Sprintf("failed to get primary key:%s", err.Error()))
	}
	if len(result) == 0 {
		return shim.Error(fmt.Sprintf("the key value not exists"))
	}

	SetLogger(fmt.Sprintf("succeed to find data with the key value of【%s】", mainKey))

	//change data in database
	dbBaseModel := models.DTOBase2Db(dtoModel)
	reqJsonValue, err := json.Marshal(&dbBaseModel)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to convert data:%s", err.Error()))
	}

	err = stubInterface.PutState(mainKey, reqJsonValue)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to update data:%s", err.Error()))
	}
	SetLogger(fmt.Sprintf("key value: 【%s】data updated", mainKey))

	SetLogger("end of data modification")

	return shim.Success([]byte("SUCCESS"))
}

// delete data
func Delete(stubInterface shim.ChaincodeStubInterface, strings []string) peer.Response {
	SetLogger("start to delet data......")
	defer SetLogger("finish deleting data......")
	if len(strings) != 1 {
		return shim.Error("parameters error")
	}

	//verify data 1、key value not empty
	key := strings[0]
	err := DataCheck(key)
	if err != nil {
		return shim.Error(fmt.Sprintf("data formatting error:%s", err.Error()))
	}

	mainKey := constructKey(key)

	SetLogger(fmt.Sprintf("search for data with the key value of【%s】", mainKey))

	//verify data 2、key value exists
	result, err := stubInterface.GetState(mainKey)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to get primary key:%s", err.Error()))
	}
	if len(result) == 0 {
		return shim.Error(fmt.Sprintf("key value not exist"))
	}

	//delete data in database
	err = stubInterface.DelState(mainKey)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to delete data:%s", err.Error()))
	}
	SetLogger(fmt.Sprintf("key value: 【%s】 data deleted", mainKey))

	SetLogger("finish deleting data")

	return shim.Success([]byte("SUCCESS"))
}

// data query
func Get(stubInterface shim.ChaincodeStubInterface, strings []string) peer.Response {
	SetLogger("start to get data......")
	defer SetLogger("finish getting data......")
	if len(strings) != 1 {
		return shim.Error("parameters error")
	}
	//verify data 1、key value not empty
	key := strings[0]
	err := DataCheck(key)
	if err != nil {
		return shim.Error(fmt.Sprintf("data formatting error:%s", err.Error()))
	}

	mainKey := constructKey(key)
	SetLogger(fmt.Sprintf("search for【%s】data", mainKey))

	//verify data 2、key value exists
	result, err := stubInterface.GetState(mainKey)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to get primary key:%s", err.Error()))
	}
	if len(result) == 0 {
		SetLogger(fmt.Sprintf("data with key value of not found ", mainKey))
		return shim.Error("data not found")
	}

	var dbBaseModel models.DBBaseModel
	err = json.Unmarshal(result, &dbBaseModel)
	if err != nil {
		return shim.Error(fmt.Sprintf("data formatting error:%s", err.Error()))
	}

	//convert the data in the database to return data
	dtoModel := models.Db2DTOBase(dbBaseModel)
	SetLogger(fmt.Sprintf("query of data with the key value of【%s】completed", mainKey))

	SetLogger("finish getting data")
	return shim.Success([]byte(dtoModel.BaseValue))
}

// query of data history
func GetHistory(stubInterface shim.ChaincodeStubInterface, strings []string) peer.Response {
	SetLogger("start to get data history......")
	defer SetLogger("finish getting data history......")
	if len(strings) != 1 {
		return shim.Error("parameters error")
	}
	//verify data 1、key value not empty
	key := strings[0]
	err := DataCheck(key)
	if err != nil {
		return shim.Error(fmt.Sprintf("data formatting error:%s", err.Error()))
	}

	mainKey := constructKey(key)
	SetLogger(fmt.Sprintf("search for history info with the key value of", mainKey))

	resultsIterator, err := stubInterface.GetHistoryForKey(mainKey)
	if err != nil {
		SetLogger(fmt.Sprintf("failed to get history info:%s", err.Error()))
		return shim.Error(fmt.Sprintf("failed to get history info:%s", err.Error()))
	}

	res := []models.DTOHistoryModel{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			SetLogger(fmt.Sprintf("failed to get history info:%s", err.Error()))
			return shim.Error(fmt.Sprintf("failed to get history info:%s", err.Error()))
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

	return shim.Success(resultByte)
}
