import React, { Component } from "react";
import axios from "axios";
//import { Redirect } from "react-router-dom";

import { TextField, Button, CircularProgress } from "@material-ui/core";

class App extends Component {
    state = {
        selectedFile: null,
        vra: {},
        message: "",
    };

    // On file upload (click the upload button)
    onFileUpload = async () => {
        this.setState({
            ID: (
                <span>
                    <CircularProgress />
                    <br></br> Loading.....
                </span>
            ),
        });

        // Create an object of formData
        const formData = new FormData();

        // Update the formData object
        formData.append("file", this.state.selectedFile, this.state.selectedFile.name);

        formData.append("payload", JSON.stringify(this.state.vra));

        let config = {
            headers: { "Content-Type": "application/json", "x-access-token": localStorage.getItem("session") },
        };

        // Details of the uploaded file
        console.log(this.state.selectedFile);

        // Request made to the backend api
        // Send formData object
        var reply = await axios.post("http://192.168.1.30:3000/api/main/trafficviolation/add", formData, config);
        console.log(reply);
        this.setState({ ID: "Violation ID: " + reply.data.TVID });
    };

    fileData = () => {
        if (this.state.selectedFile) {
            return (
                <div>
                    <h2>File Details:</h2>
                    <p>File Name: {this.state.selectedFile.name}</p>
                    <p>File Type: {this.state.selectedFile.type}</p>
                    <h2>{this.state.ID}</h2>
                </div>
            );
        } else {
            return (
                <div>
                    <br />
                    <h4>Choose before Pressing the Upload button</h4>
                </div>
            );
        }
    };

    // On file select (from the pop up)
    onFileChange = (event) => {
        // Update the state
        this.setState({ selectedFile: event.target.files[0] });
    };

    render() {
        return (
            <div>
                <h2>File Traffic Violation</h2>
                {this.state.message}
                <TextField
                    className="inputs"
                    label="Date Incident"
                    variant="outlined"
                    value={this.state.vra.DateIncident}
                    onChange={(event) => {
                        let vra = this.state.vra;
                        vra.DateIncident = event.target.value;
                        this.setState({
                            vra,
                        });
                    }}
                />
                <br />
                <br />
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
                    label="Vehicle Reg No"
                    variant="outlined"
                    value={this.state.vra.VehicleRegNo}
                    onChange={(event) => {
                        let vra = this.state.vra;
                        vra.VehicleRegNo = event.target.value;
                        this.setState({
                            vra,
                        });
                    }}
                />
                <br />
                <br />
                <TextField
                    className="inputs"
                    label="Content"
                    variant="outlined"
                    value={this.state.vra.Content}
                    onChange={(event) => {
                        let vra = this.state.vra;
                        vra.Content = event.target.value;
                        this.setState({
                            vra,
                        });
                    }}
                />
                <br />
                <br />
                <TextField variant="outlined" type="file" onChange={this.onFileChange} />
                <br />
                {this.fileData()}
                <br />
                <Button onClick={this.onFileUpload} variant="contained" color="primary">
                    File New Traffic Violation
                </Button>
            </div>
        );
    }
}

export default App;
