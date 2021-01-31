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

// Definition of the Citizen structure
type citizen struct {
	Type           string   `json:"Type"`
	UID            string   `json:"UID"`
	Name           string   `json:"Name"`
	DOB            int      `json:"DOB"`
	Address        string   `json:"Address"`
	Phone          string   `json:"Phone"`
	LicenseNumbers []string `json:"LicenseNumbers"`
	Vehicles       []string `json:"Vehicles"`
}

// Definition of the RTO structure
type rto struct {
	Type       string `json:"Type"`
	RtoID      string `json:"RtoID"`
	RtoName    string `json:"RtoName"`
	RtoAddress string `json:"RtoAddress"`
}

// Definition of the RTOE structure
type rtoe struct {
	Type       string `json:"Type"`
	EmployeeID string `json:"EmployeeID"`
	UID        string `json:"UID"`
	RtoID      string `json:"RtoID"`
}

// Definition of the Insurance Company structure
type insurance struct {
	Type           string `json:"Type"`
	CompanyID      string `json:"CompanyID"`
	CompanyName    string `json:"CompanyName"`
	CompanyAddress string `json:"CompanyAddress"`
}

// Definition of the Insurance Employee structure
type insuranceE struct {
	Type       string `json:"Type"`
	EmployeeID string `json:"EmployeeID"`
	UID        string `json:"UID"`
	CompanyID  string `json:"CompanyID"`
}

// Definition of the Pollution Company structure
type pollution struct {
	Type           string `json:"Type"`
	CompanyID      string `json:"CompanyID"`
	CompanyName    string `json:"CompanyName"`
	CompanyAddress string `json:"CompanyAddress"`
}

// Definition of the Pollution Employee structure
type pollutionE struct {
	Type       string `json:"Type"`
	EmployeeID string `json:"EmployeeID"`
	UID        string `json:"UID"`
	CompanyID  string `json:"CompanyID"`
}

// Definition of the Police Station structure
type policeS struct {
	Type           string `json:"Type"`
	CompanyID      string `json:"CompanyID"`
	CompanyName    string `json:"CompanyName"`
	CompanyAddress string `json:"CompanyAddress"`
}

// Definition of the Police Officer structure
type policeO struct {
	Type       string `json:"Type"`
	EmployeeID string `json:"EmployeeID"`
	UID        string `json:"UID"`
	CompanyID  string `json:"CompanyID"`
}

// Init is called when the chaincode is instantiated by the blockchain network.
func (cc *Chaincode) Init(stub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

// Invoke is called as a result of an application request to run the chaincode.
func (cc *Chaincode) Invoke(stub shim.ChaincodeStubInterface) sc.Response {
	fcn, params := stub.GetFunctionAndParameters()
	fmt.Println("Invoke()", fcn, params)

	if fcn == "registerCitizen" {
		return cc.registerCitizen(stub, params)
	} else if fcn == "addVehicle" {
		return cc.addVehicle(stub, params)
	} else if fcn == "addLicense" {
		return cc.addLicense(stub, params)
	} else if fcn == "registerRto" {
		return cc.registerRto(stub, params)
	} else if fcn == "registerRtoE" {
		return cc.registerRtoE(stub, params)
	} else if fcn == "registerInsuranceCompany" {
		return cc.registerInsuranceCompany(stub, params)
	} else if fcn == "registerInsuranceE" {
		return cc.registerInsuranceE(stub, params)
	} else if fcn == "registerPollutionCompany" {
		return cc.registerPollutionCompany(stub, params)
	} else if fcn == "registerPollutionE" {
		return cc.registerPollutionE(stub, params)
	} else if fcn == "registerPoliceStation" {
		return cc.registerPoliceStation(stub, params)
	} else if fcn == "registerPoliceOfficer" {
		return cc.registerPoliceOfficer(stub, params)
	} else if fcn == "getProfile" {
		return cc.getProfile(stub, params)
	} else {
		fmt.Println("Invoke() did not find func: " + fcn)
		return shim.Error("Received unknown function invocation!")
	}
}

// Function to Register new Citizen
func (cc *Chaincode) registerCitizen(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, err := getTxCreatorInfo(stub)
	if !authenticateCitizen(creatorOrg, creatorCertIssuer) {
		return shim.Error("{\"Error\":\"Access Denied!\",\"Payload\":{\"MSP\":\"" + creatorOrg + "\",\"CA\":\"" + creatorCertIssuer + "\"}}")
	}

	// Check if sufficient Params passed
	if len(params) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	// Check if Params are non-empty
	for a := 0; a < 5; a++ {
		if len(params[a]) <= 0 {
			return shim.Error("Argument must be a non-empty string")
		}
	}

	UID := params[0]
	Name := params[1]
	DOB := params[2]
	Address := params[3]
	Phone := params[4]
	DOBI, err := strconv.Atoi(DOB)
	var LicenseNumbers []string
	var Vehicles []string
	if err != nil {
		return shim.Error("Error: Invalid DateTime!")
	}

	// Check if Citizen exists with Key => params[0]
	citizenAsBytes, err := stub.GetState(UID)
	if err != nil {
		return shim.Error("Failed to check if Citizen exists!")
	} else if citizenAsBytes != nil {
		return shim.Error("Citizen Already Exists!")
	}

	// Generate Citizen from params provided
	citizen := &citizen{"CTZN",
		UID, Name, DOBI, Address, Phone, LicenseNumbers, Vehicles}

	// Get JSON bytes of Citizen struct
	citizenJSONasBytes, err := json.Marshal(citizen)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated Citizen with Key => params[0]
	err = stub.PutState(UID, citizenJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Returned on successful execution of the function
	return shim.Success(nil)
}

// Function to add new Vehicle to Citizen Profile
func (cc *Chaincode) addVehicle(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, err := getTxCreatorInfo(stub)
	if !authenticateRTO(creatorOrg, creatorCertIssuer) {
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
	UID := params[0]
	RegNo := params[1]

	// Check if Citizen exists with Key => UID
	citizenAsBytes, err := stub.GetState(UID)
	if err != nil {
		return shim.Error("Failed to get Citizen Details!")
	} else if citizenAsBytes == nil {
		return shim.Error("Error: Citizen Does NOT Exist!")
	}

	// Create Update struct var
	citizenToUpdate := citizen{}
	err = json.Unmarshal(citizenAsBytes, &citizenToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	// Update Citizen.Vehicles to append => RegNo
	citizenToUpdate.Vehicles = append(citizenToUpdate.Vehicles, RegNo)

	// Get JSON bytes of Citizen struct
	citizenJSONasBytes, err := json.Marshal(citizenToUpdate)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated Citizen with Key => params[0]
	err = stub.PutState(UID, citizenJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Returned on successful execution of the function
	return shim.Success(nil)
}

// Function to add new License to Citizen Profile
func (cc *Chaincode) addLicense(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, err := getTxCreatorInfo(stub)
	if !authenticateRTO(creatorOrg, creatorCertIssuer) {
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
	UID := params[0]
	License := params[1]

	// Check if Citizen exists with Key => RegNo
	citizenAsBytes, err := stub.GetState(UID)
	if err != nil {
		return shim.Error("Failed to get Citizen Details!")
	} else if citizenAsBytes == nil {
		return shim.Error("Error: Citizen Does NOT Exist!")
	}

	// Create Update struct var
	citizenToUpdate := citizen{}
	err = json.Unmarshal(citizenAsBytes, &citizenToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	// Update Citizen.LicenseNumbers to append => License
	citizenToUpdate.LicenseNumbers = append(citizenToUpdate.LicenseNumbers, License)

	// Get JSON bytes of Citizen struct
	citizenJSONasBytes, err := json.Marshal(citizenToUpdate)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated Citizen with Key => params[0]
	err = stub.PutState(UID, citizenJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Returned on successful execution of the function
	return shim.Success(nil)
}

// Function to Register new RTO
func (cc *Chaincode) registerRto(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, err := getTxCreatorInfo(stub)
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

	RtoID := params[0]
	RtoName := params[1]
	RtoAddress := params[2]

	// Check if RTO exists with Key => params[0]
	rtoAsBytes, err := stub.GetState(RtoID)
	if err != nil {
		return shim.Error("Failed to check if RTO exists!")
	} else if rtoAsBytes != nil {
		return shim.Error("RTO Already Exists!")
	}

	// Generate RTO from params provided
	rto := &rto{"RTO",
		RtoID, RtoName, RtoAddress}

	// Get JSON bytes of RTO struct
	rtoJSONasBytes, err := json.Marshal(rto)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated RTO with Key => params[0]
	err = stub.PutState(RtoID, rtoJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Returned on successful execution of the function
	return shim.Success(nil)
}

// Function to Register new RTOE
func (cc *Chaincode) registerRtoE(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, err := getTxCreatorInfo(stub)
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

	EmployeeID := params[0]
	UID := params[1]
	RtoID := params[2]

	// Check if RTOE exists with Key => params[0]
	rtoeAsBytes, err := stub.GetState(EmployeeID)
	if err != nil {
		return shim.Error("Failed to check if RTOE exists!")
	} else if rtoeAsBytes != nil {
		return shim.Error("RTOE Already Exists!")
	}

	// Generate RTOE from params provided
	rtoe := &rtoe{"RTO_E",
		EmployeeID, UID, RtoID}

	// Get JSON bytes of RTOE struct
	rtoeJSONasBytes, err := json.Marshal(rtoe)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated RTOE with Key => params[0]
	err = stub.PutState(EmployeeID, rtoeJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Returned on successful execution of the function
	return shim.Success(nil)
}

// Function to Register new Insurance Company
func (cc *Chaincode) registerInsuranceCompany(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, err := getTxCreatorInfo(stub)
	if !authenticateInsurance(creatorOrg, creatorCertIssuer) {
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

	CompanyID := params[0]
	CompanyName := params[1]
	CompanyAddress := params[2]

	// Check if IC exists with Key => params[0]
	icAsBytes, err := stub.GetState(CompanyID)
	if err != nil {
		return shim.Error("Failed to check if IC exists!")
	} else if icAsBytes != nil {
		return shim.Error("IC Already Exists!")
	}

	// Generate IC from params provided
	ic := &insurance{"INSRNS",
		CompanyID, CompanyName, CompanyAddress}

	// Get JSON bytes of IC struct
	icJSONasBytes, err := json.Marshal(ic)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated IC with Key => params[0]
	err = stub.PutState(CompanyID, icJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Returned on successful execution of the function
	return shim.Success(nil)
}

// Function to Register new Insurance Employee
func (cc *Chaincode) registerInsuranceE(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, err := getTxCreatorInfo(stub)
	if !authenticateInsurance(creatorOrg, creatorCertIssuer) {
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

	EmployeeID := params[0]
	UID := params[1]
	CompanyID := params[2]

	// Check if ICE exists with Key => params[0]
	iceAsBytes, err := stub.GetState(EmployeeID)
	if err != nil {
		return shim.Error("Failed to check if ICE exists!")
	} else if iceAsBytes != nil {
		return shim.Error("ICE Already Exists!")
	}

	// Generate ICE from params provided
	ice := &insuranceE{"INSRNS_E",
		EmployeeID, UID, CompanyID}

	// Get JSON bytes of ICE struct
	iceJSONasBytes, err := json.Marshal(ice)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated ICE with Key => params[0]
	err = stub.PutState(EmployeeID, iceJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Returned on successful execution of the function
	return shim.Success(nil)
}

// Function to Register new Pollution Company
func (cc *Chaincode) registerPollutionCompany(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, err := getTxCreatorInfo(stub)
	if !authenticatePollution(creatorOrg, creatorCertIssuer) {
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

	CompanyID := params[0]
	CompanyName := params[1]
	CompanyAddress := params[2]

	// Check if PC exists with Key => params[0]
	pcAsBytes, err := stub.GetState(CompanyID)
	if err != nil {
		return shim.Error("Failed to check if PC exists!")
	} else if pcAsBytes != nil {
		return shim.Error("PC Already Exists!")
	}

	// Generate PC from params provided
	pc := &pollution{"POLUTN",
		CompanyID, CompanyName, CompanyAddress}

	// Get JSON bytes of PC struct
	pcJSONasBytes, err := json.Marshal(pc)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated PC with Key => params[0]
	err = stub.PutState(CompanyID, pcJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Returned on successful execution of the function
	return shim.Success(nil)
}

// Function to Register new Pollution Employee
func (cc *Chaincode) registerPollutionE(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, err := getTxCreatorInfo(stub)
	if !authenticatePollution(creatorOrg, creatorCertIssuer) {
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

	EmployeeID := params[0]
	UID := params[1]
	CompanyID := params[2]

	// Check if PCE exists with Key => params[0]
	pceAsBytes, err := stub.GetState(EmployeeID)
	if err != nil {
		return shim.Error("Failed to check if PCE exists!")
	} else if pceAsBytes != nil {
		return shim.Error("PCE Already Exists!")
	}

	// Generate PCE from params provided
	pce := &pollutionE{"POLUTN_E",
		EmployeeID, UID, CompanyID}

	// Get JSON bytes of PCE struct
	pceJSONasBytes, err := json.Marshal(pce)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated PCE with Key => params[0]
	err = stub.PutState(EmployeeID, pceJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Returned on successful execution of the function
	return shim.Success(nil)
}

// Function to Register new Police
func (cc *Chaincode) registerPoliceStation(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, err := getTxCreatorInfo(stub)
	if !authenticatePolice(creatorOrg, creatorCertIssuer) {
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

	StationID := params[0]
	StationName := params[1]
	StationAddress := params[2]

	// Check if PS exists with Key => params[0]
	psAsBytes, err := stub.GetState(StationID)
	if err != nil {
		return shim.Error("Failed to check if PS exists!")
	} else if psAsBytes != nil {
		return shim.Error("PS Already Exists!")
	}

	// Generate PS from params provided
	ps := &policeS{"PLCE",
		StationID, StationName, StationAddress}

	// Get JSON bytes of PS struct
	psJSONasBytes, err := json.Marshal(ps)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated PS with Key => params[0]
	err = stub.PutState(StationID, psJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Returned on successful execution of the function
	return shim.Success(nil)
}

// Function to Register new Police Officer
func (cc *Chaincode) registerPoliceOfficer(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check Access
	creatorOrg, creatorCertIssuer, err := getTxCreatorInfo(stub)
	if !authenticatePolice(creatorOrg, creatorCertIssuer) {
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

	OfficerID := params[0]
	UID := params[1]
	StationID := params[2]

	// Check if PO exists with Key => params[0]
	poAsBytes, err := stub.GetState(OfficerID)
	if err != nil {
		return shim.Error("Failed to check if PO exists!")
	} else if poAsBytes != nil {
		return shim.Error("PO Already Exists!")
	}

	// Generate PO from params provided
	po := &policeO{"PLCE_E",
		OfficerID, UID, StationID}

	// Get JSON bytes of PO struct
	poJSONasBytes, err := json.Marshal(po)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put State of newly generated PO with Key => params[0]
	err = stub.PutState(OfficerID, poJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Returned on successful execution of the function
	return shim.Success(nil)
}

// Function to read an Profile
func (cc *Chaincode) getProfile(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// Check if sufficient Params passed
	if len(params) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	// Check if Params are non-empty
	if len(params[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}

	// Get State of Profile with Key => params[0]
	profileAsBytes, err := stub.GetState(params[0])
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + params[0] + "\"}"
		return shim.Error(jsonResp)
	} else if profileAsBytes == nil {
		jsonResp := "{\"Error\":\"Evidence does not exist!\"}"
		return shim.Error(jsonResp)
	}

	// Returned on successful execution of the function
	return shim.Success(profileAsBytes)
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

// Authenticate => Citizen
func authenticateCitizen(mspID string, certCN string) bool {
	return (mspID == "CitizenMSP") && (certCN == "ca.citizen.vtan.com")
}

// Authenticate => RTO
func authenticateRTO(mspID string, certCN string) bool {
	return (mspID == "RTOMSP") && (certCN == "ca.rto.vtan.com")
}

// Authenticate => Insurance
func authenticateInsurance(mspID string, certCN string) bool {
	return (mspID == "InsuranceMSP") && (certCN == "ca.insurance.vtan.com")
}

// Authenticate => Pollution
func authenticatePollution(mspID string, certCN string) bool {
	return (mspID == "PollutionMSP") && (certCN == "ca.pollution.vtan.com")
}

// Authenticate => Police
func authenticatePolice(mspID string, certCN string) bool {
	return (mspID == "PoliceMSP") && (certCN == "ca.police.vtan.com")
}
