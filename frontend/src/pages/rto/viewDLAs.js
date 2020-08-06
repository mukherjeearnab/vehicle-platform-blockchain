import React, { Component } from "react";
import { Table, TableBody, TableContainer, TableHead, TableCell, TableRow, CircularProgress } from "@material-ui/core";
import Paper from "@material-ui/core/Paper";
import { Link, Redirect } from "react-router-dom";

class App extends Component {
    state = {
        DLAS: [],
        message: "",
    };

    async componentDidMount() {
        this.setState({
            message: (
                <span>
                    <CircularProgress />
                    <br></br> Loading.....
                </span>
            ),
        });

        if (!localStorage.getItem("session")) this.setState({ redirect: <Redirect to="/" /> });
        if (localStorage.getItem("session")) {
            const requestOptions = {
                method: "GET",
                headers: { "Content-Type": "application/json", "x-access-token": localStorage.getItem("session") },
            };
            let response = await fetch("http://192.168.1.30:3000/api/auth/verify/", requestOptions);
            let res = await response.json();
            if (res.status === 0) this.setState({ redirect: <Redirect to="/" /> });
        }
        console.log(localStorage.getItem("session"));
        const requestOptions = {
            method: "POST",
            headers: { "Content-Type": "application/json", "x-access-token": localStorage.getItem("session") },
            body: JSON.stringify({
                payload: {
                    Type: "DL",
                    RtoID: "",
                },
            }),
        };
        console.log(requestOptions);
        let response = await fetch("http://192.168.1.30:3000/api/main/rto/queryDLVRA/", requestOptions);
        let res = await response.json();
        console.log(res);
        this.setState({ DLAS: res, message: "" });
    }

    render() {
        return (
            <div>
                <h1>VR Applications</h1>
                <div>
                    <TableContainer component={Paper}>
                        <Table aria-label="simple table">
                            <TableHead>
                                <TableRow>
                                    <TableCell align="left">
                                        <b>Application ID</b>
                                    </TableCell>
                                    <TableCell align="left">
                                        <b>UID</b>
                                    </TableCell>
                                    <TableCell align="left">
                                        <b>RTO ID</b>
                                    </TableCell>
                                    <TableCell align="left">
                                        <b>DATE TIME</b>
                                    </TableCell>
                                </TableRow>
                            </TableHead>
                            <TableBody>
                                {this.state.DLAS.slice(0)
                                    .reverse()
                                    .map((content, index) => {
                                        content = content.Value;
                                        const date = new Date(content.DateTime).toString();

                                        return (
                                            <TableRow key={content.ID}>
                                                <TableCell align="left">
                                                    <Link to={`/firViewer/${content.ApplicationID}`}>
                                                        {content.ApplicationID}
                                                    </Link>
                                                </TableCell>
                                                <TableCell align="left">{content.UID}</TableCell>
                                                <TableCell align="left">{content.RtoID}</TableCell>
                                                <TableCell align="left">{date}</TableCell>
                                            </TableRow>
                                        );
                                    })}
                            </TableBody>
                        </Table>
                    </TableContainer>
                </div>
                <br />
                {this.state.message}
            </div>
        );
    }
}

export default App;
