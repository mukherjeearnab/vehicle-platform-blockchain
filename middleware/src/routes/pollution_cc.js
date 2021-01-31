const express = require("express");
const md5 = require("md5");
const JWTmiddleware = require("../helpers/jwtVerifyMiddleware");
const Pollution = require("../../fabric/pollution_cc");

const router = new express.Router();

router.get("/api/main/pollution/get/:id", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    const ID = req.params.id;
    try {
        let data = await Pollution.GetPUCC(req.user, ID);
        res.status(200).send(data);
    } catch (error) {
        console.log(error);
        res.status(404).send({ message: "PUCC NOT found!" });
    }
});

router.post("/api/main/pollution/create", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    try {
        PUCCData = req.body.payload;
        PUCCData.DateTime = Math.floor(new Date() / 1000).toString();
        PUCCData.ID = md5(JSON.stringify(PUCCData) + new Date().toString());
        await Pollution.CreatePUCC(req.user, PUCCData);
        res.status(200).send({
            message: "PUCC has been successfully added!",
            id: PUCCData.ID,
        });
    } catch (error) {
        console.log(error);
        res.status(500).send({ message: "Error! PUCC NOT Added!" });
    }
});

module.exports = router;
