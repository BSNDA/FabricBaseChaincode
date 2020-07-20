package bsnchaincode

import (
	"bsnBaseCC/models"
	"bsnBaseCC/utils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	strings2 "strings"
)

type BsnChainCode struct {
}

// Set Logger
func SetLogger(logInfo ...interface{}) {
	utils.SetLogger(logInfo)
}

// Data Check
func DataCheck(model string) error {
	if strings2.TrimSpace(model) == "" {
		return errors.New("baseKey field cannot be empty")
	}
	return nil
}

func (t *BsnChainCode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	SetLogger("ChainCode Init start......")
	defer SetLogger("ChainCode Init end......")
	dbBaseModel := models.DBBaseModel{BaseKey: "cc_key_", BaseInfo: "Welcome to use ChainCode "}
	reqJsonValue, err := json.Marshal(&dbBaseModel)
	if err != nil {
		return shim.Error(fmt.Sprintf("Data conversion failed:%s", err.Error()))
	}
	err = stub.PutState(dbBaseModel.BaseKey, reqJsonValue)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func (t *BsnChainCode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	function, args := stub.GetFunctionAndParameters()
	switch function {
	case "set": // set
		return Set(stub, args)
	case "update": // update
		return Update(stub, args)
	case "delete": // delete
		return Delete(stub, args)
	case "get": // get
		return Get(stub, args)
	case "getHistory": // getHistory
		return GetHistory(stub, args)
	default:
		SetLogger("Invalid function")
		break
	}
	return shim.Error("Invalid Request")
}
