const mongoose = require("mongoose");
const userJSON = require("./json/user-model.json");
require("../db/mongoose");

const userInfo = new mongoose.Schema(userJSON);

const user = mongoose.model("userInfo_vehiclePlatform", userInfo, "userinfo_vehicleplatform");

module.exports = user;
