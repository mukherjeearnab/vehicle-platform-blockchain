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
                        Press <b>LOAD INVESTIGATION</b> to View Investigation
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

        let response = await fetch("http://192.168.1.30:3000/api/main/rto/updateDLA/" + this.state.ID, requestOptions);
        let res = await response.json();
        console.log(res);
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
                <h3>UID: {this.state.application.UID}</h3>
                <h3>RTO ID: {this.state.application.RtoID}</h3>
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
                <h2>View DL Application</h2>
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
