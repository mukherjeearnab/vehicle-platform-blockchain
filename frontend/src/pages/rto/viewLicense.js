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
                        Press <b>Load License</b> to View License Details
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

        let response = await fetch("http://192.168.1.30:3000/api/main/rto/getDL/" + this.state.ID, requestOptions);
        let res = await response.json();
        console.log(res);

        this.setState({ application: res });

        var output = <div>{this.createContent()}</div>;

        this.setState({ message: output });
    };

    createContent = () => {
        return (
            <div>
                <h3>
                    UID:{" "}
                    <Link to={"/citizen/viewProfile/" + this.state.application.UID}>{this.state.application.UID}</Link>
                </h3>
                <h3>License Number: {this.state.application.LicenseNumber}</h3>
                <h3>Vehicle Type: {this.state.application.VehicleType}</h3>
                <h3>ExpiryDate: {this.state.application.ExpiryDate}</h3>
                <h3>Signing EmployeeID: {this.state.application.Employee}</h3>
                <br />
                <br />
            </div>
        );
    };

    render() {
        return (
            <div>
                <h2>View Driving License</h2>
                <TextField
                    className="inputs"
                    label="License Number"
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
                    Load License
                </Button>
                <br /> <br />
                {this.state.message}
            </div>
        );
    }
}

export default App;
