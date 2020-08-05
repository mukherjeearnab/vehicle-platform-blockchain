const ipfsAPI = require("ipfs-api");
const express = require("express");
const path = require("path");
const multer = require("multer");
const fs = require("fs");
const JWTmiddleware = require("../helpers/jwtVerifyMiddleware");
const TrafficViolation = require("../../fabric/trafficviolation_cc");

const router = new express.Router();
const ipfs = ipfsAPI("ipfs.infura.io", "5001", { protocol: "https" });
const uploadPath = path.join(process.cwd(), "uploads");
var upload = multer({ dest: uploadPath });

router.get("/api/main/trafficviolation/read/:id", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    const ID = req.params.id;
    try {
        let data = await TrafficViolation.GetTV(req.user, ID);
        res.status(200).send(data);
    } catch (error) {
        console.log(error);
        res.status(404).send({ message: "Traffic Violation NOT found!" });
    }
});

router.post("/api/main/trafficviolation/add", upload.single("file"), JWTmiddleware, (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    try {
        const oldname = "uploads/" + req.file.filename;
        const newname = "uploads/" + req.file.filename + "." + req.file.originalname.split(".").pop();
        fs.renameSync(oldname, newname, console.log);

        let file = fs.readFileSync(newname);
        let fileBuffer = new Buffer(file);

        ipfs.files.add(fileBuffer, (err, file) => {
            if (err) {
                console.log(err);
            }
            trafficviolationData = JSON.parse(req.body.payload);
            trafficviolationData.Evidence = file[0].path;
            trafficviolationData.DateFiling = Math.floor(new Date() / 1000).toString();
            trafficviolationData.TVID = md5(JSON.stringify(trafficviolationData) + new Date().toString());
            TrafficViolation.AddTrafficViolation(req.user, trafficviolationData).then(() => {
                fs.unlinkSync(newname);
                res.status(200).send({
                    message: "Traffic Violation has been successfully filed!",
                    data: trafficviolationData,
                });
            });
        });
    } catch (error) {
        console.log(error);
        res.status(500).send({ message: "Traffic Violation NOT Added!" });
    }
});

module.exports = router;
