import React, { Component } from "react";
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
                        Press <b>LOAD APPLICATION</b> to View Application
                    </p>
                ),
            });
        }
    }

    onLoadInvestigation = async () => {
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

        let response = await fetch("http://192.168.1.30:3000/api/main/rto/getDLVRA/" + this.state.ID, requestOptions);
        let res = await response.json();
        console.log(res);

        this.setState({ application: res });

        var output = <div>{this.createContent()}</div>;

        this.setState({ message: output });
    };

    onUpdateStatus = async () => {
        console.log(this.state.content);
        const requestOptions = {
            method: "POST",
            headers: { "Content-Type": "application/json", "x-access-token": localStorage.getItem("session") },
            body: JSON.stringify({
                payload: {
                    Content: this.state.contt,
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

        let response = await fetch("http://192.168.1.30:3000/api/main/rto/updateVRA/" + this.state.ID, requestOptions);
        let res = await response.json();
        console.log(res);
        this.setState({ message: res.message });
    };

    createContent = () => {
        let status;
        console.log(this.state.application);
        if (this.state.application.Status)
            status = this.state.application.Status.slice(0)
                .reverse()
                .map((content, index) => {
                    const date = new Date(content.Date).toString();

                    return (
                        <tr>
                            <td style={{ border: "1px solid black" }}>{date}</td>
                            <td style={{ border: "1px solid black" }}>{content.Employee}</td>
                            <td style={{ border: "1px solid black" }}>{content.Content}</td>
                        </tr>
                    );
                });

        return (
            <div>
                <h3>Application ID: {this.state.application.ApplicationID}</h3>
                <h3>Date Time: {this.state.application.DateTime}</h3>
                <h3>Make: {this.state.application.Make}</h3>
                <h3>Model: {this.state.application.Model}</h3>
                <h3>ModelVariant: {this.state.application.ModelVariant}</h3>
                <h3>ModelYear: {this.state.application.ModelYear}</h3>
                <h3>Color: {this.state.application.Color}</h3>
                <h3>Engine No: {this.state.application.EngineNo}</h3>
                <h3>Chassis No: {this.state.application.ChassisNo}</h3>
                <h3>Owner: {this.state.application.Owner}</h3>
                <h3>Creator: {this.state.application.Creator}</h3>
                <h2>Status</h2>
                <table style={{ border: "1px solid black" }}>
                    <tr style={{ border: "1px solid black" }}>
                        <th>Date</th>
                        <th>Employee ID</th>
                        <th>Content</th>
                    </tr>
                    {status}
                </table>
                <br />
                <br />
            </div>
        );
    };

    render() {
        return (
            <div>
                <h2>View VR Application</h2>
                <TextField
                    className="inputs"
                    label="Application ID"
                    variant="outlined"
                    value={this.state.ID}
                    onChange={(event) => {
                        this.setState({
                            ID: event.target.value,
                        });
                    }}
                />
                <br /> <br />
                <Button onClick={this.onLoadInvestigation} variant="contained" color="primary">
                    Load Application
                </Button>
                <br /> <br />
                {this.state.message}
                <TextField
                    className="inputs"
                    label="Status Content"
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
                    Update Application Status
                </Button>
            </div>
        );
    }
}

export default App;
