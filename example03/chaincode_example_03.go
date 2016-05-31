package main

import (
        "errors"
        "fmt"
        "strconv"
        "encoding/json"
        "github.com/hyperledger/fabric/core/chaincode/shim"
        )

type SimpleChaincode struct {
}

type Baggage struct {
    item string
    position string
    temperature int
    onDelivery bool
}


func (t *SimpleChaincode) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
    if len(args) != 0 {
        return nil, errors.New("Incorrect number of arguments. Expecting 0")
    }
    
    return nil, nil
}

func (t *SimpleChaincode) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
    var key string
    var value Baggage
    var err error
    
    if len(args) != 5 {
        return nil, errors.New("Incorrect number of arguments. Expecting 5")
    }
    
    key = args[0]
    value.item = args[1]
    value.position = args[2]
    value.temperature = strconv.Atoi(args[3])
    value.onDelivery = strconv.ParseBool(args[4])
    valbytes, err := json.Marshal(value)
    if err != nil {
        return nil, err
    }
    
    
    err = stub.PutState(key, valbyte)
    if err != nil {
        return nil, err
    }
    
    
    return nil, nil
}


func (t *SimpleChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
    if function != "query" {
        return nil, errors.New("Invalid query function name. Expecting \"query\"")
    }
    var key string
    var value Baggage
    var err error
    
    if len(args) != 1 {
        return nil, errors.New("Incorrect number of arguments. Expecting name of the person to query")
    }
    
    key = args[0]
    
    valbytes, err := stub.GetState(key)
    if err != nil {
        jsonResp := "{\"Error\":\"Failed to get state for " + key + "\"}"
        return nil, errors.New(jsonResp)
    }
    
    if valbytes == nil {
        jsonResp := "{\"Error\":\"Nil amount for " + key + "\"}"
        return nil, errors.New(jsonResp)
    }
    
    err = json.Unmarshal(valbytes, &value)
    
    jsonResp := "{\"Name\":\"" + key + "\",\"Amount\":\"" + string(valbytes) + "\"}"
    fmt.Printf("Query Response:%s\n", jsonResp)
    return valbytes, nil
}

func main() {
    err := shim.Start(new(SimpleChaincode))
    if err != nil {
        fmt.Printf("Error starting Simple chaincode: %s", err)
    }
}





