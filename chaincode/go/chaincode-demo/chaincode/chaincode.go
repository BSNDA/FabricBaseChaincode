package chaincode

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-contract-api-go/metadata"
	"time"
)

func NewChaincode() *DemoChaincode {

	demo := new(DemoChaincode)

	demo.Info = metadata.InfoMetadata{
		Version: "0.0.1",
		Title:   "demo",
	}

	demo.Name = "demo"
	demo.UnknownTransaction = demo.unknown
	demo.BeforeTransaction = demo.before
	demo.AfterTransaction = demo.after

	return demo
}

type DemoChaincode struct {
	contractapi.Contract
}

/// Set store key value on the chain , and return 'success' if successful .
func (d *DemoChaincode) Set(ctx contractapi.TransactionContextInterface, key string, value string) (string, error) {

	err := ctx.GetStub().PutState(key, []byte(value))
	if err != nil {
		return "", err
	} else {
		return "success", nil
	}
}

/// Query query key on the chain , and return value if successful .
func (d *DemoChaincode) Query(ctx contractapi.TransactionContextInterface, key string) (string, error) {
	valueBytes, err := ctx.GetStub().GetState(key)
	if err != nil {
		return "", fmt.Errorf("Failed to read from world state. %s", err.Error())
	}
	return string(valueBytes), nil
}

/// Delete remove key on the chain , and return 'success' if successful .
func (d *DemoChaincode) Delete(ctx contractapi.TransactionContextInterface, key string) (string, error) {
	err := ctx.GetStub().DelState(key)
	if err != nil {
		return "", err
	} else {
		return "success", nil
	}
}

/// History query key operation record on the chain .
func (d *DemoChaincode) History(ctx contractapi.TransactionContextInterface, key string) (string, error) {

	history, err := ctx.GetStub().GetHistoryForKey(key)
	if err != nil {
		return "", fmt.Errorf("Failed to read history. %s", err.Error())
	}

	type KeyRecord struct {
		Key      string `json:"key"`
		Value    string `json:"value"`
		TxId     string `json:"txId"`
		IsDelete bool   `json:"isDelete"`
		Time     string `json:"time"`
	}

	var records []KeyRecord

	for {
		if history.HasNext() {
			res, err := history.Next()
			if err != nil {
				continue
			}
			kr := KeyRecord{
				Key:      key,
				Value:    base64.StdEncoding.EncodeToString(res.Value),
				TxId:     res.TxId,
				IsDelete: res.IsDelete,
				Time:     time.Unix(res.Timestamp.Seconds, int64(res.Timestamp.Nanos)).Format("2006-01-02 15:04:05.000 -0700 MST"),
			}
			//tm := time.Unix(res.Timestamp.Seconds,int64(res.Timestamp.Nanos))
			//kr.Time = tm.Format("2006-01-02 15:04:05.000 -0700 MST")
			records = append(records, kr)

		} else {
			break
		}
	}

	jsonBytes, _ := json.Marshal(&records)
	return string(jsonBytes), nil
}

// GetEvaluateTransactions returns a list of function names that should be tagged in the
// metadata as "evaluate" to indicate to a user of the chaincode that they should query
// rather than invoke these functions
func (d *DemoChaincode) GetEvaluateTransactions() []string {
	return []string{"Query", "History"}
}

func (d *DemoChaincode) unknown(ctx contractapi.TransactionContextInterface) error {
	fcn, _ := ctx.GetStub().GetFunctionAndParameters()
	return fmt.Errorf("%s is not found", fcn)
}

func (d *DemoChaincode) before(ctx contractapi.TransactionContextInterface) {
	fmt.Println("DemoChaincode.before")
}

func (d *DemoChaincode) after(ctx contractapi.TransactionContextInterface, data interface{}) {
	fmt.Println(fmt.Sprintf("After function called with %v", data))

}
