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

        let response = await fetch("http://192.168.1.30:3000/api/main/pollution/get/" + this.state.ID, requestOptions);
        let res = await response.json();
        console.log(res);

        this.setState({ application: res });

        var output = <div>{this.createContent()}</div>;

        this.setState({ message: output });
    };

    createContent = () => {
        return (
            <div>
                <h3>Cert ID: {this.state.application.CertID}</h3>
                <h3>DateTime: {this.state.application.DateTime}</h3>
                <h3>
                    Vehicle Reg No:{" "}
                    <Link to={"/rto/transferOwner/" + this.state.application.VehicleRegNo}>
                        {this.state.application.VehicleRegNo}
                    </Link>
                </h3>
                <h3>Signing EmployeeID: {this.state.application.Employee}</h3>
                <table style={{ border: "1px solid black" }}>
                    <tr style={{ border: "1px solid black" }}>
                        <th>Certificate Content</th>
                    </tr>
                    <tr style={{ border: "1px solid black" }}>
                        <td>{this.state.application.Content}</td>
                    </tr>
                </table>
                <br />
                <br />
            </div>
        );
    };

    render() {
        return (
            <div>
                <h2>View Pollution Under Control Certificate</h2>
                <TextField
                    className="inputs"
                    label="Certificate ID"
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
                    Load Certificate
                </Button>
                <br /> <br />
                {this.state.message}
            </div>
        );
    }
}

export default App;
