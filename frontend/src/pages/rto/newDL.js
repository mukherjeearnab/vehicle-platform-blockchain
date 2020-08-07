import React, { Component } from "react";
//import { Redirect } from "react-router-dom";

import { TextField, Button, CircularProgress } from "@material-ui/core";

class App extends Component {
    state = {
        vra: {},
        message: "",
    };

    onAddApp = async () => {
        console.log(this.state.vra);

        const requestOptions = {
            method: "POST",
            headers: { "Content-Type": "application/json", "x-access-token": localStorage.getItem("session") },
            body: JSON.stringify({
                payload: this.state.vra,
            }),
        };

        this.setState({
            message: (
                <span>
                    <CircularProgress />
                    <br></br> Loading.....
                </span>
            ),
        });

        let response = await fetch("http://192.168.1.30:3000/api/main/rto/createDL", requestOptions);
        let res = await response.json();
        console.log(res);
        this.setState({ message: "License Number: " + res.id + "\n" });
    };

    render() {
        return (
            <div>
                <h2>New Driver License Issue</h2>
                {this.state.message}
                <TextField
                    className="inputs"
                    label="License Number"
                    variant="outlined"
                    value={this.state.vra.LicenseNumber}
                    onChange={(event) => {
                        let vra = this.state.vra;
                        vra.LicenseNumber = event.target.value;
                        this.setState({
                            vra,
                        });
                    }}
                />
                <br />
                <br />
                <TextField
                    className="inputs"
                    label="UID"
                    variant="outlined"
                    value={this.state.vra.UID}
                    onChange={(event) => {
                        let vra = this.state.vra;
                        vra.UID = event.target.value;
                        this.setState({
                            vra,
                        });
                    }}
                />
                <br />
                <br />
                <TextField
                    className="inputs"
                    label="Vehicle Type"
                    variant="outlined"
                    value={this.state.vra.VehicleType}
                    onChange={(event) => {
                        let vra = this.state.vra;
                        vra.VehicleType = event.target.value;
                        this.setState({
                            vra,
                        });
                    }}
                />
                <br />
                <br />
                <TextField
                    className="inputs"
                    label="Expiry Date"
                    variant="outlined"
                    value={this.state.vra.ExpiryDate}
                    onChange={(event) => {
                        let vra = this.state.vra;
                        vra.ExpiryDate = event.target.value;
                        this.setState({
                            vra,
                        });
                    }}
                />
                <br />
                <br />
                <br />
                <br />
                <Button onClick={this.onAddApp} variant="contained" color="primary">
                    Create Driving License
                </Button>
            </div>
        );
    }
}

export default App;
