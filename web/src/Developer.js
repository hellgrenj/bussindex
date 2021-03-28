import React from "react";
import { useDispatch, useSelector } from "react-redux";
import {useEffect, useState, useRef} from "react";
import { deleteDeveloperThunk, developersSelector, fetchDevelopersThunk, postDeveloperThunk} from "./slices/developers";
import "materialize-css/dist/css/materialize.min.css";
import "materialize-css/dist/js/materialize.min.js";
import "./Developer.css";

function Developer() {
  const { developers, loading, hasErrors, errorMessage } = useSelector(
    developersSelector,
  );
  const dispatch = useDispatch();
    // local state
    const [developerName, setDeveloperName] = useState(null);
    const [developerDoE, setDeveloperDoE] = useState(null); 

    // refs
    const nameEl = useRef(null);
    const doeEl = useRef(null);
    const nameChangeHandler = (e) => {
      setDeveloperName(e.target.value);
    }

    const doeChangeHandler = (date) => {
      console.log(date)
      setDeveloperDoE(date)
    }
  const createDeveloper = () => {
    dispatch(postDeveloperThunk({name: developerName, dateOfEmployment: developerDoE}))
    setDeveloperDoE(null)
    setDeveloperName(null)
    nameEl.current.value = '';
    doeEl.current.value = '';
  }
  const deleteDeveloper = (developerId) => {
    dispatch(deleteDeveloperThunk(developerId))
  };
  const listDevelopers = () => {
    return developers.map((dev) =>
      <span
        key={dev.ID}
        value={dev.ID}
        className="developerCards card blue-grey darken-1"
      >
        {dev.Name} <br/>
        {new Date(dev.DateOfEmployment).toISOString().split('T')[0]}
        <i
            className="material-icons action" onClick={() => deleteDeveloper(dev.ID)}
          >
            delete
          </i>
      </span>
    );
  };
  const renderDevelopers = () => {
    if (hasErrors) {
      return (
        <>
          <p>{errorMessage}</p>
          {listDevelopers()}
        </>
      );
    }
    if (developers) {
      return listDevelopers();
    }
  };
  useEffect(() => {
    var el = document.querySelectorAll('.datepicker');
    M.Datepicker.init(el, {format:'yyyy-mm-dd', onSelect: doeChangeHandler,
    i18n: {
      months:['Januari','Februari','Mars','April','Maj','Juni','Juli','Augusti','September','Oktober', 'November', 'December'],
      monthsShort:['Jan','Feb','Mar','Apr','Maj','Jun','Jul','Aug','Sep','Okt', 'Nov', 'Dec']
    }
    });
  }, []);
  useEffect(() => {
    dispatch(fetchDevelopersThunk())
  }, []);
  return (
    <>
   
   <div className="row">
   <h2>Utvecklare</h2>
        <div className="col s6">
          <input ref={nameEl} placeholder="Namn" onChange={nameChangeHandler}></input>
          <input ref={doeEl} type="text" placeholder="Anställningsdatum" className="datepicker"></input>
          <button className="btn-large" onClick={createDeveloper}>Skapa utvecklare</button>
        </div>
      </div>
      <div className="row">
        {renderDevelopers()}
      </div>
   </>
  );
}
export default Developer;
