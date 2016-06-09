package main
/*
 テスト中
 
 */
import (
        "errors"
        "fmt"
        "github.com/hyperledger/fabric/core/chaincode/shim"
        )

type SimpleChaincode struct {
}

var myLogger = logging.MustGetLogger("asset_mgm")

func (t *SimpleChaincode) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {    
    if len(args) != 0 {
        return nil, errors.New("Incorrect number of arguments. Expecting 0")
    }
    
    return nil, nil
}

func (t *SimpleChaincode) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
    var key string
    var value string
    var err error
    
    if len(args) != 2 {
        return nil, errors.New("Incorrect number of arguments. Expecting 2")
    }
    
    key = args[0]
    value = args[1]
    valbyte := []byte(value)
    
    
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
    
    message := "Key:" + key + " Value:" + string(valbytes)
    return []byte(message), nil
}

func main() {
    err := shim.Start(new(SimpleChaincode))
    if err != nil {
        fmt.Printf("Error starting Simple chaincode: %s", err)
    }
}





