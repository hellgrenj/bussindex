import React, { useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { developersSelector } from "./slices/developers";
import { fetchSystemsThunk, addDevToSystemThunk, removeDevFromSystemThunk, postSystemThunk, deleteSystemThunk, systemsSelector } from "./slices/systems";
import { fetchDevelopersThunk} from "./slices/developers";
import "./System.css";
function System() {
  const dispatch = useDispatch();
  const { systems, loading, hasErrors, errorMessage } = useSelector(
    systemsSelector,
  );
  const { developers } = useSelector(developersSelector);

  // local state
  const [selectedSystem, setSelectedSystem] = useState(null);

  // dispatch our fetchSystems thunk when component first mounts
  useEffect(() => {
    dispatch(fetchSystemsThunk());
    dispatch(fetchDevelopersThunk());
  }, [dispatch]);

  useEffect(() => {
    if(selectedSystem) { 
      setSelectedSystem(systems.filter(s => s.ID == selectedSystem.ID)[0])
    }
  }, [systems])

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
          <i className="material-icons action" onClick={() => editSystem(system)}>
            edit
          </i>
        
      </span>
    );
  };
  const renderSystems = () => {
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
  const listDevelopers = () => {
    if(developers) {  
      return developers.map((dev) => {
      if(selectedSystem.DevIds && selectedSystem.DevIds.includes(dev.ID)) {
        return <li key={dev.ID} className="collection-item active-developer" value={dev.ID} onClick={removeDevFromSystem}>{dev.Name}</li> 
      } else {
       return <li key={dev.ID} className="collection-item" value={dev.ID} onClick={addDevToSystem}>{dev.Name}</li>
      }
      
    });
    } else {
      return "Inga utvecklare är inlagda i systemet"
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
              <li className="collection-header"><h5>Markerade utvecklare arbetar med <b>{selectedSystem.Name}</b></h5></li>
             {listDevelopers()}
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
  const addDevToSystem = (e) => {
    const devId = e.target.value;
    dispatch(addDevToSystemThunk(selectedSystem.ID, devId))
  }
  const removeDevFromSystem = (e) => {
    const devId = e.target.value;
    dispatch(removeDevFromSystemThunk(selectedSystem.ID, devId))
  }
  const deleteSystem = (systemId) => {
    setSelectedSystem(null)
    console.log("delete system", systemId);
    dispatch(deleteSystemThunk(systemId))
  };
  const editSystem = (system) => {
    console.log("edit system", system.Name);
    setSelectedSystem(system)
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
