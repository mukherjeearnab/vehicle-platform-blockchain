const mongoose = require("mongoose");

const connectionString = "mongodb+srv://user:pass@reelitin.5jxp1.mongodb.net/vehblock";

mongoose.connect(connectionString, {
    useNewUrlParser: true,
    useCreateIndex: true,
    useFindAndModify: false,
    useUnifiedTopology: true,
});
