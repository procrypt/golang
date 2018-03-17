package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type SimpleAsset struct {

}

// Init is called during the Chaincode instantiation to initialize and
// data. Chaincode "update" also calls this function to reset or to migrate data.
func(t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) peer.Response {
	// Get args from transaction proposal
	args := stub.GetStringArgs()
	if len(args) != 2 {
		return shim.Error("Incorrect arguments. Expecting a key and a value.")
	}

	// Set up any variables or assets here by calling stub.PutState()
	// We store the key and the value on the ledger
	err := stub.PutState(args[0], []byte(args[1]))
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to create asset: %s", args[0]))
	}
	return shim.Success(nil)
}

// Invoke is called per transaction on the Chaincode.
// Each transaction is either a "get" or a "set" on the asset created by Init method.
// The "set" method may create new assets by specifying a new key-value pair.
func(t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	// Extract the function and args from transaction proposal
	fn, args := stub.GetFunctionAndParameters()

	var result string
	var err error

	if fn == "set" {
		result, err = set(stub, args)
	} else {
		result, err = get(stub, args)
	}
	if err != nil {
		return shim.Error(err.Error())
	}
	// Return the result as success payload
	return shim.Success([]byte(result))
}

// Set stores the asset (both key and a value) on the ledger.
// If the key exist it will overwrite the value with a new one.
func set(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 2 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a key and a value.")
	}

	err := stub.PutState(args[0], []byte(args[1]))
	if err != nil {
		return "", fmt.Errorf("Failed to ser asset: %s", args[0])
	}
	return args[1], nil
}

// Get returns the value of the specified asset key.
func get(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("ncorrect arguments. Expecting a key")
	}
	value, err := stub.GetState(args[0])
	if err != nil {
		return "", fmt.Errorf("failed to get the asset: %s with error: %s", args[0], err)
	}
	if value == nil {
		return "", fmt.Errorf("asset not found: %s", args[0])
	}
	return string(value), nil
}

func main() {
	err := shim.Start(new(SimpleAsset))
	if err != nil {
		fmt.Printf("Error starting SimpleAsset chaincode %s", err)
	}
}
