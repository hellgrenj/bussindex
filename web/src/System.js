import React, { useEffect }  from "react";
import { useDispatch, useSelector } from 'react-redux'  
import { fetchSystems, postSystem, systemsSelector } from './slices/systems'
import "./System.css"
function System() {
  const dispatch = useDispatch()
  const { systems, loading, hasErrors } = useSelector(systemsSelector)	
  
  // dispatch our fetchSystems thunk when component first mounts
   useEffect(() => {
    dispatch(fetchSystems())
  }, [dispatch])


  const renderSystems = () => {
    if (loading) return <p>Loading systems...</p>
    if (hasErrors) return <p>Cannot display systems...</p>
    console.log(systems[0])
    return systems.map(system =>
      
      <span key={system.ID} className="systemCards card blue-grey darken-1">
        {system.Description}
      </span>
    )
  }

  const createSystem = (e) => {
    if (e.key === 'Enter') {
      console.log('do validate');
      dispatch(postSystem({description:e.target.value}))
    }
  }

  return (
    <>
    <div className="row">
      <h2>System</h2>
      <div className="col s6"><input placeholder="ange namn och tryck enter" onKeyDown={createSystem}></input></div>
      </div>
      {renderSystems()}
    </>
  );
}
export default System;
