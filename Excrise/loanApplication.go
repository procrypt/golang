package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"fmt"
	"strconv"
)

//type PersonalInfo struct {
//	FirstName string `json:"first_name"`
//	LastName  string `json:"last_name"`
//}

type LoanApplication struct {

}

func (l *LoanApplication) Init(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("Init Function")

	var id int
	var firstName string
	//var lastName string

	args := stub.GetStringArgs()
	if len(args) != 2 {
		return shim.Error("Expected three arguments Loan ID, First Name and Last Name")
	}

	id,_ = strconv.Atoi(args[0])
	firstName = args[1]
	//lastName  = args[2]

	err := stub.PutState(args[0], []byte(args[1]))
	if err != nil {
		return shim.Error("Error committing data to ledger")
	}
	fmt.Printf("Data committed to ledger, LoanID: %d, First Name: %s, Last Name: <Not Implemeted>\n", id, firstName)
	return shim.Success(nil)

}

func (l *LoanApplication) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fn, args := stub.GetFunctionAndParameters()

	var result string
	var err error

	if fn == "GetLoanApplication" {
		result, err = GetLoanApplication(stub, args)
	} else {
		result, err = CreateLoanApplication(stub, args)
	}
	if err != nil {
		return shim.Error("Error")
	}

	return shim.Success([]byte(result))
}

func CreateLoanApplication(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 2 {
		return "", fmt.Errorf("expected three arguments Loan ID, First Name and Last Name")
	}

	var id int
	var firstName string
	//var lastName

	id,_ = strconv.Atoi(args[0])
	firstName = args[1]
	//lastName = args[2]

	err := stub.PutState(args[0], []byte(args[1]))
	if err != nil {
		return "", fmt.Errorf("error committing data to ledger")
	}

	fmt.Printf("Data committed to ledger, LoanID: %d, First Name: %s, Last Name: <Not Implemeted>\n", id, firstName)
	return "", nil
	}

func GetLoanApplication(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("expects LoanID")
	}

	var loanID string
	loanID = args[0]

	bs, err := stub.GetState(loanID)
	if err != nil {
		return "", fmt.Errorf("error getting the loan application")
	}
	if bs == nil {
		return "", fmt.Errorf("loan application not found: %s", loanID)
	}

	fmt.Printf("Loan Application details for %s is %s\n", loanID, string(bs))
	return string(bs), nil

}

func main() {
	if err := shim.Start(new(LoanApplication)); err != nil {
		fmt.Println("Error starting chaincode")
	}
}
