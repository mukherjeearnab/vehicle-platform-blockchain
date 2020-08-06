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

        let response = await fetch("http://192.168.1.30:3000/api/main/rto/newVRA", requestOptions);
        let res = await response.json();
        console.log(res);
        this.setState({ message: "VR Application ID: " + res.id + "<br/>" });
    };

    render() {
        return (
            <div>
                <h2>New DL Application</h2>
                {this.state.message}
                <TextField
                    className="inputs"
                    label="Make"
                    variant="outlined"
                    value={this.state.vra.Make}
                    onChange={(event) => {
                        let vra = this.state.vra;
                        vra.Make = event.target.value;
                        this.setState({
                            vra,
                        });
                    }}
                />
                <br />
                <br />
                <TextField
                    className="inputs"
                    label="Model"
                    variant="outlined"
                    value={this.state.vra.Model}
                    onChange={(event) => {
                        let vra = this.state.vra;
                        vra.Model = event.target.value;
                        this.setState({
                            vra,
                        });
                    }}
                />
                <br />
                <br />
                <TextField
                    className="inputs"
                    label="Variant"
                    variant="outlined"
                    value={this.state.vra.ModelVariant}
                    onChange={(event) => {
                        let vra = this.state.vra;
                        vra.ModelVariant = event.target.value;
                        this.setState({
                            vra,
                        });
                    }}
                />
                <br />
                <br />
                <TextField
                    className="inputs"
                    label="Model Year"
                    variant="outlined"
                    value={this.state.vra.ModelYear}
                    onChange={(event) => {
                        let vra = this.state.vra;
                        vra.ModelYear = event.target.value;
                        this.setState({
                            vra,
                        });
                    }}
                />
                <br />
                <br />
                <TextField
                    className="inputs"
                    label="Color"
                    variant="outlined"
                    value={this.state.vra.Color}
                    onChange={(event) => {
                        let vra = this.state.vra;
                        vra.Color = event.target.value;
                        this.setState({
                            vra,
                        });
                    }}
                />
                <br />
                <br />
                <TextField
                    className="inputs"
                    label="Engine No."
                    variant="outlined"
                    value={this.state.vra.EngineNo}
                    onChange={(event) => {
                        let vra = this.state.vra;
                        vra.EngineNo = event.target.value;
                        this.setState({
                            vra,
                        });
                    }}
                />
                <br />
                <br />
                <TextField
                    className="inputs"
                    label="Chassis No."
                    variant="outlined"
                    value={this.state.vra.ChassisNo}
                    onChange={(event) => {
                        let vra = this.state.vra;
                        vra.ChassisNo = event.target.value;
                        this.setState({
                            vra,
                        });
                    }}
                />
                <br />
                <br />
                <TextField
                    className="inputs"
                    label="Owner UID"
                    variant="outlined"
                    value={this.state.vra.Owner}
                    onChange={(event) => {
                        let vra = this.state.vra;
                        vra.Owner = event.target.value;
                        this.setState({
                            vra,
                        });
                    }}
                />
                <br />
                <br />
                <Button onClick={this.onAddApp} variant="contained" color="primary">
                    Init. VR Application
                </Button>
            </div>
        );
    }
}

export default App;
