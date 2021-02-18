import React, { useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { fetchSystemsThunk, postSystemThunk, deleteSystemThunk, systemsSelector } from "./slices/systems";
import "./System.css";
function System() {
  const dispatch = useDispatch();
  const { systems, loading, hasErrors, errorMessage } = useSelector(
    systemsSelector,
  );

  // local state
  const [selectedSystem, setSelectedSystem] = useState(null);
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
        {system.Name}
        
          <i
            className="material-icons action"
            onClick={() => deleteSystem(system.ID)}
          >
            delete
          </i>
          <i className="material-icons action" onClick={() => editSystem(system.Name)}>
            edit
          </i>
        
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
  const renderDevelopersWorkinOnSystem = () => {
    if(selectedSystem) {
      return (
        <>
        <div className="row">
          <div className="col s6 developers-section">
          <i
            className="material-icons action"
            onClick={() => setSelectedSystem(null)}
          >close</i>
            <ul className="collection with-header">
              <li className="collection-header"><h5>Utvecklare som arbetar med <b>{selectedSystem}</b></h5></li>
              <li className="collection-item">Johan</li>
              <li className="collection-item active-developer">Fredrik</li>
              <li className="collection-item">Viktor</li>
              <li className="collection-item active-developer">Sofia</li>
            </ul>
          </div>
        </div>
        </>
      );
    } else {
      return null
    }
   
  }
  const createSystem = (e) => {
    if (e.key === "Enter") {
      console.log("do validate");
      dispatch(postSystemThunk({ name: e.target.value }));
      e.target.value = "";
    }
  };

  const deleteSystem = (systemId) => {
    setSelectedSystem(null)
    console.log("delete system", systemId);
    dispatch(deleteSystemThunk(systemId))
  };
  const editSystem = (systemName) => {
    console.log("edit system", systemName);
    setSelectedSystem(systemName)
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
      <div className="row">
        {renderSystems()}
      </div>
      {renderDevelopersWorkinOnSystem()}
    </>
  );
}
export default System;
