package chaincode

import (
	"bsnBaseChaincode/utils"
	"errors"
	"strings"
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

func SetLogger(logInfo ...interface{}) {
	utils.SetLogger(logInfo)
}

// Data Check
func DataCheck(model string) error {
	if strings.TrimSpace(model) == "" {
		return errors.New("baseKey field cannot be empty")
	}
	return nil
}
