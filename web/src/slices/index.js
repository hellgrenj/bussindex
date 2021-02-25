import { combineReducers } from 'redux'
import developersReducer from './developers'
import systemsReducer from './systems'

const rootReducer = combineReducers({
  systems: systemsReducer,
  developers: developersReducer
})

export default rootReducer