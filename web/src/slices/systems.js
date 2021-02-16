import { createSlice } from '@reduxjs/toolkit'

export const initialState = {
  loading: false,
  hasErrors: false,
  systems: [],
}

// A slice for systems with our 3 reducers
const systemsSlice = createSlice({
  name: 'systems',
  initialState,
  reducers: {
    getSystems: state => {
      state.loading = true
    },
    getSystemsSuccess: (state, { payload }) => {
      state.systems = payload
      state.loading = false
      state.hasErrors = false
    },
    getSystemsFailure: state => {
      state.loading = false
      state.hasErrors = true
    },
  },
})

// Three actions generated from the slice
export const { getSystems, getSystemsSuccess, getSystemsFailure } = systemsSlice.actions

// A selector
export const systemsSelector = state => state.systems

// The reducer
export default systemsSlice.reducer

// Asynchronous thunk actions
export function fetchSystems() {
  return async dispatch => {
    dispatch(getSystems())

    try {
      const response = await fetch('http://localhost:8080/system')
      const data = await response.json()
     
      dispatch(getSystemsSuccess(data))
    } catch (error) {
      dispatch(getSystemsFailure())
    }
  }
}
export function postSystem(system) {
    return async dispatch => {
     
      try {
        console.log('posting')
        const response = await fetch('http://localhost:8080/system',  {
            method: 'POST',
            headers: {
              'Accept': 'application/json',
              'Content-Type': 'application/json'
            },
            body: JSON.stringify(system)
          })
        const data = await response.json()
        console.log(data)
        dispatch(fetchSystems())
      } catch (error) {
        dispatch(getSystemsFailure())
      }
    }
  }

