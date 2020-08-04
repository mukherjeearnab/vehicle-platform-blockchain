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

module.exports = router;
