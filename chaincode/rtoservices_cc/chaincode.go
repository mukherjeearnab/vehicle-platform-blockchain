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

type status struct {
	Date     int    `json:"Date"`
	Content  string `json:"Content"`
	Employee string `json:"Employee"`
}

type license struct {
	License  string `json:"License"`
	Employee string `json:"Employee"`
}

type registration struct {
	RegNo    int    `json:"RegNo"`
	Employee string `json:"Employee"`
}

// Definition of the Driver's License Application structure
type dlapp struct {
	Type          string   `json:"Type"`
	ApplicationID string   `json:"ApplicationID"`
	DateTime      int      `json:"DateTime"`
	UID           string   `json:"UID"`
	RtoID         string   `json:"RtoID"`
	Status        []status `json:"Status"`
	License       license  `json:"License"`
}

// Definition of the Vehicle Registration Application structure
type vrapp struct {
	Type          string       `json:"Type"`
	ApplicationID string       `json:"ApplicationID"`
	DateTime      int          `json:"DateTime"`
	Make          string       `json:"Make"`
	Model         string       `json:"Model"`
	ModelVariant  string       `json:"ModelVariant"`
	ModelYear     string       `json:"ModelYear"`
	Color         string       `json:"Color"`
	EngineNo      string       `json:"EngineNo"`
	ChassisNo     string       `json:"ChassisNo"`
	Owner         string       `json:"Owner"`
	Creator       string       `json:"Creator"`
	Status        []status     `json:"Status"`
	Registration  registration `json:"Registration"`
}

// Definition of the Driver's License structure
type dlicense struct {
	Type          string `json:"Type"`
	LicenseNumber string `json:"LicenseNumber"`
	UID           string `json:"UID"`
	VehicleType   string `json:"VehicleType"`
	ExpiryDate    int    `json:"ExpiryDate"`
	Employee      string `json:"Employee"`
}

// Init is called when the chaincode is instantiated by the blockchain network.
func (cc *Chaincode) Init(stub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

// Invoke is called as a result of an application request to run the chaincode.
func (cc *Chaincode) Invoke(stub shim.ChaincodeStubInterface) sc.Response {
	fcn, params := stub.GetFunctionAndParameters()
	fmt.Println("Invoke()", fcn, params)

	if fcn == "newDLApplication" {
		return cc.newDLApplication(stub, params)
	} else if fcn == "newVRApplication" {
		return cc.newVRApplication(stub, params)
	} else if fcn == "getDLVRApplication" {
		return cc.getDLVRApplication(stub, params)
	} else if fcn == "updateDLApplication" {
		return cc.updateDLApplication(stub, params)
	} else if fcn == "updateVRApplication" {
		return cc.updateVRApplication(stub, params)
	} else if fcn == "createDL" {
		return cc.createDL(stub, params)
	} else if fcn == "getDL" {
		return cc.getDL(stub, params)
	} else {
		fmt.Println("Invoke() did not find func: " + fcn)
		return shim.Error("Received unknown function invocation!")
	}
}

// Function to add new DL Application
func (cc *Chaincode) newDLApplication(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, creator, err := getTxCreatorInfo(stub)
	if !authenticateCitizen(creatorOrg, creatorCertIssuer) {
		return shim.Error("{\"Error\":\"Access Denied!\",\"Payload\":{\"MSP\":\"" + creatorOrg + "\",\"CA\":\"" + creatorCertIssuer + "\"}}")
	}

	// Check if sufficient Params passed
	if len(params) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	// Check if Params are non-empty
	for a := 0; a < 3; a++ {
		if len(params[a]) <= 0 {
			return shim.Error("Argument must be a non-empty string")
		}
	}

	ApplicationID := params[0]
	DateTime := params[1]
	UID := creator
	RtoID := params[2]
	var Status []status
	var License license

	DateTimeI, err := strconv.Atoi(DateTime)
	if err != nil {
		return shim.Error("Error: Invalid DateTime!")
	}

	// Check if Application exists with Key => params[0]
	applicationAsBytes, err := stub.GetState(ApplicationID)
	if err != nil {
		return shim.Error("Failed to check if Application exists!")
	} else if applicationAsBytes != nil {
		return shim.Error("Application Already Exists!")
	}

	// Generate Application from params provided
	application := &dlapp{"DRVLCN_RA",
		ApplicationID, DateTimeI, UID, RtoID, Status, License}

	// Get JSON bytes of Application struct
	applicationJSONasBytes, err := json.Marshal(application)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated Application with Key => params[0]
	err = stub.PutState(ApplicationID, applicationJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Returned on successful execution of the function
	return shim.Success(nil)
}

// Function to add new VR Application
func (cc *Chaincode) newVRApplication(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, creator, err := getTxCreatorInfo(stub)
	if !authenticateCitizen(creatorOrg, creatorCertIssuer) {
		return shim.Error("{\"Error\":\"Access Denied!\",\"Payload\":{\"MSP\":\"" + creatorOrg + "\",\"CA\":\"" + creatorCertIssuer + "\"}}")
	}

	// Check if sufficient Params passed
	if len(params) != 10 {
		return shim.Error("Incorrect number of arguments. Expecting 10")
	}

	// Check if Params are non-empty
	for a := 0; a < 10; a++ {
		if len(params[a]) <= 0 {
			return shim.Error("Argument must be a non-empty string")
		}
	}

	ApplicationID := params[0]
	DateTime := params[1]
	Make := params[2]
	Model := params[3]
	ModelVariant := params[4]
	ModelYear := params[5]
	Color := params[6]
	EngineNo := params[7]
	ChassisNo := params[8]
	Owner := params[9]
	var Status []status
	var Registration registration

	DateTimeI, err := strconv.Atoi(DateTime)
	if err != nil {
		return shim.Error("Error: Invalid DateTime!")
	}

	// Check if Application exists with Key => params[0]
	applicationAsBytes, err := stub.GetState(ApplicationID)
	if err != nil {
		return shim.Error("Failed to check if Application exists!")
	} else if applicationAsBytes != nil {
		return shim.Error("Application Already Exists!")
	}

	// Generate Application from params provided
	application := &vrapp{"VEHCL_RA",
		ApplicationID, DateTimeI, Make, Model, ModelVariant, ModelYear,
		Color, EngineNo, ChassisNo, Owner, creator, Status, Registration}

	// Get JSON bytes of Application struct
	applicationJSONasBytes, err := json.Marshal(application)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated Application with Key => params[0]
	err = stub.PutState(ApplicationID, applicationJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Returned on successful execution of the function
	return shim.Success(nil)
}

// Function to read an Application
func (cc *Chaincode) getDLVRApplication(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check if sufficient Params passed
	if len(params) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	// Check if Params are non-empty
	if len(params[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}

	// Get State of Application with Key => params[0]
	applicationAsBytes, err := stub.GetState(params[0])
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + params[0] + "\"}"
		return shim.Error(jsonResp)
	} else if applicationAsBytes == nil {
		jsonResp := "{\"Error\":\"Application does not exist!\"}"
		return shim.Error(jsonResp)
	}

	// Returned on successful execution of the function
	return shim.Success(applicationAsBytes)
}

// Function to Update DL Application Status
func (cc *Chaincode) updateDLApplication(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, creator, err := getTxCreatorInfo(stub)
	if !authenticateRTO(creatorOrg, creatorCertIssuer) {
		return shim.Error("{\"Error\":\"Access Denied!\",\"Payload\":{\"MSP\":\"" + creatorOrg + "\",\"CA\":\"" + creatorCertIssuer + "\"}}")
	}

	// Check if sufficient Params passed
	if len(params) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	// Check if Params are non-empty
	for a := 0; a < 3; a++ {
		if len(params[a]) <= 0 {
			return shim.Error("Argument must be a non-empty string")
		}
	}

	// Copy the Values from params[]
	ApplicationID := params[0]
	Date := params[1]
	Content := params[2]
	Employee := creator
	DateI, err := strconv.Atoi(Date)
	if err != nil {
		return shim.Error("Error: Invalid Date!")
	}

	// Check if Application exists with Key => ApplicationID
	applicationAsBytes, err := stub.GetState(ApplicationID)
	if err != nil {
		return shim.Error("Failed to get Application Details!")
	} else if applicationAsBytes == nil {
		return shim.Error("Error: Application Does NOT Exist!")
	}

	// Create Update struct var
	applicationToUpdate := dlapp{}
	err = json.Unmarshal(applicationAsBytes, &applicationToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	NewStatus := status{DateI, Content, Employee}

	// Update Application.Status to append => NewStatus
	applicationToUpdate.Status = append(applicationToUpdate.Status, NewStatus)

	// Get JSON bytes of Vehicle struct
	applicationJSONasBytes, err := json.Marshal(applicationToUpdate)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated Vehicle with Key => params[0]
	err = stub.PutState(ApplicationID, applicationJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Returned on successful execution of the function
	return shim.Success(nil)
}

// Function to Update VR Application Status
func (cc *Chaincode) updateVRApplication(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, creator, err := getTxCreatorInfo(stub)
	if !authenticateRTO(creatorOrg, creatorCertIssuer) {
		return shim.Error("{\"Error\":\"Access Denied!\",\"Payload\":{\"MSP\":\"" + creatorOrg + "\",\"CA\":\"" + creatorCertIssuer + "\"}}")
	}

	// Check if sufficient Params passed
	if len(params) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	// Check if Params are non-empty
	for a := 0; a < 3; a++ {
		if len(params[a]) <= 0 {
			return shim.Error("Argument must be a non-empty string")
		}
	}

	// Copy the Values from params[]
	ApplicationID := params[0]
	Date := params[1]
	Content := params[2]
	Employee := creator
	DateI, err := strconv.Atoi(Date)
	if err != nil {
		return shim.Error("Error: Invalid Date!")
	}

	// Check if Application exists with Key => ApplicationID
	applicationAsBytes, err := stub.GetState(ApplicationID)
	if err != nil {
		return shim.Error("Failed to get Application Details!")
	} else if applicationAsBytes == nil {
		return shim.Error("Error: Application Does NOT Exist!")
	}

	// Create Update struct var
	applicationToUpdate := dlapp{}
	err = json.Unmarshal(applicationAsBytes, &applicationToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	NewStatus := status{DateI, Content, Employee}

	// Update Application.Status to append => NewStatus
	applicationToUpdate.Status = append(applicationToUpdate.Status, NewStatus)

	// Get JSON bytes of Vehicle struct
	applicationJSONasBytes, err := json.Marshal(applicationToUpdate)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated Vehicle with Key => params[0]
	err = stub.PutState(ApplicationID, applicationJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Returned on successful execution of the function
	return shim.Success(nil)
}

// Function to Create new DL
func (cc *Chaincode) createDL(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, creator, err := getTxCreatorInfo(stub)
	if !authenticateRTO(creatorOrg, creatorCertIssuer) {
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

	LicenseNumber := params[0]
	UID := params[1]
	VehicleType := params[2]
	ExpiryDate := params[3]

	ExpiryDateI, err := strconv.Atoi(ExpiryDate)
	if err != nil {
		return shim.Error("Error: Invalid ExpiryDate!")
	}

	// Check if License exists with Key => params[0]
	licenseAsBytes, err := stub.GetState(LicenseNumber)
	if err != nil {
		return shim.Error("Failed to check if License exists!")
	} else if licenseAsBytes != nil {
		return shim.Error("License Already Exists!")
	}

	// Generate License from params provided
	license := &dlicense{"DRVLCN",
		LicenseNumber, UID, VehicleType, ExpiryDateI, creator}

	// Get JSON bytes of License struct
	licenseJSONasBytes, err := json.Marshal(license)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated License with Key => params[0]
	err = stub.PutState(LicenseNumber, licenseJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Add License Number to The Citizen Profile with UID
	args := util.ToChaincodeArgs("addLicense", UID, LicenseNumber)
	response := stub.InvokeChaincode("profilemanager_cc", args, "mainchannel")
	if response.Status != shim.OK {
		return shim.Error(response.Message)
	}

	// Returned on successful execution of the function
	return shim.Success(nil)
}

// Function to read an Driver's License
func (cc *Chaincode) getDL(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check if sufficient Params passed
	if len(params) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	// Check if Params are non-empty
	if len(params[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}

	// Get State of Application with Key => params[0]
	applicationAsBytes, err := stub.GetState(params[0])
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + params[0] + "\"}"
		return shim.Error(jsonResp)
	} else if applicationAsBytes == nil {
		jsonResp := "{\"Error\":\"Application does not exist!\"}"
		return shim.Error(jsonResp)
	}

	// Returned on successful execution of the function
	return shim.Success(applicationAsBytes)
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

// Authenticate => Citizen
func authenticateCitizen(mspID string, certCN string) bool {
	return (mspID == "CitizenMSP") && (certCN == "ca.citizen.vehicle.com")
}

// Authenticate => RTO
func authenticateRTO(mspID string, certCN string) bool {
	return (mspID == "RTOMSP") && (certCN == "ca.rto.vehicle.com")
}
