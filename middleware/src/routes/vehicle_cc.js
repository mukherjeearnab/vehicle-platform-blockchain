const express = require("express");
const md5 = require("md5");
const JWTmiddleware = require("../helpers/jwtVerifyMiddleware");
const Vehicle = require("../../fabric/vehicle_cc");

const router = new express.Router();

router.get("/api/main/vehicle/get/:id", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    const ID = req.params.id;
    try {
        let data = await Vehicle.GetVehicleProfile(req.user, ID);
        res.status(200).send(data);
    } catch (error) {
        console.log(error);
        res.status(404).send({ message: "Vehicle NOT found!" });
    }
});

router.post("/api/main/vehicle/create", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    try {
        VehicleData = req.body.payload;
        await Vehicle.CreateVehicleProfile(req.user, VehicleData);
        res.status(200).send({
            message: "Vehicle has been successfully Created!",
            id: VehicleData.RegNo,
        });
    } catch (error) {
        console.log(error);
        res.status(500).send({ message: "Error! Vehicle NOT Created!" });
    }
});

router.post("/api/main/vehicle/transfer/:id", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    const ID = req.params.id;
    try {
        VehicleData = req.body.payload;
        VehicleData.RegNo = ID;
        VehicleData.Date = Math.floor(new Date() / 1000).toString();
        await Vehicle.TransferOwnership(req.user, VehicleData);
        res.status(200).send({
            message: "Vehicle Ownership transfer successful!",
            id: VehicleData.RegNo,
            curr: VehicleData.Curr,
        });
    } catch (error) {
        console.log(error);
        res.status(500).send({ message: "Error! Vehicle ownership NOT transfered!" });
    }
});

module.exports = router;
