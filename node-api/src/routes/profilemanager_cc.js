const express = require("express");
const md5 = require("md5");
const JWTmiddleware = require("../helpers/jwtVerifyMiddleware");
const ProfileManager = require("../../fabric/profilemanager_cc");

const router = new express.Router();

router.get("/api/main/profile/get/:id", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    const ID = req.params.id;
    try {
        let data = await ProfileManager.GetProfile(req.user, ID);
        res.status(200).send(data);
    } catch (error) {
        console.log(error);
        res.status(404).send({ message: "Profile NOT found!" });
    }
});

router.post("/api/main/profile/regCi", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    try {
        profileData = JSON.parse(req.body.payload);
        await ProfileManager.RegisterCitizen(req.user, profileData);
        res.status(200).send({
            message: "Citizen Profile has been successfully added!",
            id: profileData.UID,
        });
    } catch (error) {
        console.log(error);
        res.status(500).send({ message: "Error! Citizen Profile NOT Added!" });
    }
});

router.post("/api/main/profile/addVeh", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    try {
        const ID = req.params.id;
        profileData = JSON.parse(req.body.payload);
        await ProfileManager.AddVehicle(req.user, profileData);
        res.status(200).send({ message: `Vehicle has been Successfully Added to the Profile ${profileData.UID}.` });
    } catch (error) {
        console.log(error);
        res.status(500).send({ message: `Vehicle NOT Added to the Profile ${profileData.UID}.` });
    }
});

router.post("/api/main/profile/addLic", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    try {
        const ID = req.params.id;
        profileData = JSON.parse(req.body.payload);
        await ProfileManager.AddLicense(req.user, profileData);
        res.status(200).send({ message: `License has been Successfully Added to the Profile ${profileData.UID}.` });
    } catch (error) {
        console.log(error);
        res.status(500).send({ message: `License NOT Added to the Profile ${profileData.UID}.` });
    }
});

//TODO: Add remaining Registering Routes

router.post("/api/main/profile/regRto", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    try {
        profileData = JSON.parse(req.body.payload);
        await ProfileManager.RegisterRto(req.user, profileData);
        res.status(200).send({
            message: "RTO Profile has been successfully added!",
            id: profileData.RtoID,
        });
    } catch (error) {
        console.log(error);
        res.status(500).send({ message: "Error! RTO Profile NOT Added!" });
    }
});

router.post("/api/main/profile/regRtoE", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    try {
        profileData = JSON.parse(req.body.payload);
        await ProfileManager.RegisterRtoE(req.user, profileData);
        res.status(200).send({
            message: "RTO Employee Profile has been successfully added!",
            id: profileData.EmployeeID,
        });
    } catch (error) {
        console.log(error);
        res.status(500).send({ message: "Error! RTO Employee Profile NOT Added!" });
    }
});

router.post("/api/main/profile/regIns", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    try {
        profileData = JSON.parse(req.body.payload);
        await ProfileManager.RegisterInsuranceCompany(req.user, profileData);
        res.status(200).send({
            message: "IC Profile has been successfully added!",
            id: profileData.CompanyID,
        });
    } catch (error) {
        console.log(error);
        res.status(500).send({ message: "Error! IC Profile NOT Added!" });
    }
});

router.post("/api/main/profile/regInsE", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    try {
        profileData = JSON.parse(req.body.payload);
        await ProfileManager.RegisterInsuranceE(req.user, profileData);
        res.status(200).send({
            message: "Insurance Employee Profile has been successfully added!",
            id: profileData.EmployeeID,
        });
    } catch (error) {
        console.log(error);
        res.status(500).send({ message: "Error! Insurance Employee Profile NOT Added!" });
    }
});

router.post("/api/main/profile/regPol", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    try {
        profileData = JSON.parse(req.body.payload);
        await ProfileManager.RegisterPollutionCompany(req.user, profileData);
        res.status(200).send({
            message: "PC Profile has been successfully added!",
            id: profileData.CompanyID,
        });
    } catch (error) {
        console.log(error);
        res.status(500).send({ message: "Error! PC Profile NOT Added!" });
    }
});

router.post("/api/main/profile/regPolE", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    try {
        profileData = JSON.parse(req.body.payload);
        await ProfileManager.RegisterPollutionE(req.user, profileData);
        res.status(200).send({
            message: "Pollution Employee Profile has been successfully added!",
            id: profileData.EmployeeID,
        });
    } catch (error) {
        console.log(error);
        res.status(500).send({ message: "Error! Pollution Employee Profile NOT Added!" });
    }
});

router.post("/api/main/profile/regPlc", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    try {
        profileData = JSON.parse(req.body.payload);
        await ProfileManager.RegisterPoliceStation(req.user, profileData);
        res.status(200).send({
            message: "Police Station Profile has been successfully added!",
            id: profileData.StationID,
        });
    } catch (error) {
        console.log(error);
        res.status(500).send({ message: "Error! Police Station Profile NOT Added!" });
    }
});

router.post("/api/main/profile/regPlcO", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    try {
        profileData = JSON.parse(req.body.payload);
        await ProfileManager.RegisterPoliceOfficer(req.user, profileData);
        res.status(200).send({
            message: "Police Officer Profile has been successfully added!",
            id: profileData.OfficerID,
        });
    } catch (error) {
        console.log(error);
        res.status(500).send({ message: "Error! Police Officer Profile NOT Added!" });
    }
});

module.exports = router;
