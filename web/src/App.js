import React from "react";
import System from "./System";
import Worker from "./Worker";
import Home from "./Home";
import { BrowserRouter as Router, Link, Route, Switch } from "react-router-dom";
import "./App.css";
function App() {
  
  return (
    <Router>
      <nav>
        <div className="nav-wrapper">
          <Link to="/" className="brand-logo">bussindex</Link>

          <ul id="nav-mobile" className="right hide-on-med-and-down">
            <li>
              <Link to="/system">System</Link>
            </li>
            <li>
              <Link to="/worker">Worker</Link>
            </li>
          </ul>
        </div>
      </nav>
      <div className="container">
        <div className="row">
          <div className="col s12 mainView">
            <Switch>
              <Route path="/system">
                <System />
              </Route>
              <Route path="/worker">
                <Worker />
              </Route>
              <Route path="/">
                <Home />
              </Route>
            </Switch>
          </div>
        </div>
      </div>
    </Router>
  );
}
export default App;
