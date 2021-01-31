const AddInsurancePolicy = require("./addInsurancePolicy");
const ClaimInsurancePolicy = require("./claimInsurancePolicy");
const ReadInsurancePolicyClaim = require("./readInsurancePolicyClaim");
const UpdateInsuranceClaim = require("./updateInsuranceClaim");
const ApproveInsuranceClaim = require("./approveInsuranceClaim");

const payload = {
    AddInsurancePolicy,
    ClaimInsurancePolicy,
    ReadInsurancePolicyClaim,
    UpdateInsuranceClaim,
    ApproveInsuranceClaim,
};

module.exports = payload;
