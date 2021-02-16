import { combineReducers } from 'redux'

import systemsReducer from './systems'

const rootReducer = combineReducers({
  systems: systemsReducer,
})

export default rootReducer