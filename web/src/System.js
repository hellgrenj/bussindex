import React, { useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import { fetchSystemsThunk, postSystemThunk, deleteSystemThunk, systemsSelector } from "./slices/systems";
import "./System.css";
function System() {
  const dispatch = useDispatch();
  const { systems, loading, hasErrors, errorMessage } = useSelector(
    systemsSelector,
  );

  // dispatch our fetchSystems thunk when component first mounts
  useEffect(() => {
    dispatch(fetchSystemsThunk());
  }, [dispatch]);

  const listSystems = () => {
    return systems.map((system) =>
      <span
        key={system.ID}
        value={system.ID}
        className="systemCards card blue-grey darken-1"
      >
        <h5>{system.Name}</h5>
        <p>
          <span
            className="material-icons action"
            onClick={() => deleteSystem(system.ID)}
          >
            delete
          </span>
          <i className="material-icons action" onClick={() => editSystem(system.ID)}>
            edit
          </i>
        </p>
      </span>
    );
  };
  const renderSystems = () => {
    if (loading) return <p>Loading systems...</p>;
    if (hasErrors) {
      return (
        <>
          <p>{errorMessage}</p>
          {listSystems()}
        </>
      );
    }
    if (systems) {
      return listSystems();
    }
  };

  const createSystem = (e) => {
    if (e.key === "Enter") {
      console.log("do validate");
      dispatch(postSystemThunk({ name: e.target.value }));
      e.target.value = "";
    }
  };

  const deleteSystem = (systemId) => {
    console.log("delete system", systemId);
    dispatch(deleteSystemThunk(systemId))
  };
  const editSystem = (systemId) => {
    console.log("edit system", systemId);
  };

  return (
    <>
      <div className="row">
        <h2>System</h2>
        <div className="col s6">
          <input
            placeholder="ange namn och tryck enter"
            onKeyDown={createSystem}
          >
          </input>
        </div>
      </div>
      {renderSystems()}
    </>
  );
}
export default System;
