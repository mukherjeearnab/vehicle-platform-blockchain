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

router.post("/api/main/judgement/add", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    try {
        judgementData = JSON.parse(req.body.payload);
        judgementData.ID = md5(JSON.stringify(judgementData) + new Date().toString());
        await ProfileManager.AddProfileManager(req.user, judgementData);
        res.status(200).send({
            message: "ProfileManager has been successfully added!",
            id: judgementData.ID,
        });
    } catch (error) {
        console.log(error);
        res.status(500).send({ message: "Error! ProfileManager NOT Added!" });
    }
});

router.post("/api/main/judgement/addevidence/:id", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    try {
        const ID = req.params.id;
        judgementData = JSON.parse(req.body.payload);
        judgementData.ID = ID;
        await ProfileManager.AddEvidence(req.user, judgementData);
        res.status(200).send({ message: "Evidence has been Successfully Added to the ProfileManager Report!" });
    } catch (error) {
        console.log(error);
        res.status(404).send({ message: "ProfileManager NOT found!" });
    }
});

router.post("/api/main/judgement/addsentence/:id", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    try {
        const ID = req.params.id;
        judgementData = JSON.parse(req.body.payload);
        judgementData.ID = ID;
        await ProfileManager.AddSentence(req.user, judgementData);
        res.status(200).send({ message: "Sentence has been Successfully Added to the ProfileManager Report!" });
    } catch (error) {
        console.log(error);
        res.status(404).send({ message: "ProfileManager NOT found!" });
    }
});

router.post("/api/main/judgement/setcomplete/:id", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    try {
        const ID = req.params.id;
        await ProfileManager.SetComplete(req.user, ID);
        res.status(200).send({ message: "ProfileManager has been Successfully Set to Complete!" });
    } catch (error) {
        console.log(error);
        res.status(404).send({ message: "ProfileManager NOT found!" });
    }
});

module.exports = router;
