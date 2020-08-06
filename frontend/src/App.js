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
            </ThemeProvider>
        </div>
    );
}

export default App;
