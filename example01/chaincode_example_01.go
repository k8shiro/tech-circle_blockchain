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
    /*var itemId string
    var item string
    var err error
    
    if len(args) != 2 {
        return nil, errors.New("Incorrect number of arguments. Expecting 2")
    }
    
    itemId = args[0]
    item = args[1]
    
    err = stub.PutState(itemId, []byte(item))
    if err != nil {
        return nil, err
    }
    */
    return nil, nil
}

// Transaction makes payment of X units from A to B
func (t *SimpleChaincode) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
    /*var itemId string
    var item string
    var err error
    
    if len(args) != 2 {
        return nil, errors.New("Incorrect number of arguments. Expecting 2")
    }
    
    itemId = args[0]
    item = args[1]
    
    err = stub.PutState(itemId, []byte(item))
    if err != nil {
        return nil, err
    }
    */
    return nil, nil
}


// Query callback representing the query of a chaincode
func (t *SimpleChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
    /*if function != "query" {
        return nil, errors.New("Invalid query function name. Expecting \"query\"")
    }
    var itemId string
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
    */
    return []byte("Hi"), nil
}

func main() {
    err := shim.Start(new(SimpleChaincode))
    if err != nil {
        fmt.Printf("Error starting Simple chaincode: %s", err)
    }
}