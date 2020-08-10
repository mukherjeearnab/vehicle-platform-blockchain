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

// Definition of the Policy structure
type policy struct {
	Type         string   `json:"Type"`
	InsuranceID  string   `json:"InsuranceID"`
	VehicleRegNo string   `json:"VehicleRegNo"`
	UID          string   `json:"UID"`
	Date         int      `json:"Date"`
	Duration     string   `json:"Duration"`
	Content      string   `json:"Content"`
	Claims       []string `json:"Claims"`
}

type status struct {
	Date     int    `json:"Date"`
	Content  string `json:"Content"`
	Employee string `json:"Employee"`
}

//Definition of the Claim structure
type claim struct {
	Type              string   `json:"Type"`
	ClaimID           string   `json:"ClaimID"`
	InsurancePolicyID string   `json:"InsurancePolicyID"`
	DateTime          int      `json:"DateTime"`
	Content           string   `json:"Content"`
	Status            []status `json:"Status"`
	Approve           bool     `json:"Approve"`
}

// Init is called when the chaincode is instantiated by the blockchain network.
func (cc *Chaincode) Init(stub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

// Invoke is called as a result of an application request to run the chaincode.
func (cc *Chaincode) Invoke(stub shim.ChaincodeStubInterface) sc.Response {
	fcn, params := stub.GetFunctionAndParameters()
	fmt.Println("Invoke()", fcn, params)

	if fcn == "addInsurancePolicy" {
		return cc.addInsurancePolicy(stub, params)
	} else if fcn == "claimInsurancePolicy" {
		return cc.claimInsurancePolicy(stub, params)
	} else if fcn == "readInsurancePolicyClaim" {
		return cc.readInsurancePolicyClaim(stub, params)
	} else if fcn == "updateInsuranceClaim" {
		return cc.updateInsuranceClaim(stub, params)
	} else if fcn == "approveInsuranceClaim" {
		return cc.approveInsuranceClaim(stub, params)
	} else {
		fmt.Println("Invoke() did not find func: " + fcn)
		return shim.Error("Received unknown function invocation!")
	}
}

// Function to add new Insurance
func (cc *Chaincode) addInsurancePolicy(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, err := getTxCreatorInfo(stub)
	if !authenticateInsurance(creatorOrg, creatorCertIssuer) {
		return shim.Error("{\"Error\":\"Access Denied!\",\"Payload\":{\"MSP\":\"" + creatorOrg + "\",\"CA\":\"" + creatorCertIssuer + "\"}}")
	}

	// Check if sufficient Params passed
	if len(params) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 6")
	}

	// Check if Params are non-empty
	for a := 0; a < 6; a++ {
		if len(params[a]) <= 0 {
			return shim.Error("Argument must be a non-empty string")
		}
	}

	InsuranceID := params[0]
	VehicleRegNo := params[1]
	UID := params[2]
	Date := params[3]
	Duration := params[4]
	Content := params[5]
	var Claims []string

	DateI, err := strconv.Atoi(Date)
	if err != nil {
		return shim.Error("Error: Invalid Date!")
	}

	// Check if Insurance exists with Key => params[0]
	policyAsBytes, err := stub.GetState(InsuranceID)
	if err != nil {
		return shim.Error("Failed to check if Insurance exists!")
	} else if policyAsBytes != nil {
		return shim.Error("Insurance Already Exists!")
	}

	// Generate Insurance from params provided
	policy := &policy{"INSRNS_PLCY",
		InsuranceID, VehicleRegNo, UID, DateI, Duration, Content, Claims}

	// Get JSON bytes of Insurance struct
	policyJSONasBytes, err := json.Marshal(policy)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated Insurance with Key => params[0]
	err = stub.PutState(InsuranceID, policyJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Add Insurance ID to The Vehicle with RegNo
	args := util.ToChaincodeArgs("addInsurancePolicy", VehicleRegNo, InsuranceID)
	response := stub.InvokeChaincode("vehicle_cc", args, "mainchannel")
	if response.Status != shim.OK {
		return shim.Error(response.Message)
	}

	// Returned on successful execution of the function
	return shim.Success(nil)
}

// Function to read an Insurance
func (cc *Chaincode) readInsurancePolicyClaim(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check if sufficient Params passed
	if len(params) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	// Check if Params are non-empty
	if len(params[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}

	// Get State of Insurance with Key => params[0]
	policyAsBytes, err := stub.GetState(params[0])
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + params[0] + "\"}"
		return shim.Error(jsonResp)
	} else if policyAsBytes == nil {
		jsonResp := "{\"Error\":\"Insurance does not exist!\"}"
		return shim.Error(jsonResp)
	}

	// Returned on successful execution of the function
	return shim.Success(policyAsBytes)
}

// Function to add new Claim
func (cc *Chaincode) claimInsurancePolicy(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, err := getTxCreatorInfo(stub)
	if !authenticateCitizen(creatorOrg, creatorCertIssuer) {
		return shim.Error("{\"Error\":\"Access Denied!\",\"Payload\":{\"MSP\":\"" + creatorOrg + "\",\"CA\":\"" + creatorCertIssuer + "\"}}")
	}

	// Check if sufficient Params passed
	if len(params) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	// Check if Params are non-empty
	for a := 0; a < 4; a++ {
		if len(params[a]) <= 0 {
			return shim.Error("Argument must be a non-empty string")
		}
	}

	ClaimID := params[0]
	InsurancePolicyID := params[1]
	DateTime := params[2]
	Content := params[3]
	var Status []status
	Approve := false

	DateTimeI, err := strconv.Atoi(DateTime)
	if err != nil {
		return shim.Error("Error: Invalid DateTime!")
	}

	// Check if Policy exists with Key => InsurancePolicyID
	policyAsBytes, err := stub.GetState(InsurancePolicyID)
	if err != nil {
		return shim.Error("Failed to get Policy Details!")
	} else if policyAsBytes == nil {
		return shim.Error("Error: Policy Does NOT Exist!")
	}

	// Check if Insurance exists with Key => params[0]
	claimAsBytes, err := stub.GetState(ClaimID)
	if err != nil {
		return shim.Error("Failed to check if Insurance exists!")
	} else if claimAsBytes != nil {
		return shim.Error("Insurance Already Exists!")
	}

	// Generate Insurance from params provided
	claim := &claim{"INSRNS_CLM",
		ClaimID, InsurancePolicyID, DateTimeI, Content, Status, Approve}

	// Get JSON bytes of Insurance struct
	claimJSONasBytes, err := json.Marshal(claim)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated Insurance with Key => params[0]
	err = stub.PutState(ClaimID, claimJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Create Update struct var
	policyToUpdate := policy{}
	err = json.Unmarshal(policyAsBytes, &policyToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	// Update Policy.Claims to append => ClaimID
	policyToUpdate.Claims = append(policyToUpdate.Claims, ClaimID)

	// Get JSON bytes of Vehicle struct
	policyJSONasBytes, err := json.Marshal(policyToUpdate)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated Policy with Key => params[0]
	err = stub.PutState(InsurancePolicyID, policyJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Returned on successful execution of the function
	return shim.Success(nil)
}

// Function to Update Insurance Claim Status
func (cc *Chaincode) updateInsuranceClaim(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, creator, err := getTxCreatorInfoC(stub)
	if !authenticateInsurance(creatorOrg, creatorCertIssuer) {
		return shim.Error("{\"Error\":\"Access Denied!\",\"Payload\":{\"MSP\":\"" + creatorOrg + "\",\"CA\":\"" + creatorCertIssuer + "\"}}")
	}

	// Check if sufficient Params passed
	if len(params) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 9")
	}

	// Check if Params are non-empty
	for a := 0; a < 3; a++ {
		if len(params[a]) <= 0 {
			return shim.Error("Argument must be a non-empty string")
		}
	}

	// Copy the Values from params[]
	ClaimID := params[0]
	Date := params[1]
	Content := params[2]
	Employee := creator
	DateI, err := strconv.Atoi(Date)
	if err != nil {
		return shim.Error("Error: Invalid Date!")
	}

	// Check if Claim exists with Key => ClaimID
	claimAsBytes, err := stub.GetState(ClaimID)
	if err != nil {
		return shim.Error("Failed to get Claim Details!")
	} else if claimAsBytes == nil {
		return shim.Error("Error: Claim Does NOT Exist!")
	}

	// Create Update struct var
	claimToUpdate := claim{}
	err = json.Unmarshal(claimAsBytes, &claimToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	NewStatus := status{DateI, Content, Employee}

	// Update Claim.Status to append => NewStatus
	claimToUpdate.Status = append(claimToUpdate.Status, NewStatus)

	// Get JSON bytes of Vehicle struct
	claimJSONasBytes, err := json.Marshal(claimToUpdate)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated Vehicle with Key => params[0]
	err = stub.PutState(ClaimID, claimJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Returned on successful execution of the function
	return shim.Success(nil)
}

// Function to Approve Insurance Claim
func (cc *Chaincode) approveInsuranceClaim(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, creator, err := getTxCreatorInfoC(stub)
	if !authenticateInsurance(creatorOrg, creatorCertIssuer) {
		return shim.Error("{\"Error\":\"Access Denied!\",\"Payload\":{\"MSP\":\"" + creatorOrg + "\",\"CA\":\"" + creatorCertIssuer + "\"}}")
	}

	// Check if sufficient Params passed
	if len(params) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	// Check if Params are non-empty
	for a := 0; a < 2; a++ {
		if len(params[a]) <= 0 {
			return shim.Error("Argument must be a non-empty string")
		}
	}

	// Copy the Values from params[]
	ClaimID := params[0]
	Date := params[1]
	Content := "Insurance Claim Approved!"
	Employee := creator
	DateI, err := strconv.Atoi(Date)
	if err != nil {
		return shim.Error("Error: Invalid Date!")
	}

	// Check if Claim exists with Key => ClaimID
	claimAsBytes, err := stub.GetState(ClaimID)
	if err != nil {
		return shim.Error("Failed to get Claim Details!")
	} else if claimAsBytes == nil {
		return shim.Error("Error: Claim Does NOT Exist!")
	}

	// Create Update struct var
	claimToUpdate := claim{}
	err = json.Unmarshal(claimAsBytes, &claimToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	NewStatus := status{DateI, Content, Employee}

	// Update Claim.Status to append => NewStatus
	claimToUpdate.Status = append(claimToUpdate.Status, NewStatus)
	// Update Claim.Approve = true
	claimToUpdate.Approve = true

	// Get JSON bytes of Vehicle struct
	claimJSONasBytes, err := json.Marshal(claimToUpdate)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated Vehicle with Key => params[0]
	err = stub.PutState(ClaimID, claimJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Returned on successful execution of the function
	return shim.Success(nil)
}

// ---------------------------------------------
// Helper Functions
// ---------------------------------------------

// Authentication
// ++++++++++++++

// Get Tx Creator Info
func getTxCreatorInfo(stub shim.ChaincodeStubInterface) (string, string, error) {
	var mspid string
	var err error
	var cert *x509.Certificate
	mspid, err = cid.GetMSPID(stub)

	if err != nil {
		fmt.Printf("Error getting MSP identity: %sn", err.Error())
		return "", "", err
	}

	cert, err = cid.GetX509Certificate(stub)
	if err != nil {
		fmt.Printf("Error getting client certificate: %sn", err.Error())
		return "", "", err
	}

	return mspid, cert.Issuer.CommonName, nil
}

// Get Tx Creator Info with CommonName
func getTxCreatorInfoC(stub shim.ChaincodeStubInterface) (string, string, string, error) {
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

// Authenticate => Insurance
func authenticateInsurance(mspID string, certCN string) bool {
	return (mspID == "InsuranceMSP") && (certCN == "ca.insurance.vehicle.com")
}

// Authenticate => Citizen
func authenticateCitizen(mspID string, certCN string) bool {
	return (mspID == "CitizenMSP") && (certCN == "ca.citizen.vehicle.com")
}
