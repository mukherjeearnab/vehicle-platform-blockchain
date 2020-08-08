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
                        Press <b>Load Citizen Profile</b> to View Profile Details
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

        let response = await fetch("http://192.168.1.30:3000/api/main/profile/get/" + this.state.ID, requestOptions);
        let res = await response.json();
        console.log(res);

        this.setState({ application: res });

        var output = <div>{this.createContent()}</div>;

        this.setState({ message: output });
    };

    createContent = () => {
        let PUCC, INS;
        console.log(this.state.application);
        if (this.state.application.Vehicles)
            INS = this.state.application.Vehicles.slice(0)
                .reverse()
                .map((content, index) => {
                    return (
                        <tr>
                            <td style={{ border: "1px solid black" }}>
                                <Link to={"/rto/transferOwner/" + content}>{content}</Link>
                            </td>
                        </tr>
                    );
                });
        if (this.state.application.LicenseNumbers)
            PUCC = this.state.application.LicenseNumbers.slice(0)
                .reverse()
                .map((content, index) => {
                    return (
                        <tr>
                            <td style={{ border: "1px solid black" }}>
                                <Link to={"/rto/viewLicense/" + content}>{content}</Link>
                            </td>
                        </tr>
                    );
                });

        return (
            <div>
                <h3>UID: {this.state.application.UID}</h3>
                <h3>Name: {this.state.application.Name}</h3>
                <h3>Date Of Birth: {this.state.application.DOB}</h3>
                <h3>Address: {this.state.application.Address}</h3>
                <h3>Phone: {this.state.application.Phone}</h3>
                <h2>License Record</h2>
                <table style={{ border: "1px solid black" }}>
                    <tr style={{ border: "1px solid black" }}>
                        <th>Lincense Number</th>
                    </tr>
                    {PUCC}
                </table>
                <h2>Vehicle Ownership Record</h2>
                <table style={{ border: "1px solid black" }}>
                    <tr style={{ border: "1px solid black" }}>
                        <th>Vehicle Reg No</th>
                    </tr>
                    {INS}
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
                    label="Citizen UID"
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
                    Load Citizen Profile
                </Button>
                <br /> <br />
                {this.state.message}
            </div>
        );
    }
}

export default App;
