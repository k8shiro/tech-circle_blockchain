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

// ============================================================================================================================
// Main
// ============================================================================================================================

func main() {
    err := shim.Start(new(SimpleChaincode))
    if err != nil {
        fmt.Printf("Error starting Simple chaincode: %s", err)
    }
}

func (t *SimpleChaincode) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
    var A string
    var Aval int
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
    err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
    if err != nil {
        return nil, err
    }
    return nil, nil
}
