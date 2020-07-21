package main

import (
	"crypto/x509"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/common/util"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/chaincode/shim/ext/cid"
	sc "github.com/hyperledger/fabric/protos/peer"
	"strconv"
)

// Chaincode is the definition of the chaincode structure.
type Chaincode struct {
}

// Definition of the Certificate structure
type certificate struct {
	Type         string `json:"Type"`
	CertID       string `json:"CertID"`
	DateTime     int    `json:"DateTime"`
	VehicleRegNo string `json:"VehicleRegNo"`
	Employee     string `json:"Employee"`
	Content      string `json:"Content"`
}

// Init is called when the chaincode is instantiated by the blockchain network.
func (cc *Chaincode) Init(stub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

// Invoke is called as a result of an application request to run the chaincode.
func (cc *Chaincode) Invoke(stub shim.ChaincodeStubInterface) sc.Response {
	fcn, params := stub.GetFunctionAndParameters()
	fmt.Println("Invoke()", fcn, params)

	if fcn == "createCert" {
		return cc.createCert(stub, params)
	} else if fcn == "createCert" {
		return cc.getCert(stub, params)
	} else {
		fmt.Println("Invoke() did not find func: " + fcn)
		return shim.Error("Received unknown function invocation!")
	}
}

// Function to add new Certificate
func (cc *Chaincode) createCert(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, sender, err := getTxCreatorInfo(stub)
	if !authenticatePollution(creatorOrg, creatorCertIssuer) {
		return shim.Error("{\"Error\":\"Access Denied!\",\"Payload\":{\"MSP\":\"" + creatorOrg + "\",\"CA\":\"" + creatorCertIssuer + "\"}}")
	}

	// Check if sufficient Params passed
	if len(params) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	// Check if Params are non-empty
	for a := 0; a < 4; a++ {
		if len(params[a]) <= 0 {
			return shim.Error("Argument must be a non-empty string")
		}
	}

	CertID := params[0]
	DateTime := params[1]
	VehicleRegNo := params[2]
	Content := params[3]
	Employee := sender
	DateTimeI, err := strconv.Atoi(DateTime)
	if err != nil {
		return shim.Error("Error: Invalid DateTime!")
	}

	// Check if Certificate exists with Key => params[0]
	certificateAsBytes, err := stub.GetState(CertID)
	if err != nil {
		return shim.Error("Failed to check if Certificate exists!")
	} else if certificateAsBytes != nil {
		return shim.Error("Certificate Already Exists!")
	}

	// Generate Certificate from params provided
	certificate := &certificate{"POLUTN_CERT",
		CertID, DateTimeI, VehicleRegNo, Employee, Content}

	// Get JSON bytes of Certificate struct
	certificateJSONasBytes, err := json.Marshal(certificate)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated Certificate with Key => params[0]
	err = stub.PutState(CertID, certificateJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Add Certificate ID to The Investigation with InvestigationID
	/* UPDATE THIS WHEN VEHICLE_CC IS COMPLETE!
	args := util.ToChaincodeArgs("addCertificate", InvestigationID, ID)
	response := stub.InvokeChaincode("investigation_cc", args, "mainchannel")
	if response.Status != shim.OK {
		return shim.Error(response.Message)
	}
	*/

	// Returned on successful execution of the function
	return shim.Success(nil)
}

// Function to read an Certificate
func (cc *Chaincode) getCert(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check if sufficient Params passed
	if len(params) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	// Check if Params are non-empty
	if len(params[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}

	// Get State of Certificate with Key => params[0]
	certificateAsBytes, err := stub.GetState(params[0])
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + params[0] + "\"}"
		return shim.Error(jsonResp)
	} else if certificateAsBytes == nil {
		jsonResp := "{\"Error\":\"Certificate does not exist!\"}"
		return shim.Error(jsonResp)
	}

	// Returned on successful execution of the function
	return shim.Success(certificateAsBytes)
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
func authenticatePollution(mspID string, certCN string) bool {
	return (mspID == "PoliceMSP") && (certCN == "ca.police.example.com") || (mspID == "ForensicsMSP") && (certCN == "ca.forensics.example.com") || (mspID == "CitizenMSP") && (certCN == "ca.citizen.example.com")
}
