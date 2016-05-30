/*
 https://github.com/IBM-Blockchain/marbles-chaincode/blob/master/hyperledger/part2/part2_chaincode.go
 */

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
    var id string
    var temperature int
    var err error
    
    if len(args) != 2 {
        return nil, errors.New("Incorrect number of arguments. Expecting 2")
    }
    
    id = args[0]
    temperature, err = strconv.Atoi(args[1])
    if err != nil {
        return nil, err
    }
    
    fmt.Printf("temperature = %d,\n", temperature)
    
    err = stub.PutState(id, []byte(strconv.Itoa(temperature)))
    if err != nil {
        return nil, err
    }

    return nil, nil
}



func (t *SimpleChaincode) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
    var id string
    var temperature int
    var err error
    
    if len(args) != 2 {
        return nil, errors.New("Incorrect number of arguments. Expecting 2")
    }
    
    id = args[0]
    temperature, err = strconv.Atoi(args[1])
    if err != nil {
        return nil, err
    }
    
    err = stub.PutState(id, []byte(strconv.Itoa(temperature)))
    if err != nil {
        return nil, err
    }
    
    return nil, nil
}


func (t *SimpleChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
    if function != "query" {
        return nil, errors.New("Invalid query function name. Expecting \"query\"")
    }
    var id string
    var err error
    
    if len(args) != 1 {
        return nil, errors.New("Incorrect number of arguments. Expecting name of the person to query")
    }
    
    id = args[0]
    temperatureBytes, err := stub.GetState(id)
    if err != nil {
        return nil, err
    }
    
    if temperatureBytes == nil {
        jsonResp := "{\"Error\":\"Nil amount for " + id + "\"}"
        return nil, errors.New(jsonResp)
    }
    
    message := "ID:" + id + ", Temperature:" + string(temperatureBytes)
    return  []byte(message), nil
}


func main() {
    err := shim.Start(new(SimpleChaincode))
    if err != nil {
        fmt.Printf("Error starting Simple chaincode: %s", err)
    }
}