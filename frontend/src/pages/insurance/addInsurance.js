import React, { Component } from "react";
import { Link } from "react-router-dom";
import { TextField, Button, CircularProgress } from "@material-ui/core";

class App extends Component {
    state = {
        application: {},
        message: "",
        ID: "",
        contt: "",
        contt2: "",
    };

    onLoad = async () => {
        console.log(this.state.application);
        const requestOptions = {
            method: "GET",
            headers: { "Content-Type": "application/json", "x-access-token": localStorage.getItem("session") },
        };

        this.setState({
            message: (
                <span>
                    <CircularProgress />
                    <br></br> Loading.....
                </span>
            ),
        });

        let response = await fetch("http://192.168.1.30:3000/api/main/vehicle/get/" + this.state.ID, requestOptions);
        let res = await response.json();
        console.log(res);

        this.setState({ application: res });

        var output = <div>{this.createContent()}</div>;

        this.setState({ message: output });
    };

    onUpdateStatus = async () => {
        console.log(this.state.contt);
        const requestOptions = {
            method: "POST",
            headers: { "Content-Type": "application/json", "x-access-token": localStorage.getItem("session") },
            body: JSON.stringify({
                payload: {
                    VehicleRegNo: this.state.application.RegNo,
                    UID: this.state.application.Owner,
                    Duration: this.state.contt,
                    Content: this.state.contt2,
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

        let response = await fetch("http://192.168.1.30:3000/api/main/insurance/policy/add", requestOptions);
        let res = await response.json();
        console.log(res);
        this.setState({ message: res.message + " | " + res.id });
    };

    createContent = () => {
        let OWNER, PUCC, INS;
        console.log(this.state.application);
        if (this.state.application.OwnershipHistory)
            OWNER = this.state.application.OwnershipHistory.slice(0)
                .reverse()
                .map((content, index) => {
                    const date = new Date(content.Date).toString();

                    return (
                        <tr>
                            <td style={{ border: "1px solid black" }}>{content.Prev}</td>
                            <td style={{ border: "1px solid black" }}>{content.Curr}</td>
                            <td style={{ border: "1px solid black" }}>{content.Employee}</td>
                            <td style={{ border: "1px solid black" }}>{date}</td>
                        </tr>
                    );
                });

        if (this.state.application.Insurance)
            INS = this.state.application.Insurance.slice(0)
                .reverse()
                .map((content, index) => {
                    return (
                        <tr>
                            <td style={{ border: "1px solid black" }}>
                                <Link to={"/pollution/viewCert/" + content}>{content}</Link>
                            </td>
                        </tr>
                    );
                });
        if (this.state.application.PollutionCert)
            PUCC = this.state.application.PollutionCert.slice(0)
                .reverse()
                .map((content, index) => {
                    return (
                        <tr>
                            <td style={{ border: "1px solid black" }}>
                                <Link to={"/pollution/viewCert/" + content}>{content}</Link>
                            </td>
                        </tr>
                    );
                });

        return (
            <div>
                <h3>Reg No: {this.state.application.RegNo}</h3>
                <h3>Make: {this.state.application.Make}</h3>
                <h3>Model: {this.state.application.Model}</h3>
                <h3>ModelVariant: {this.state.application.ModelVariant}</h3>
                <h3>ModelYear: {this.state.application.ModelYear}</h3>
                <h3>Color: {this.state.application.Color}</h3>
                <h3>Engine No: {this.state.application.EngineNo}</h3>
                <h3>Chassis No: {this.state.application.ChassisNo}</h3>
                <h3>Owner: {this.state.application.Owner}</h3>
                <h3>Creator: {this.state.application.Creator}</h3>
                <h2>PUCC Record</h2>
                <table style={{ border: "1px solid black" }}>
                    <tr style={{ border: "1px solid black" }}>
                        <th>Certificate ID</th>
                    </tr>
                    {PUCC}
                </table>
                <h2>Insurance Record</h2>
                <table style={{ border: "1px solid black" }}>
                    <tr style={{ border: "1px solid black" }}>
                        <th>Policy ID</th>
                    </tr>
                    {INS}
                </table>
                <h2>Ownership History</h2>
                <table style={{ border: "1px solid black" }}>
                    <tr style={{ border: "1px solid black" }}>
                        <th>Prev Owner</th>
                        <th>Curr Owner</th>
                        <th>Employee ID</th>
                        <th>Date</th>
                    </tr>
                    {OWNER}
                </table>
                <br />
                <br />
            </div>
        );
    };

    render() {
        return (
            <div>
                <h2>View Vehicle Profile</h2>
                <TextField
                    className="inputs"
                    label="Vehicle Reg No"
                    variant="outlined"
                    value={this.state.ID}
                    onChange={(event) => {
                        this.setState({
                            ID: event.target.value,
                        });
                    }}
                />
                <br /> <br />
                <Button onClick={this.onLoad} variant="contained" color="primary">
                    Load Profile
                </Button>
                <br /> <br />
                {this.state.message}
                <h4>Insurance Policy Params</h4>
                <TextField
                    className="inputs"
                    label="Duration"
                    variant="outlined"
                    value={this.state.contt}
                    onChange={(event) => {
                        console.log(this.state.contt);
                        this.setState({
                            contt: event.target.value,
                        });
                    }}
                />
                <TextField
                    className="inputs"
                    label="Content"
                    variant="outlined"
                    value={this.state.contt2}
                    onChange={(event) => {
                        console.log(this.state.contt2);
                        this.setState({
                            contt2: event.target.value,
                        });
                    }}
                />
                <br /> <br />
                <Button onClick={this.onUpdateStatus} variant="contained" color="primary">
                    Add Policy
                </Button>
            </div>
        );
    }
}

export default App;
