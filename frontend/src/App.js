import React from "react";
import { ThemeProvider, CssBaseline } from "@material-ui/core";
import Theme from "./Theme";
import "./App.css";

import { Route, Link } from "react-router-dom";
import Login from "./pages/login";
import HomeCi from "./pages/HomeCi";
import HomePo from "./pages/HomePo";
import HomeIn from "./pages/HomeIn";
import HomePu from "./pages/HomePu";
import HomeRt from "./pages/HomeRt";

import newDLA from "./pages/rto/newDLA";
import newVRA from "./pages/rto/newVRA";
import viewVRAs from "./pages/rto/viewVRAs";
import viewDLAs from "./pages/rto/viewDLAs";
import updateDLA from "./pages/rto/updateDLA";
import updateVRA from "./pages/rto/updateVRA";
import newVR from "./pages/rto/newVR";
import newDL from "./pages/rto/newDL";
import transOwner from "./pages/rto/transferOwnership";
import newViolation from "./pages/police/newViolation";
import viewProfile from "./pages/citizen/viewProfile";
import viewLicense from "./pages/rto/viewLicense";
import newCert from "./pages/pollution/newCert";
import viewCert from "./pages/pollution/viewCert";

function App() {
    return (
        <div className="App">
            <ThemeProvider theme={Theme}>
                <CssBaseline />
                <Link to="/">
                    <h1>Vehicle Platform</h1>
                </Link>
                <Route exact path="/" component={Login}></Route>
                <Route exact path="/HomeCi" component={HomeCi}></Route>
                <Route exact path="/HomePo" component={HomePo}></Route>
                <Route exact path="/HomeIn" component={HomeIn}></Route>
                <Route exact path="/HomePu" component={HomePu}></Route>
                <Route exact path="/HomeRt" component={HomeRt}></Route>

                <Route exact path="/rto/newDLA" component={newDLA}></Route>
                <Route exact path="/rto/newVRA" component={newVRA}></Route>
                <Route exact path="/rto/viewVRAs" component={viewVRAs}></Route>
                <Route exact path="/rto/viewDLAs" component={viewDLAs}></Route>
                <Route exact path="/rto/updateDLA/:id" component={updateDLA}></Route>
                <Route exact path="/rto/updateVRA/:id" component={updateVRA}></Route>
                <Route exact path="/rto/createVP" component={newVR}></Route>
                <Route exact path="/rto/createDL" component={newDL}></Route>
                <Route exact path="/police/newViolation" component={newViolation}></Route>
                <Route exact path="/rto/transferOwner/:id" component={transOwner}></Route>
                <Route exact path="/citizen/viewProfile/:id" component={viewProfile}></Route>
                <Route exact path="/rto/viewLicense/:id" component={viewLicense}></Route>
                <Route exact path="/pollution/newCert" component={newCert}></Route>
                <Route exact path="/pollution/viewCert/:id" component={viewCert}></Route>
            </ThemeProvider>
        </div>
    );
}

export default App;
