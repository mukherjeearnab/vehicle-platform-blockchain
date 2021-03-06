const express = require("express");
const md5 = require("md5");
const JWTmiddleware = require("../helpers/jwtVerifyMiddleware");
const Insurance = require("../../fabric/insurance_cc");

const router = new express.Router();

router.get("/api/main/insurance/get/:id", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    const ID = req.params.id;
    try {
        let data = await Insurance.ReadInsurancePolicyClaim(req.user, ID);
        res.status(200).send(data);
    } catch (error) {
        console.log(error);
        res.status(404).send({ message: "Insurance Policy Claim NOT found!" });
    }
});

router.post("/api/main/insurance/policy/add", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    try {
        InsuranceData = req.body.payload;
        InsuranceData.Date = Math.floor(new Date() / 1000).toString();
        InsuranceData.InsuranceID = md5(JSON.stringify(InsuranceData) + new Date().toString());
        await Insurance.AddInsurancePolicy(req.user, InsuranceData);
        res.status(200).send({
            message: "Insurance Policy has been successfully added!",
            id: InsuranceData.InsuranceID,
        });
    } catch (error) {
        console.log(error);
        res.status(500).send({ message: "Error! Insurance Policy NOT Added!" });
    }
});

router.post("/api/main/insurance/claim/add", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    try {
        InsuranceData = req.body.payload;
        InsuranceData.DateTime = Math.floor(new Date() / 1000).toString();
        InsuranceData.ClaimID = md5(JSON.stringify(InsuranceData) + new Date().toString());
        await Insurance.ClaimInsurancePolicy(req.user, InsuranceData);
        res.status(200).send({
            message: "Insurance Claim has been successfully added!",
            id: InsuranceData.ClaimID,
        });
    } catch (error) {
        console.log(error);
        res.status(500).send({ message: "Error! Insurance Claim NOT Added!" });
    }
});

router.post("/api/main/insurance/claim/update/:id", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    const ID = req.params.id;
    try {
        InsuranceData = req.body.payload;
        InsuranceData.ClaimID = ID;
        InsuranceData.Date = Math.floor(new Date() / 1000).toString();
        await Insurance.UpdateInsuranceClaim(req.user, InsuranceData);
        res.status(200).send({
            message: "Insurance Claim has been successfully Updated!",
            id: InsuranceData.ClaimID,
        });
    } catch (error) {
        console.log(error);
        res.status(500).send({ message: "Error! Insurance Claim NOT Updated!" });
    }
});

router.post("/api/main/insurance/claim/approve/:id", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    const ID = req.params.id;
    try {
        InsuranceData = req.body.payload;
        InsuranceData.ClaimID = ID;
        InsuranceData.Date = Math.floor(new Date() / 1000).toString();
        await Insurance.ApproveInsuranceClaim(req.user, InsuranceData);
        res.status(200).send({
            message: "Insurance Claim has been successfully Approved!",
            id: InsuranceData.ClaimID,
        });
    } catch (error) {
        console.log(error);
        res.status(500).send({ message: "Error! Insurance Claim NOT Approved!" });
    }
});

module.exports = router;
