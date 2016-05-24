/*
 https://github.com/IBM-Blockchain/marbles-chaincode/blob/master/hyperledger/part2/part2_chaincode.go
 を元に改変
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
    var A string    // ユーザ
    var Aval int // 値
    var err error
    
    if len(args) != 2 {
        return nil, errors.New("Incorrect number of arguments. Expecting 4")
    }
    
    A = args[0]
    Aval, err = strconv.Atoi(args[1])
    if err != nil {
        return nil, errors.New("Expecting integer value for asset holding")
    }
    
    fmt.Printf("Aval = %d,\n", Aval)
    
    // ブロックチェーンへの書き込み
    err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
    if err != nil {
        return nil, err
    }
    return nil, nil
}



// 単純な値の書き換え
func (t *SimpleChaincode) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
    var A string    // ユーザ
    var Aval int
    var err error
    
    if len(args) != 2 {
        return nil, errors.New("Incorrect number of arguments. Expecting 2")
    }
    
    A = args[0]
    Aval, err = strconv.Atoi(args[1])
    if err != nil {
        return nil, errors.New("Expecting integer value for asset holding")
    }
    
    // ブロックチェーン上の値の書き換え
    err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
    if err != nil {
        return nil, err
    }
    
    return nil, nil
}


// 値の取得
func (t *SimpleChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
    if function != "query" {
        return nil, errors.New("Invalid query function name. Expecting \"query\"")
    }
    var A string // ユーザ
    var err error
    
    if len(args) != 1 {
        return nil, errors.New("Incorrect number of arguments. Expecting name of the person to query")
    }
    
    A = args[0]
    
    // ブロックチェーン上のユーザの値の取得
    Avalbytes, err := stub.GetState(A)
    if err != nil {
        jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
        return nil, errors.New(jsonResp)
    }
    
    if Avalbytes == nil {
        jsonResp := "{\"Error\":\"Nil amount for " + A + "\"}"
        return nil, errors.New(jsonResp)
    }
    
    message := "Name:" + A + ", Amount:" + string(Avalbytes) + "}"
    return  []byte(jsonResp), nil
}


func main() {
    err := shim.Start(new(SimpleChaincode))
    if err != nil {
        fmt.Printf("Error starting Simple chaincode: %s", err)
    }
}