package main

import (
        "errors"
        "fmt"
        "strconv"
        
        "github.com/hyperledger/fabric/core/chaincode/shim"
        )

// SimpleChaincode example simple Chaincode implementation
type ChaincodeExample01 struct {
}

func (t *ChaincodeExample01) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
    if len(args) != 0 {
        return nil, errors.New("Incorrect number of arguments. Expecting 0")
    }
    
    return nil, nil
}

// Transaction makes payment of X units from A to B
func (t *ChaincodeExample01) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
    var itemId, itemName string    // Entities
    //var itemName string // Asset holdings
    var err error
    
    if len(args) != 2 {
        return nil, errors.New("Incorrect number of arguments. Expecting 4")
    }
    
    // Initialize the chaincode
    itemId = args[0]
    itemName = args[1]
    if err != nil {
        return nil, errors.New("Expecting integer value for asset holding")
    }
    
    
    // Write the state to the ledger
    err = stub.PutState(itemId, []byte(itemName))
    if err != nil {
        return nil, err
    }
    
    
    return nil, nil
}


// Query callback representing the query of a chaincode
func (t *ChaincodeExample01) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
    if function != "query" {
        return nil, errors.New("Invalid query function name. Expecting \"query\"")
    }
    var itemId string // Entities
    var err error
    
    if len(args) != 1 {
        return nil, errors.New("Incorrect number of arguments. Expecting name of the person to query")
    }
    
    itemId = args[0]
    
    // Get the state from the ledger
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
    err := shim.Start(new(ChaincodeExample01))
    if err != nil {
        fmt.Printf("Error starting Simple chaincode: %s", err)
    }
}


