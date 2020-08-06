import React, { Component } from "react";
//import { Redirect } from "react-router-dom";

import { TextField, Button, CircularProgress } from "@material-ui/core";

class App extends Component {
    state = {
        RtoID: "",
        message: "",
    };

    onAddApp = async () => {
        const requestOptions = {
            method: "POST",
            headers: { "Content-Type": "application/json", "x-access-token": localStorage.getItem("session") },
            body: JSON.stringify({
                payload: {
                    RtoID: this.state.RtoID,
                },
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

        let response = await fetch("http://192.168.1.30:3000/api/main/rto/newDLA", requestOptions);
        let res = await response.json();
        console.log(res);
        this.setState({ message: "DL Application ID: " + res.id });
    };

    render() {
        return (
            <div>
                <h2>New DL Application</h2>
                {this.state.message}
                <TextField
                    className="inputs"
                    label="RTO ID"
                    variant="outlined"
                    value={this.state.RtoID}
                    onChange={(event) => {
                        this.setState({ RtoID: event.target.value });
                    }}
                />
                <br />
                <br />
                <Button onClick={this.onAddApp} variant="contained" color="primary">
                    Init. DL Application
                </Button>
            </div>
        );
    }
}

export default App;
