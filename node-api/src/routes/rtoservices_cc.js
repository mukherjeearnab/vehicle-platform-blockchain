const express = require("express");
const md5 = require("md5");
const JWTmiddleware = require("../helpers/jwtVerifyMiddleware");
const RTOServices = require("../../fabric/rtoservices_cc");

const router = new express.Router();

router.get("/api/main/rto/getDLVRA/:id", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    const ID = req.params.id;
    try {
        let data = await RTOServices.GetDLVRApplication(req.user, ID);
        res.status(200).send(data);
    } catch (error) {
        console.log(error);
        res.status(404).send({ message: "Application NOT found!" });
    }
});

router.post("/api/main/rto/newDLA", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    try {
        RTOServicesData = JSON.parse(req.body.payload);
        RTOServicesData.DateTime = Math.floor(new Date() / 1000).toString();
        RTOServicesData.ApplicationID = md5(JSON.stringify(RTOServicesData) + new Date().toString());
        await RTOServices.NewDLApplication(req.user, RTOServicesData);
        res.status(200).send({
            message: "DL Application has been successfully added!",
            id: RTOServicesData.ApplicationID,
        });
    } catch (error) {
        console.log(error);
        res.status(500).send({ message: "Error! DL Application NOT Added!" });
    }
});

router.post("/api/main/rto/newVRA", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    try {
        RTOServicesData = JSON.parse(req.body.payload);
        RTOServicesData.DateTime = Math.floor(new Date() / 1000).toString();
        RTOServicesData.ApplicationID = md5(JSON.stringify(RTOServicesData) + new Date().toString());
        await RTOServices.NewVRApplication(req.user, RTOServicesData);
        res.status(200).send({
            message: "VR Application has been successfully added!",
            id: RTOServicesData.ApplicationID,
        });
    } catch (error) {
        console.log(error);
        res.status(500).send({ message: "Error! VR Application NOT Added!" });
    }
});

router.post("/api/main/rto/updateDLA/:id", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    const ID = req.params.id;
    try {
        RTOServicesData = JSON.parse(req.body.payload);
        RTOServicesData.ApplicationID = ID;
        RTOServicesData.Date = Math.floor(new Date() / 1000).toString();
        await RTOServices.UpdateDLApplication(req.user, RTOServicesData);
        res.status(200).send({
            message: "DL Application has been successfully Updated!",
            id: RTOServicesData.ApplicationID,
        });
    } catch (error) {
        console.log(error);
        res.status(500).send({ message: "Error! DL Application NOT Updated!" });
    }
});

router.post("/api/main/rto/updateVRA/:id", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    const ID = req.params.id;
    try {
        RTOServicesData = JSON.parse(req.body.payload);
        RTOServicesData.ApplicationID = ID;
        RTOServicesData.Date = Math.floor(new Date() / 1000).toString();
        await RTOServices.UpdateVRApplication(req.user, RTOServicesData);
        res.status(200).send({
            message: "VR Application has been successfully Updated!",
            id: RTOServicesData.ApplicationID,
        });
    } catch (error) {
        console.log(error);
        res.status(500).send({ message: "Error! VR Application NOT Updated!" });
    }
});

router.post("/api/main/rto/createDL", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    try {
        RTOServicesData = JSON.parse(req.body.payload);
        await RTOServices.CreateDL(req.user, RTOServicesData);
        res.status(200).send({
            message: "DL has been successfully Created!",
            id: RTOServicesData.LicenseNumber,
        });
    } catch (error) {
        console.log(error);
        res.status(500).send({ message: "Error! License NOT Created!" });
    }
});

router.get("/api/main/rto/getDL/:id", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    const ID = req.params.id;
    try {
        let data = await RTOServices.GetDL(req.user, ID);
        res.status(200).send(data);
    } catch (error) {
        console.log(error);
        res.status(404).send({ message: "DL NOT found!" });
    }
});

module.exports = router;
