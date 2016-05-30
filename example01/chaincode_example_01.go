/*
 https://github.com/IBM-Blockchain/marbles-chaincode/blob/master/hyperledger/part2/part2_chaincode.go
 */

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



type User struct {
    Name string
    Age  int
}



func (t *SimpleChaincode) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
    var A string
    var Aval User
    var err error
    
    if len(args) != 3 {
        return nil, errors.New("Incorrect number of arguments. Expecting 3")
    }
    
    A = args[0]
    Aval.Name = args[1]
    Aval.Age, err = strconv.Atoi(args[2])
    if err != nil {
        return nil, errors.New("Expecting integer value for asset holding")
    }
    
    Avalbytes, err := json.Marshal(Aval)
    if err != nil {
        return nil, errors.New("err")
    }
    
    err = stub.PutState(A, Avalbytes)
    if err != nil {
        return nil, err
    }
    return nil, nil
}



func (t *SimpleChaincode) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
    var A string
    var Aval User
    var err error
    
    if len(args) != 3 {
        return nil, errors.New("Incorrect number of arguments. Expecting 3")
    }
    
    A = args[0]
    Aval.Name = args[1]
    Aval.Age, err = strconv.Atoi(args[2])
    if err != nil {
        return nil, errors.New("Expecting integer value for asset holding")
    }
    
    Avalbytes, err := json.Marshal(Aval)
    if err != nil {
        return nil, errors.New("err")
    }
    
    
    err = stub.PutState(A, Avalbytes)
    if err != nil {
        return nil, err
    }
    
    return nil, nil
}



func (t *SimpleChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
    if function != "query" {
        return nil, errors.New("Invalid query function name. Expecting \"query\"")
    }
    
    var A string
    var Aval User
    var err error
    
    if len(args) != 1 {
        return nil, errors.New("Incorrect number of arguments. Expecting name of the person to query")
    }
    
    A = args[0]
    
    
    Avalbytes, err := stub.GetState(A)
    if err != nil {
        jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
        return nil, errors.New(jsonResp)
    }
    
    if Avalbytes == nil {
        jsonResp := "{\"Error\":\"Nil amount for " + A + "\"}"
        return nil, errors.New(jsonResp)
    }
    
    err = json.Unmarshal(Avalbytes, &Aval)
    if err != nil {
        fmt.Errorf("%s", err)
    }
    
    
    message := "ID:" + A + ", Name:" + Aval.Name + ", Age:"+ strconv.Itoa(Aval.Age) +"}"
    return  []byte(message), nil
}


func main() {
    err := shim.Start(new(SimpleChaincode))
    if err != nil {
        fmt.Printf("Error starting Simple chaincode: %s", err)
    }
}