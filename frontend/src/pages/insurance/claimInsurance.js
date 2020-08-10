import React, { Component } from "react";
import { Link } from "react-router-dom";
import { TextField, Button, CircularProgress } from "@material-ui/core";

class App extends Component {
    state = {
        application: {},
        message: "",
        ID: "",
        contt: "",
    };

    async componentDidMount() {
        const { id } = this.props.match.params;
        if (id !== "0") {
            this.setState({
                ID: id,
                message: (
                    <p>
                        Press <b>Load Vehicle Profile</b> to View Vehicle Details
                    </p>
                ),
            });
        }
    }

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

        let response = await fetch("http://192.168.1.30:3000/api/main/insurance/get/" + this.state.ID, requestOptions);
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
                    Content: this.state.contt,
                    InsurancePolicyID: this.state.ID,
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

        let response = await fetch("http://192.168.1.30:3000/api/main/insurance/claim/add", requestOptions);
        let res = await response.json();
        console.log(res);
        this.setState({ message: res.message + " | " + res.id });
    };

    createContent = () => {
        let PUCC;

        if (this.state.application.Claims)
            PUCC = this.state.application.Claims.slice(0)
                .reverse()
                .map((content, index) => {
                    return (
                        <tr>
                            <td style={{ border: "1px solid black" }}>
                                <Link to={"/isurance/viewClaim/" + content}>{content}</Link>
                            </td>
                        </tr>
                    );
                });

        return (
            <div>
                <h3>Insurance ID: {this.state.application.InsuranceID}</h3>
                <h3>Vehicle Reg No: {this.state.application.VehicleRegNo}</h3>
                <h3>Vehicle Owner: {this.state.application.UID}</h3>
                <h3>Date: {this.state.application.Date}</h3>
                <h3>Duration: {this.state.application.Duration}</h3>
                <h3>Content: {this.state.application.Content}</h3>
                <h2>Policy Claims</h2>
                <table style={{ border: "1px solid black" }}>
                    <tr style={{ border: "1px solid black" }}>
                        <th>Claim ID</th>
                    </tr>
                    {PUCC}
                </table>
                <br />
                <br />
            </div>
        );
    };

    render() {
        return (
            <div>
                <h2>View Insurance Policy</h2>
                <TextField
                    className="inputs"
                    label="Policy ID"
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
                    Load Policy
                </Button>
                <br /> <br />
                {this.state.message}
                <h4>Claim Policy</h4>
                <TextField
                    className="inputs"
                    label="Claim Content"
                    variant="outlined"
                    value={this.state.contt}
                    onChange={(event) => {
                        console.log(this.state.contt);
                        this.setState({
                            contt: event.target.value,
                        });
                    }}
                />
                <br /> <br />
                <Button onClick={this.onUpdateStatus} variant="contained" color="primary">
                    Add Insurance Claim
                </Button>
            </div>
        );
    }
}

export default App;
