const GetProfile = require("./getProfile");
const AddChargeSheet = require("./addChargeSheet");
const AddAccusedPerson = require("./addAccusedPerson");
const AddBriefReport = require("./addBriefReport");
const AddChargedPerson = require("./addChargedPerson");
const AddFIRID = require("./addFIRID");
const AddInvestigatingOfficer = require("./addInvestigatingOfficer");
const AddInvestigationID = require("./addInvestigationID");
const AddSectionOfLaw = require("./addSectionOfLaw");

const payload = {
    GetProfile,
    AddChargeSheet,
    AddAccusedPerson,
    AddBriefReport,
    AddChargedPerson,
    AddFIRID,
    AddInvestigatingOfficer,
    AddInvestigationID,
    AddSectionOfLaw,
};

module.exports = payload;
