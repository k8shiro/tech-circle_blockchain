package main

import (
        "errors"
        "fmt"
        "strconv"
        
        "github.com/hyperledger/fabric/core/chaincode/shim"
        )

type SimpleChaincode struct {
}

func (t *SimpleChaincode) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
    var A, B string    // Entities
    var Aval, Bval int // Asset holdings
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

func (t *SimpleChaincode) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
    var A, B string    // Entities
    var Aval, Bval int // Asset holdings
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


func (t *SimpleChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
    if function != "query" {
        return nil, errors.New("Invalid query function name. Expecting \"query\"")
    }
    var itemId string
    var err error
    
    if len(args) != 1 {
        return nil, errors.New("Incorrect number of arguments. Expecting name of the person to query")
    }
    
    itemId = args[0]
    
    itemBytes, err := stub.GetState(itemId)
    if err != nil {
        jsonResp := "{\"Error\":\"Failed to get state for " + itemId + "\"}"
        return nil, errors.New(jsonResp)
    }
    
    if itemBytes == nil {
        jsonResp := "{\"Error\":\"Nil amount for " + itemId + "\"}"
        return nil, errors.New(jsonResp)
    }
    
    jsonResp := "{\"Name\":\"" + itemId + "\",\"Amount\":\"" + string(itemBytes) + "\"}"
    fmt.Printf("Query Response:%s\n", jsonResp)
    return itemBytes, nil
}

func main() {
    err := shim.Start(new(SimpleChaincode))
    if err != nil {
        fmt.Printf("Error starting Simple chaincode: %s", err)
    }
}