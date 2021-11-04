package main

import (
	"bsnBaseChaincode/chaincode"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {

	cc, err := contractapi.NewChaincode(chaincode.NewChaincode())

	if err != nil {
		fmt.Println("new chaincode has err " + err.Error())
		return
	}

	if err := cc.Start(); err != nil {
		fmt.Println("start chaincode has err " + err.Error())
	}

}
