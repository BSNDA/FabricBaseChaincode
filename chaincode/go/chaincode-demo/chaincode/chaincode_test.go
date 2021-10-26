package chaincode

import (
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shimtest"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"testing"
)

func TestNewChaincode(t *testing.T) {

	chaincode, err := contractapi.NewChaincode(NewChaincode())

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(chaincode.Info)
	fmt.Println(chaincode.DefaultContract)
	fmt.Println(chaincode.TransactionSerializer)

}
func TestDemoChaincode(t *testing.T) {

	chaincode, err := contractapi.NewChaincode(new(DemoChaincode))

	if err != nil {
		t.Fatal(err)
	}

	uid := "1234567890"

	mockStub := shimtest.NewMockStub("smartContractTest", chaincode)

	var args [][]byte
	args = append(args, []byte("Set"))
	args = append(args, []byte("abc"))
	args = append(args, []byte("123"))

	response := mockStub.MockInvoke(uid, args)
	fmt.Println(string(response.Payload))

	args[0] = []byte("Query")
	response = mockStub.MockInvoke(uid, args)
	fmt.Println(string(response.Payload))

	args[0] = []byte("History")
	response = mockStub.MockInvoke(uid, args)
	fmt.Println(response.Status)
}
