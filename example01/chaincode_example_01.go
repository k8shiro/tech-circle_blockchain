package main

import (
        "errors"
        "fmt"
        "strconv"
        
        "github.com/hyperledger/fabric/core/chaincode/shim"
        )

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode1 struct {
}

func (t *SimpleChaincode1) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
    if len(args) != 0 {
        return nil, errors.New("Incorrect number of arguments. Expecting 0")
    }
    
    return nil, nil
}

// Transaction makes payment of X units from A to B
func (t *SimpleChaincode1) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
    var A string    // Entities
    var Aval int // Asset holdings
    var err error
    
    if len(args) != 2 {
        return nil, errors.New("Incorrect number of arguments. Expecting 4")
    }
    
    // Initialize the chaincode
    A = args[0]
    Aval, err = strconv.Atoi(args[1])
    if err != nil {
        return nil, errors.New("Expecting integer value for asset holding")
    }
    
    
    // Write the state to the ledger
    err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
    if err != nil {
        return nil, err
    }
    
    
    return nil, nil
}


// Query callback representing the query of a chaincode
func (t *SimpleChaincode1) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
    if function != "query" {
        return nil, errors.New("Invalid query function name. Expecting \"query\"")
    }
    var A string // Entities
    var err error
    
    if len(args) != 1 {
        return nil, errors.New("Incorrect number of arguments. Expecting name of the person to query")
    }
    
    A = args[0]
    
    // Get the state from the ledger
    Avalbytes, err := stub.GetState(A)
    if err != nil {
        jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
        return nil, errors.New(jsonResp)
    }
    
    if Avalbytes == nil {
        jsonResp := "{\"Error\":\"Nil amount for " + A + "\"}"
        return nil, errors.New(jsonResp)
    }
    
    jsonResp := "{\"Name\":\"" + A + "\",\"Amount\":\"" + string(Avalbytes) + "\"}"
    fmt.Printf("Query Response:%s\n", jsonResp)
    return Avalbytes, nil
}

func main() {
    err := shim.Start(new(SimpleChaincode1))
    if err != nil {
        fmt.Printf("Error starting Simple chaincode: %s", err)
    }
}


