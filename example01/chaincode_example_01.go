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

// Transaction makes payment of X units from itemID to B
func (t *SimpleChaincode1) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
    var itemID string    // Entities
    var item string // Asset holdings
    var err error
    
    if len(args) != 2 {
        return nil, errors.New("Incorrect number of arguments. Expecting 4")
    }
    
    // Initialize the chaincode
    itemID = args[0]
    item, err = args[1]
    if err != nil {
        return nil, errors.New("Expecting integer value for asset holding")
    }
    
    
    // Write the state to the ledger
    err = stub.PutState(itemID, []byte(item))
    if err != nil {
        return nil, err
    }
    
    
    return nil, nil
}


// Query callback representing the query of a chaincode
func (t *SimpleChaincode1) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {

    return []byte("OK"), nil
}

func main() {
    err := shim.Start(new(SimpleChaincode1))
    if err != nil {
        fmt.Printf("Error starting Simple chaincode: %s", err)
    }
}


