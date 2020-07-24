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

type ownership struct {
	Prev     string `json:"Prev"`
	Curr     string `json:"Curr"`
	Employee string `json:"Employee"`
	Date     int    `json:"Date"`
}

// Definition of the Vehicle structure
type vehicle struct {
	Type             string      `json:"Type"`
	RegNo            string      `json:"RegNo"`
	Make             string      `json:"Make"`
	Model            string      `json:"Model"`
	ModelVariant     string      `json:"ModelVariant"`
	ModelYear        string      `json:"ModelYear"`
	Color            string      `json:"Color"`
	EngineNo         string      `json:"EngineNo"`
	ChassisNo        string      `json:"ChassisNo"`
	Owner            string      `json:"Owner"`
	Insurance        []string    `json:"Insurance"`
	PollutionCert    []string    `json:"PollutionCert"`
	OwnershipHistory []ownership `json:"OwnershipHistory"`
}

// Init is called when the chaincode is instantiated by the blockchain network.
func (cc *Chaincode) Init(stub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

// Invoke is called as a result of an application request to run the chaincode.
func (cc *Chaincode) Invoke(stub shim.ChaincodeStubInterface) sc.Response {
	fcn, params := stub.GetFunctionAndParameters()
	fmt.Println("Invoke()", fcn, params)

	if fcn == "createVehicleProfile" {
		return cc.createVehicleProfile(stub, params)
	} else if fcn == "getVehicleProfile" {
		return cc.getVehicleProfile(stub, params)
	} else if fcn == "addPollutionCertificate" {
		return cc.addPollutionCertificate(stub, params)
	} else if fcn == "addInsurancePolicy" {
		return cc.addInsurancePolicy(stub, params)
	} else if fcn == "transferOwnership" {
		return cc.transferOwnership(stub, params)
	} else {
		fmt.Println("Invoke() did not find func: " + fcn)
		return shim.Error("Received unknown function invocation!")
	}
}

// Function to add new Vehicle
func (cc *Chaincode) createVehicleProfile(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, creator, err := getTxCreatorInfo(stub)
	if !authenticateRTO(creatorOrg, creatorCertIssuer) {
		return shim.Error("{\"Error\":\"Access Denied!\",\"Payload\":{\"MSP\":\"" + creatorOrg + "\",\"CA\":\"" + creatorCertIssuer + "\"}}")
	}

	// Check if sufficient Params passed
	if len(params) != 9 {
		return shim.Error("Incorrect number of arguments. Expecting 9")
	}

	// Check if Params are non-empty
	for a := 0; a < 9; a++ {
		if len(params[a]) <= 0 {
			return shim.Error("Argument must be a non-empty string")
		}
	}

	RegNo := params[0]
	Make := params[1]
	Model := params[2]
	ModelVariant := params[3]
	ModelYear := params[4]
	Color := params[5]
	EngineNo := params[6]
	ChassisNo := params[7]
	Owner := params[8]
	var Insurance []string
	var PollutionCert []string
	var OwnershipHistory []ownership

	// Check if Vehicle exists with Key => params[0]
	vehicleAsBytes, err := stub.GetState(RegNo)
	if err != nil {
		return shim.Error("Failed to check if Vehicle exists!")
	} else if vehicleAsBytes != nil {
		return shim.Error("Vehicle Already Exists!")
	}

	// Generate Vehicle from params provided
	vehicle := &vehicle{"VEHCL",
		RegNo, Make, Model, ModelVariant, ModelYear, Color,
		EngineNo, ChassisNo, Owner, Insurance, PollutionCert, OwnershipHistory}

	// Get JSON bytes of Vehicle struct
	vehicleJSONasBytes, err := json.Marshal(vehicle)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated Vehicle with Key => params[0]
	err = stub.PutState(RegNo, vehicleJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Add Vehicle ID to The Citizen with UID
	args := util.ToChaincodeArgs("addVehicle", Owner, RegNo)
	response := stub.InvokeChaincode("profilemanager_cc", args, "mainchannel")
	if response.Status != shim.OK {
		return shim.Error(response.Message)
	}

	// Returned on successful execution of the function
	return shim.Success(nil)
}

// Function to read an Vehicle
func (cc *Chaincode) getVehicleProfile(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check if sufficient Params passed
	if len(params) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	// Check if Params are non-empty
	if len(params[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}

	// Get State of Vehicle with Key => params[0]
	vehicleAsBytes, err := stub.GetState(params[0])
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + params[0] + "\"}"
		return shim.Error(jsonResp)
	} else if vehicleAsBytes == nil {
		jsonResp := "{\"Error\":\"Vehicle does not exist!\"}"
		return shim.Error(jsonResp)
	}

	// Returned on successful execution of the function
	return shim.Success(vehicleAsBytes)
}

// Function to add new Pollution Certificate
func (cc *Chaincode) addPollutionCertificate(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, creator, err := getTxCreatorInfo(stub)
	if !authenticatePollution(creatorOrg, creatorCertIssuer) {
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
	RegNo := params[0]
	Cert := params[1]

	// Check if Vehicle exists with Key => RegNo
	vehicleAsBytes, err := stub.GetState(RegNo)
	if err != nil {
		return shim.Error("Failed to get Vehicle Details!")
	} else if vehicleAsBytes == nil {
		return shim.Error("Error: Vehicle Does NOT Exist!")
	}

	// Create Update struct var
	vehicleToUpdate := vehicle{}
	err = json.Unmarshal(vehicleAsBytes, &vehicleToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	// Update Vehicle.PollutionCert to append => Cert
	vehicleToUpdate.PollutionCert = append(vehicleToUpdate.PollutionCert, Cert)

	// Get JSON bytes of Vehicle struct
	vehicleJSONasBytes, err := json.Marshal(vehicleToUpdate)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated Vehicle with Key => params[0]
	err = stub.PutState(RegNo, vehicleJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Returned on successful execution of the function
	return shim.Success(nil)
}

// Function to add new Insurance Policy
func (cc *Chaincode) addInsurancePolicy(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, creator, err := getTxCreatorInfo(stub)
	if !authenticateInsurance(creatorOrg, creatorCertIssuer) {
		return shim.Error("{\"Error\":\"Access Denied!\",\"Payload\":{\"MSP\":\"" + creatorOrg + "\",\"CA\":\"" + creatorCertIssuer + "\"}}")
	}

	// Check if sufficient Params passed
	if len(params) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 9")
	}

	// Check if Params are non-empty
	for a := 0; a < 2; a++ {
		if len(params[a]) <= 0 {
			return shim.Error("Argument must be a non-empty string")
		}
	}

	// Copy the Values from params[]
	RegNo := params[0]
	Insurance := params[1]

	// Check if Vehicle exists with Key => RegNo
	vehicleAsBytes, err := stub.GetState(RegNo)
	if err != nil {
		return shim.Error("Failed to get Vehicle Details!")
	} else if vehicleAsBytes == nil {
		return shim.Error("Error: Vehicle Does NOT Exist!")
	}

	// Create Update struct var
	vehicleToUpdate := vehicle{}
	err = json.Unmarshal(vehicleAsBytes, &vehicleToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	// Update Vehicle.PollutionCert to append => Cert
	vehicleToUpdate.Insurance = append(vehicleToUpdate.Insurance, Insurance)

	// Get JSON bytes of Vehicle struct
	vehicleJSONasBytes, err := json.Marshal(vehicleToUpdate)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated Vehicle with Key => params[0]
	err = stub.PutState(RegNo, vehicleJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Returned on successful execution of the function
	return shim.Success(nil)
}

// Function to add Transfer Ownership
func (cc *Chaincode) transferOwnership(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, creator, err := getTxCreatorInfo(stub)
	if !authenticateRTO(creatorOrg, creatorCertIssuer) {
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
	RegNo := params[0]
	Curr := params[1]
	Date := params[1]
	DateI, err := strconv.Atoi(Date)
	if err != nil {
		return shim.Error("Error: Invalid Date!")
	}

	// Check if Vehicle exists with Key => RegNo
	vehicleAsBytes, err := stub.GetState(RegNo)
	if err != nil {
		return shim.Error("Failed to get Vehicle Details!")
	} else if vehicleAsBytes == nil {
		return shim.Error("Error: Vehicle Does NOT Exist!")
	}

	// Create Update struct var
	vehicleToUpdate := vehicle{}
	err = json.Unmarshal(vehicleAsBytes, &vehicleToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	Record := ownership{vehicleToUpdate.Owner, Curr, creator, DateI}

	// Update Vehicle.OwnershipHistory to append => Record
	vehicleToUpdate.OwnershipHistory = append(vehicleToUpdate.OwnershipHistory, Record)
	// Update Vehicle.Owner = Curr
	vehicleToUpdate.Owner = Curr

	// Get JSON bytes of Vehicle struct
	vehicleJSONasBytes, err := json.Marshal(vehicleToUpdate)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated Vehicle with Key => params[0]
	err = stub.PutState(RegNo, vehicleJSONasBytes)
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

// Authenticate => RTO
func authenticateRTO(mspID string, certCN string) bool {
	return (mspID == "RTOMSP") && (certCN == "ca.rto.vehicle.com")
}

// Authenticate => Insurance
func authenticateInsurance(mspID string, certCN string) bool {
	return (mspID == "InsuranceMSP") && (certCN == "ca.insurance.vehicle.com")
}

// Authenticate => Pollution
func authenticatePollution(mspID string, certCN string) bool {
	return (mspID == "PollutionMSP") && (certCN == "ca.pollution.vehicle.com")
}
