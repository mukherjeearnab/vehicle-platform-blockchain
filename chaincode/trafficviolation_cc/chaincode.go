package main

import (
	"crypto/x509"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/chaincode/shim/ext/cid"
	sc "github.com/hyperledger/fabric/protos/peer"
	"strconv"
)

// Chaincode is the definition of the chaincode structure.
type Chaincode struct {
}

// Definition of the TV structure
type tv struct {
	Type          string `json:"Type"`
	TVID          string `json:"TVID"`
	DateIncident  int    `json:"DateIncident"`
	DateFiling    int    `json:"DateFiling"`
	OfficerID     string `json:"OfficerID"`
	LicenseNumber string `json:"LicenseNumber"`
	VehicleRegNo  string `json:"VehicleRegNo"`
	Content       string `json:"Content"`
	Evidence      string `json:"Evidence"`
}

// Init is called when the chaincode is instantiated by the blockchain network.
func (cc *Chaincode) Init(stub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

// Invoke is called as a result of an application request to run the chaincode.
func (cc *Chaincode) Invoke(stub shim.ChaincodeStubInterface) sc.Response {
	fcn, params := stub.GetFunctionAndParameters()
	fmt.Println("Invoke()", fcn, params)

	if fcn == "fileTV" {
		return cc.fileTV(stub, params)
	} else if fcn == "getTV" {
		return cc.getTV(stub, params)
	} else {
		fmt.Println("Invoke() did not find func: " + fcn)
		return shim.Error("Received unknown function invocation!")
	}
}

// Function to add new Certificate
func (cc *Chaincode) fileTV(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, creator, err := getTxCreatorInfo(stub)
	if !authenticatePolice(creatorOrg, creatorCertIssuer) {
		return shim.Error("{\"Error\":\"Access Denied!\",\"Payload\":{\"MSP\":\"" + creatorOrg + "\",\"CA\":\"" + creatorCertIssuer + "\"}}")
	}

	// Check if sufficient Params passed
	if len(params) != 7 {
		return shim.Error("Incorrect number of arguments. Expecting 6")
	}

	// Check if Params are non-empty
	for a := 0; a < 7; a++ {
		if len(params[a]) <= 0 {
			return shim.Error("Argument must be a non-empty string")
		}
	}

	TVID := params[0]
	DateIncident := params[1]
	DateFiling := params[2]
	OfficerID := creator
	LicenseNumber := params[3]
	VehicleRegNo := params[4]
	Content := params[5]
	Evidence := params[6]

	DateIncidentI, err := strconv.Atoi(DateIncident)
	if err != nil {
		return shim.Error("Error: Invalid DateIncident!")
	}
	DateFilingI, err := strconv.Atoi(DateFiling)
	if err != nil {
		return shim.Error("Error: Invalid DateFiling!")
	}

	// Check if TV exists with Key => params[0]
	tvAsBytes, err := stub.GetState(TVID)
	if err != nil {
		return shim.Error("Failed to check if TV exists!")
	} else if tvAsBytes != nil {
		return shim.Error("TV Already Exists!")
	}

	// Generate TV from params provided
	tv := &tv{"TRFC_VIO",
		TVID, DateIncidentI, DateFilingI, OfficerID, LicenseNumber, VehicleRegNo, Content, Evidence}

	// Get JSON bytes of TV struct
	tvJSONasBytes, err := json.Marshal(tv)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated TV with Key => params[0]
	err = stub.PutState(TVID, tvJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Returned on successful execution of the function
	return shim.Success(nil)
}

// Function to read an Certificate
func (cc *Chaincode) getTV(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check if sufficient Params passed
	if len(params) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	// Check if Params are non-empty
	if len(params[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}

	// Get State of TV with Key => params[0]
	tvAsBytes, err := stub.GetState(params[0])
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + params[0] + "\"}"
		return shim.Error(jsonResp)
	} else if tvAsBytes == nil {
		jsonResp := "{\"Error\":\"Certificate does not exist!\"}"
		return shim.Error(jsonResp)
	}

	// Returned on successful execution of the function
	return shim.Success(tvAsBytes)
}

// ---------------------------------------------
// Helper Functions
// ---------------------------------------------

// Authentication
// ++++++++++++++

// Get Tx Creator Info
func getTxCreatorInfo(stub shim.ChaincodeStubInterface) (string, string, string, error) {
	var mspid string
	var err error
	var cert *x509.Certificate
	mspid, err = cid.GetMSPID(stub)

	if err != nil {
		fmt.Printf("Error getting MSP identity: %sn", err.Error())
		return "", "", "", err
	}

	cert, err = cid.GetX509Certificate(stub)
	if err != nil {
		fmt.Printf("Error getting client certificate: %sn", err.Error())
		return "", "", "", err
	}

	return mspid, cert.Issuer.CommonName, cert.Subject.CommonName, nil
}

// Authenticate => Pollution
func authenticatePolice(mspID string, certCN string) bool {
	return (mspID == "PoliceMSP") && (certCN == "ca.police.vtan.com")
}
