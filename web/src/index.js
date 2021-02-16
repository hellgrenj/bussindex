import React from "react";
import ReactDOM from "react-dom";
import { configureStore } from '@reduxjs/toolkit'
import { Provider } from 'react-redux'
import "regenerator-runtime/runtime.js";
import "./index.css";
import App from "./App";
import rootReducer from "./slices";
const store = configureStore({ reducer: rootReducer });
ReactDOM.render(
  <Provider store={store}>
    <App />
  </Provider>,
  document.getElementById("root"),
);
