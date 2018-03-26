package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type Dapp struct {
}

func (d *Dapp) Init(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("Init Function")
	var a string
	var value int
	var er error

	args := stub.GetStringArgs()
	if len(args) != 2 {
		return shim.Error("Need at least two arguments")
	}
	a = args[0]
	value, er = strconv.Atoi(args[1])
	if er != nil {
		return shim.Error(fmt.Sprintf("error converting ascii to int,%s", er))
	}
	err := stub.PutState(args[0], []byte(args[1]))
	if err != nil {
		return shim.Error("Failed to put state into the ledger")
	}

	fmt.Printf("Key: %s, Value: %d, committed to ledger\n", a, value)

	return shim.Success(nil)
}

func (d *Dapp) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fn, args := stub.GetFunctionAndParameters()
	fmt.Println("Invoke Function", fn)

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
	return shim.Success([]byte(result))
}

func set(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 2 {
		return "", fmt.Errorf("expect a key and a value")
	}

	var a string
	var value int

	a = args[0]
	value, _ = strconv.Atoi(args[1])

	err := stub.PutState(args[0], []byte(args[1]))
	if err != nil {
		return "", fmt.Errorf("error putting it into the ledger")
	}

	fmt.Printf("Key: %s, Value: %d, committed to ledger\n", a, value)
	return args[1], nil
}

func get(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("expects a key")
	}

	var a string
	a = args[0]

	bs, err := stub.GetState(args[0])
	if err != nil {
		return "", fmt.Errorf("error getting the value")
	}

	if bs == nil {
		return "", fmt.Errorf("asset not found: %s", args[0])
	}

	fmt.Printf("requested value, for the key: %s is %s\n", a, string(bs))

	return string(bs), nil
}

func main() {
	if err := shim.Start(new(Dapp)); err != nil {
		fmt.Println("Error starting Dapp.")
		os.Exit(1)
	}
}
