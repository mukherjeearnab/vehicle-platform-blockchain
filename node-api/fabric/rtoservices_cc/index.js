const NewDLApplication = require("./newDLApplication");
const NewVRApplication = require("./newVRApplication");
const GetDLVRApplication = require("./getDLVRApplication");
const UpdateDLApplication = require("./updateDLApplication");
const UpdateVRApplication = require("./updateVRApplication");
const CreateDL = require("./createDL");
const GetDL = require("./getDL");

const payload = {
    NewDLApplication,
    NewVRApplication,
    GetDLVRApplication,
    UpdateDLApplication,
    UpdateVRApplication,
    CreateDL,
    GetDL,
};

module.exports = payload;
