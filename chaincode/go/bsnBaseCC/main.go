package main

import (
	"fmt"
	"github.com/BSNDA/FabricBaseChaincode/chaincode/go/bsnBaseCC/bsnchaincode"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func main() {
	err := shim.Start(new(bsnchaincode.BsnChainCode))
	if err != nil {
		fmt.Printf("Error starting BsnChainCode: %s", err)
	}
}
