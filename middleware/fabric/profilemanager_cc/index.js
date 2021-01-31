const GetProfile = require("./getProfile");
const RegisterCitizen = require("./registerCitizen");
const AddVehicle = require("./addVehicle");
const AddLicense = require("./addLicense");
const RegisterRto = require("./registerRto");
const RegisterRtoE = require("./registerRtoE");
const RegisterInsuranceCompany = require("./registerInsuranceCompany");
const RegisterInsuranceE = require("./registerInsuranceE");
const RegisterPollutionCompany = require("./registerPollutionCompany");
const RegisterPollutionE = require("./registerPollutionE");
const RegisterPoliceStation = require("./registerPoliceStation");
const RegisterPoliceOfficer = require("./registerPoliceOfficer");

const payload = {
    GetProfile,
    RegisterCitizen,
    AddVehicle,
    AddLicense,
    RegisterRto,
    RegisterRtoE,
    RegisterInsuranceCompany,
    RegisterInsuranceE,
    RegisterPollutionCompany,
    RegisterPollutionE,
    RegisterPoliceStation,
    RegisterPoliceOfficer,
};

module.exports = payload;
