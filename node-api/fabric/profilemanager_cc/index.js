const GetProfile = require("./getProfile");
const RegisterCitizen = require("./registerCitizen");
const AddVehicle = require("./addVehicle");
const AddLicense = require("./addLicense");

const payload = {
    GetProfile,
    RegisterCitizen,
    AddVehicle,
    AddLicense,
};

module.exports = payload;
