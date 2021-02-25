import { createSlice } from "@reduxjs/toolkit";

export const initialState = {
  loading: false,
  hasErrors: false,
  errorMessage: null,
  developers: [],
};

// A slice for developers with our 3 reducers
const developersSlice = createSlice({
  name: "developers",
  initialState,
  reducers: {
    getDevelopers: (state) => {
      state.loading = true;
    },
    getDevelopersSuccess: (state, { payload }) => {
      state.developers = payload;
      state.loading = false;
      state.hasErrors = false;
    },
    getDevelopersFailure: (state, { payload }) => {
      state.loading = false;
      state.hasErrors = true;
      state.errorMessage = payload;
    },
  },
});

// Three actions generated from the slice
export const { getDevelopers, getDevelopersSuccess, getDevelopersFailure } =
  developersSlice.actions;

// A selector
export const developersSelector = (state) => state.developers;

// The reducer
export default developersSlice.reducer;

// Asynchronous thunk actions
export function fetchDevelopersThunk() {
    return async (dispatch) => {
      dispatch(getDevelopers());
  
      try {
        const response = await fetch("http://localhost:8080/developer");
        const data = await response.json();
        dispatch(getDevelopersSuccess(data.result));
      } catch (error) {
        const data = error.json();
        dispatch(getDevelopersFailure(data.result));
      }
    };
  }
export function postDeveloperThunk(developer) {
  return async (dispatch) => {
    try {
      console.log("posting");
      const response = await fetch("http://localhost:8080/developer", {
        method: "POST",
        headers: {
          "Accept": "application/json",
          "Content-Type": "application/json",
        },
        body: JSON.stringify(developer),
      });
      const data = await response.json();
      console.log(data);
      if(data.success) {
        dispatch(fetchDevelopersThunk());
      } else {
        console.log(data.result)
        dispatch(getDevelopersFailure(data.result));
      }
      
    } catch (error) {
      dispatch(getDevelopersFailure());
    }
  };
}
